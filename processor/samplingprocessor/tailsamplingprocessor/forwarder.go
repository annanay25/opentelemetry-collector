// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tailsamplingprocessor

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.uber.org/zap"

	tracepb "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	"github.com/open-telemetry/opentelemetry-collector/config/configgrpc"
	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"github.com/open-telemetry/opentelemetry-collector/consumer/consumerdata"
	"github.com/open-telemetry/opentelemetry-collector/exporter"
	"github.com/open-telemetry/opentelemetry-collector/exporter/opencensusexporter"
	"github.com/open-telemetry/opentelemetry-collector/processor/samplingprocessor/tailsamplingprocessor/idbatcher"
)

type forwarder interface {
	// process span
	process(span *tracepb.Span) bool
}

type collectorPeer struct {
	exporter           exporter.TraceExporter
	ip                 string
	peerBatcher        idbatcher.Batcher
	spanDispatchTicker tTicker // to pop from the batcher and forward to peer
	logger             *zap.Logger
	start              sync.Once
	idToSpans          sync.Map
	globalDeleteChan   chan int
}

type ringMembershipExtensionClient struct {
	client   *http.Client
	endpoint string
}

// spanForwarder is the component that fowards spans to collector peers
// based on traceID
type spanForwarder struct {
	// to make peerQueues concurrently accessible
	sync.RWMutex
	// to start the memberSyncTicker
	start sync.Once
	// logger
	logger *zap.Logger
	// self member IP
	selfMemberIP string
	// ticker to call ring.GetState() and sync member list
	memberSyncTicker tTicker
	// stores queues for each of the collector peers
	peerQueues map[string]*collectorPeer
	// The ringmembership extension (implements extension.SupportExtension)
	// Ownership should be with the component that requires it.
	// However it will be Start()'ed by the Application service.
	// ringMembershipExtensionClient is an http client to get extension state
	ring *ringMembershipExtensionClient
	// channel to keep track of total number of traces in memory
	// across all queues
	// TODO(@annanay25): Is it good to use channels for this
	globalDeleteChan chan int
}

func (c *collectorPeer) batchDispatchOnTick() {
	c.logger.Debug("Collector peer batchDispatchOnTick invoked")
	batchIds, _ := c.peerBatcher.CloseCurrentAndTakeFirstBatch()
	// create batch from batchIds
	var td consumerdata.TraceData
	for _, v := range batchIds {
		if span, ok := c.idToSpans.Load(string(v)); ok {
			if span != nil {
				if span.(*v1.Span).Attributes == nil {
					span.(*v1.Span).Attributes = &v1.Span_Attributes{
						AttributeMap: make(map[string]*v1.AttributeValue),
					}
				} else if span.(*v1.Span).Attributes.AttributeMap == nil {
					span.(*v1.Span).Attributes.AttributeMap = make(map[string]*v1.AttributeValue)
				}
				span.(*v1.Span).Attributes.AttributeMap["otelcol.ttl"] = &v1.AttributeValue{
					Value: &v1.AttributeValue_IntValue{
						IntValue: 1,
					},
				}
				td.Spans = append(td.Spans, span.(*v1.Span))
				c.logger.Info("Forwarding this span", zap.ByteString("Span ID", span.(*v1.Span).GetSpanId()), zap.String("Peer", c.ip))

				// Null-out the map entry
				c.idToSpans.Delete(string(v))
			}
		}
	}

	stats.Record(
		context.Background(),
		statCountSpansForwarded.M(int64(len(batchIds))))

	// simply post this batch via grpc to the collector peer
	c.logger.Debug("Sending batch to collector peer")
	c.exporter.ConsumeTraceData(context.Background(), td)
}

func newCollectorPeer(logger *zap.Logger, ip string, globalDeleteChan chan int) *collectorPeer {
	factory := &opencensusexporter.Factory{}
	config := factory.CreateDefaultConfig()
	config.(*opencensusexporter.Config).ExporterSettings = configmodels.ExporterSettings{
		NameVal: "opencensus",
		TypeVal: "opencensus",
	}
	config.(*opencensusexporter.Config).GRPCSettings = configgrpc.GRPCSettings{
		Endpoint: ip + ":" + strconv.Itoa(55678),
	}

	logger.Info("Creating new collector peer instance for", zap.String("PeerIP", ip))
	exporter, err := factory.CreateTraceExporter(logger, config)
	if err != nil {
		logger.Fatal("Could not create span exporter", zap.Error(err))
		return nil
	}

	batcher, err := idbatcher.New(10, 64, uint64(runtime.NumCPU()))
	if err != nil {
		logger.Fatal("Could not create id batcher", zap.Error(err))
		return nil
	}

	cp := &collectorPeer{
		exporter:    exporter,
		logger:      logger,
		peerBatcher: batcher,
		ip:          ip,
	}
	cp.spanDispatchTicker = &policyTicker{onTick: cp.batchDispatchOnTick}
	return cp
}

