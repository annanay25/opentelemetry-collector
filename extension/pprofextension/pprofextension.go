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

package pprofextension

import (
	"net"
	"net/http"
	_ "net/http/pprof" // Needed to enable the performance profiler
	"runtime"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-service/extension"
)

type pprofExtension struct {
	config Config
	logger *zap.Logger
	server http.Server
}

var _ (extension.ServiceExtension) = (*pprofExtension)(nil)

func (p *pprofExtension) Start(host extension.Host) error {
	// Start the listener here so we can have earlier failure if port is
	// already in use.
	ln, err := net.Listen("tcp", p.config.Endpoint)
	if err != nil {
		return err
	}

	runtime.SetBlockProfileRate(p.config.BlockProfileFraction)
	runtime.SetMutexProfileFraction(p.config.MutexProfileFraction)

	p.logger.Info("Starting net/http/pprof server", zap.Any("config", p.config))
	go func() {
		// The listener ownership goes to the server.
		if err := p.server.Serve(ln); err != nil && err != http.ErrServerClosed {
			host.ReportFatalError(err)
		}
	}()

	return nil
}

func (p *pprofExtension) Shutdown() error {
	return p.server.Close()
}

func newServer(config Config, logger *zap.Logger) (*pprofExtension, error) {
	p := &pprofExtension{
		config: config,
		logger: logger,
	}

	return p, nil
}
