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

package testbed

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestGeneratorAndBackend(t *testing.T) {
	mb := NewMockBackend("mockbackend.log")

	assert.EqualValues(t, 0, mb.SpansReceived())

	err := mb.Start(BackendJaeger)
	require.NoError(t, err, "Cannot start backend")

	defer mb.Stop()

	lg, err := NewLoadGenerator()
	require.NoError(t, err, "Cannot start load generator")

	assert.EqualValues(t, 0, lg.spansSent)

	// Generate at 1000 SPS
	lg.Start(LoadOptions{SpansPerSecond: 1000})

	// Wait until at least 50 spans are sent
	WaitFor(t, func() bool { return lg.SpansSent() > 50 }, "SpansSent > 50")

	lg.Stop()

	// The backend should receive everything generated.
	assert.Equal(t, lg.SpansSent(), mb.SpansReceived())
}

// WaitFor the specific condition for up to 10 seconds. Records a test error
// if condition does not become true.
func WaitFor(t *testing.T, cond func() bool, errMsg ...interface{}) bool {
	startTime := time.Now()

	// Start with 5 ms waiting interval between condition re-evaluation.
	waitInterval := time.Millisecond * 5

	for {
		time.Sleep(waitInterval)

		// Increase waiting interval exponentially up to 500 ms.
		if waitInterval < time.Millisecond*500 {
			waitInterval = waitInterval * 2
		}

		if cond() {
			return true
		}

		if time.Since(startTime) > time.Second*10 {
			// Waited too long
			t.Error("Time out waiting for", errMsg)
			return false
		}
	}
}