// At the set frequency, get the state of the collector peer list
func (sf *spanForwarder) memberSyncOnTick() {
	// Get sorted member list from the extension
	newMembers, err := sf.ring.GetState()
	if err != nil {
		sf.logger.Info("(memberSyncOnTick) Error fetching members", zap.Error(err))
		return
	}

	sf.RLock()
	curMembers := make([]string, 0, len(sf.peerQueues))
	for k := range sf.peerQueues {
		curMembers = append(curMembers, k)
	}
	sf.RUnlock()

	// checking if curMembers == newMembers
	isEqual := true
	if len(curMembers) != len(newMembers) {
		isEqual = false
	} else {
		for k, v := range curMembers {
			if v != newMembers[k] {
				isEqual = false
			}
		}
	}

	if !isEqual {
		// Remove old members
		// Find diff(curMembers, newMembers)
		for _, c := range curMembers {
			// check if v is part of newMembers
			flag := 0
			for _, n := range newMembers {
				if c == n {
					flag = 1
				}
			}

			if flag == 0 {
				// Need a write lock here
				sf.Lock()
				// nullify the collector peer instance
				sf.peerQueues[c] = nil
				// delete the key
				delete(sf.peerQueues, c)
				sf.Unlock()
				sf.logger.Info("(memberSyncOnTick) Deleted member", zap.String("Member ip", c))
			}
		}

		// Add new members
		for _, v := range newMembers {
			if _, ok := sf.peerQueues[v]; ok {
				// exists, do nothing
			} else if v == sf.selfMemberIP {
				// Need a write lock here
				sf.Lock()
				sf.peerQueues[v] = nil
				sf.Unlock()
			} else {
				newPeer := newCollectorPeer(sf.logger, v, sf.globalDeleteChan)
				if newPeer == nil {
					return
				}
				// Need a write lock here
				sf.Lock()
				// build a new collectorPeer object
				sf.peerQueues[v] = newPeer
				sf.Unlock()
				sf.logger.Info("(memberSyncOnTick) Added member", zap.String("Member ip", v))
			}
		}
	}
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func newSpanForwarder(logger *zap.Logger, globalDeleteChan chan int, ringExtEndpoint string) (forwarder, error) {
	if ringExtEndpoint == "" {
		ringExtEndpoint = "127.0.0.1:13000"
	}

	sf := &spanForwarder{
		logger:           logger,
		globalDeleteChan: globalDeleteChan,
		ring: &ringMembershipExtensionClient{
			client:   &http.Client{},
			endpoint: ringExtEndpoint,
		},
	}

	if ip, err := externalIP(); err == nil {
		sf.selfMemberIP = ip
	} else {
		return nil, err
	}

	sf.peerQueues = make(map[string]*collectorPeer)
	sf.memberSyncTicker = &policyTicker{onTick: sf.memberSyncOnTick}

	return sf, nil
}

// Copied from github.com/dgryski/go-jump/blob/master/jump.go
func jumpHash(key uint64, numBuckets int) int32 {

	var b int64 = -1
	var j int64

	for j < int64(numBuckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}

	return int32(b)
}

func bytesToInt(spanID []byte) int64 {
	var n int64
	buf := bytes.NewBuffer(spanID)
	binary.Read(buf, binary.LittleEndian, &n)
	return n
}

// TODO(@annanay25): Use batch here instead of span?
func (sf *spanForwarder) process(span *tracepb.Span) bool {
	// Start member sync
	sf.start.Do(func() {
		sf.logger.Info("First span received, starting member sync timer")
		// Run first one manually
		sf.memberSyncOnTick()
		sf.memberSyncTicker.Start(100 * time.Millisecond)
	})

	var memberID string

	// The only time we need to acquire the lock is to see peer list
	sf.RLock()
	defer sf.RUnlock()

	// check hash of traceid
	traceIDInt64 := bytesToInt(span.TraceId)
	memberNum := int64(jumpHash(uint64(traceIDInt64), len(sf.peerQueues)))
	if memberNum == -1 {
		memberID = sf.selfMemberIP
	} else {
		memberID = reflect.ValueOf(sf.peerQueues).MapKeys()[memberNum].Interface().(string)
	}

	if memberID == sf.selfMemberIP {
		// span should be processed by this collector peer
		return false
	}

	// Append this span to the batch of that member
	peer := sf.peerQueues[memberID]

	peer.start.Do(func() {
		peer.logger.Info("First span received, starting collector peer timers")
		peer.spanDispatchTicker.Start(100 * time.Millisecond)
	})

	// there might be multiple spans of the same traceId
	// we don't want to build the trace here(?)
	peer.peerBatcher.AddToCurrentBatch(span.SpanId)

	// FIXME(@annanay25): Don't insert to globalDeleteChan for every span
	// ... Need to change this, track only trace constructions in memory.
	select {
	case sf.globalDeleteChan <- 1:
		// Store the span in idToSpans
		peer.idToSpans.Store(string(span.SpanId), span)
	default:
		// don't store this span
	}

	return true
}

func (r *ringMembershipExtensionClient) GetState() ([]string, error) {
	// Make an http client to the extension and get members
	resp, err := r.client.Get("http://" + r.endpoint + "/state")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []string
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}