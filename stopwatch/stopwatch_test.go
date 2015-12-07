// Copyright 2015, Quentin RENARD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stopwatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStopwatch(t *testing.T) {
	// Initialize
	id := "test"
	isEnabled := true
	s := NewStopwatch(id).SetIsEnabled(true)

	// Assert
	assert.Equal(t, id, s.ID())
	assert.Equal(t, isEnabled, s.IsEnabled())
}

func TestNewStopwatchFromConfiguration(t *testing.T) {
	// Initialize
	c := Configuration{
		ID:        "test",
		IsEnabled: true,
	}
	s := NewStopwatchFromConfiguration(c)

	// Assert
	assert.Equal(t, c.ID, s.ID())
	assert.Equal(t, c.IsEnabled, s.IsEnabled())
}

func TestAddEvent(t *testing.T) {
	// Initialize
	sEnabled := stopwatch{
		id:        "test",
		isEnabled: true,
	}
	sDisabled := stopwatch{
		id:        "test",
		isEnabled: false,
	}

	// Add event 1
	n := "event 1"
	d := "description 1"
	sEnabled.AddEvent(n, d)
	sDisabled.AddEvent(n, d)

	// Assert enabled
	assert.Equal(t, 1, len(sEnabled.events))
	assert.Equal(t, "test - "+n, sEnabled.events[0].name)
	assert.Equal(t, d, sEnabled.events[0].description)

	// Assert disabled
	assert.Equal(t, 0, len(sDisabled.events))

	// Add event 2
	n = "event 2"
	d = "description 2"
	sEnabled.AddEvent(n, d)
	sDisabled.AddEvent(n, d)

	// Assert enabled
	assert.Equal(t, 2, len(sEnabled.events))
	assert.Equal(t, "test - "+n, sEnabled.events[1].name)
	assert.Equal(t, d, sEnabled.events[1].description)

	// Assert disabled
	assert.Equal(t, 0, len(sDisabled.events))
}
