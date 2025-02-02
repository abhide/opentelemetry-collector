// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package featuregate

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGate(t *testing.T) {
	enabled := &atomic.Bool{}
	enabled.Store(true)
	g := &Gate{
		id:             "test",
		description:    "test gate",
		enabled:        enabled,
		stage:          StageAlpha,
		referenceURL:   "http://example.com",
		removalVersion: "v0.64.0",
	}

	assert.Equal(t, "test", g.ID())
	assert.Equal(t, "test gate", g.Description())
	assert.True(t, g.IsEnabled())
	assert.Equal(t, StageAlpha, g.Stage())
	assert.Equal(t, "http://example.com", g.ReferenceURL())
	assert.Equal(t, "v0.64.0", g.RemovalVersion())
}
