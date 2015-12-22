// Copyright 2015, Quentin RENARD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stopwatch

import (
	"fmt"
	"math"
	"strconv"
)

// Configuration represents the configuration necessary to initialize properly a stopwatch.
type Configuration struct {
	ID        string `json:"id"`
	IsEnabled bool   `json:"is_enabled"`
}

// Stopwatch represents a stopwatch capable of clocking events in your code and printing the results so that they can
// be analyzed easily
type Stopwatch interface {
	AddEvent(name string, description string) Stopwatch
	String() string
	JSON() JSON
	ID() string
	SetID(id string) Stopwatch
	IsEnabled() bool
	SetIsEnabled(isEnabled bool) Stopwatch
}

type stopwatch struct {
	id        string
	isEnabled bool
	events    []*event
}

// NewStopwatch creates a new Stopwatch based on its id
func NewStopwatch(id string) Stopwatch {
	return &stopwatch{
		id: id,
	}
}

// NewStopwatchFromConfiguration creates a new Stopwatch based on its configuration
func NewStopwatchFromConfiguration(c Configuration) Stopwatch {
	return NewStopwatch(c.ID).SetIsEnabled(c.IsEnabled)
}

// ID returns the stopwatch ID
func (s *stopwatch) ID() string {
	return s.id
}

// SetID sets the stopwatch ID
func (s *stopwatch) SetID(id string) Stopwatch {
	s.id = id
	return s
}

// IsEnabled returns whether the stopwatch is enabled
func (s *stopwatch) IsEnabled() bool {
	return s.isEnabled
}

// SetIsEnabled sets whether the stopwatch is enabled
func (s *stopwatch) SetIsEnabled(isEnabled bool) Stopwatch {
	s.isEnabled = isEnabled
	return s
}

// AddEvent adds a new event to the stopwatch based on a name and a description
func (s *stopwatch) AddEvent(name string, description string) Stopwatch {
	if s.isEnabled {
		s.events = append(s.events, newEvent(fmt.Sprintf("%s - %s", s.id, name), description))
	}
	return s
}

// String prints out the results of the stopwatch in a string
func (s *stopwatch) String() string {
	// Initialize
	var lastEvent event
	r := "Stopwatch results:\n"
	r += fmt.Sprintf("Id: %s\n", s.id)

	// Enabled
	if s.isEnabled {
		// Number of events
		r += fmt.Sprintf("Number of events: %d", len(s.events))

		// Events
		if len(s.events) > 0 {
			// Finalize introduction
			r += fmt.Sprintf("\nTime start: %s\n", s.events[0].time.Format("02 Jan 2006 15:04:05 -0700"))
			r += fmt.Sprintf("Memory start: %sMB", formatFloat64(s.events[0].heapAllocInMegabytes()))

			// Loop through events
			for i, e := range s.events {
				// Add deltas
				if i > 0 {
					// Initialize
					var signTime string
					var signMemory string

					// Compute delta time
					deltaTime := e.timeInMilliseconds() - lastEvent.timeInMilliseconds()
					if deltaTime > 0 {
						signTime = "+"
					} else {
						signTime = "-"
					}
					r += fmt.Sprintf("\n\n%s%sms\n", signTime, formatFloat64(math.Abs(deltaTime)))

					// Compute delta memory
					deltaMemory := s.events[0].heapAllocInMegabytes() - lastEvent.heapAllocInMegabytes()
					if deltaMemory >= 0 {
						signMemory = "+"
					} else {
						signMemory = "-"
					}
					r += fmt.Sprintf("%s%sMB", signMemory, formatFloat64(math.Abs(deltaMemory)))
				}

				// Add event
				r += fmt.Sprintf("\n\nName: %s", e.name)
				r += fmt.Sprintf("\nDescription: %s", e.description)

				// Update last event
				lastEvent = *e
			}
		}
	} else {
		// Stopwatch is disabled
		r += "Stopwatch is disabled"
	}

	// Return
	return r
}

func formatFloat64(i float64) string {
	return strconv.FormatFloat(i, 'f', 3, 64)
}

// JSON returns a JSON-friendly struct containing the stopwatch results
func (s *stopwatch) JSON() JSON {
	// Initialize
	json := JSON{
		Events:   []JSONEvent{},
		Timeline: []JSONTimelineItem{},
	}

	// Events
	if len(s.events) > 0 {
		// Get last event
		lastEvent := *s.events[0]
		lastEvent.name = ""
		lastEvent.description = ""

		// Loop through events
		for _, e := range s.events {
			// Add event
			json.Events = append(json.Events, newJSONEvent(e))

			// Add timeline event
			json.Timeline = append(json.Timeline, newJSONTimelineItem(&lastEvent, e))

			// Update last event
			lastEvent = *e
		}
	}

	// Return
	return json
}
