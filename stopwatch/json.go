// Copyright 2015, Quentin RENARD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stopwatch

import "time"

// JSON represents a JSON-friendly structure of the Stopwatch results
type JSON struct {
	Events   []JSONEvent        `json:"events"`
	Timeline []JSONTimelineItem `json:"timeline"`
}

// JSONEvent represents a JSON-friendly structure of a Stopwatch event
type JSONEvent struct {
	HeapAlloc   float64 `json:"memory-peak(MB)"`
	Time        string  `json:"time"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

func newJSONEvent(e *event) JSONEvent {
	return JSONEvent{
		HeapAlloc:   e.heapAllocInMegabytes(),
		Time:        e.time.Format(time.RFC3339),
		Name:        e.name,
		Description: e.description,
	}
}

// JSONTimelineItem represents a JSON-friendly structure of a Stopwatch timeline item
type JSONTimelineItem struct {
	DeltaMemory float64   `json:"delta_memory(MB)"`
	DeltaTime   float64   `json:"delta_time(ms)"`
	EventStart  JSONEvent `json:"event_start"`
	EventStop   JSONEvent `json:"event_stop"`
}

func newJSONTimelineItem(eventStart *event, eventStop *event) JSONTimelineItem {
	return JSONTimelineItem{
		DeltaMemory: eventStop.heapAllocInMegabytes() - eventStart.heapAllocInMegabytes(),
		DeltaTime:   eventStop.timeInMilliseconds() - eventStart.timeInMilliseconds(),
		EventStart:  newJSONEvent(eventStart),
		EventStop:   newJSONEvent(eventStop),
	}
}
