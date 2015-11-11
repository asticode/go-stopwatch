// Copyright 2015, Quentin RENARD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stopwatch

import (
	"runtime"
	"time"
)

type event struct {
	memStats    runtime.MemStats
	time        time.Time
	name        string
	description string
}

func newEvent(name string, description string) *event {
	// Create event
	e := event{
		name:        name,
		description: description,
		time:        time.Now(),
	}

	// Set memstats
	runtime.ReadMemStats(&e.memStats)

	// Return
	return &e
}

func (e event) timeInMilliseconds() float64 {
	return convertNanosecondsToMilliseconds(e.time.Nanosecond())
}

func convertNanosecondsToMilliseconds(i int) float64 {
	return float64(i) / 1000000
}

func (e event) heapAllocInMegabytes() float64 {
	return convertBytesToMegabytes(e.memStats.HeapAlloc)
}

func convertBytesToMegabytes(i uint64) float64 {
	return float64(i) / float64(1024*1024)
}
