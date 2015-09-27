package stopwatch

import (
	"runtime"
	"time"
)

type Event struct {
	memStats    runtime.MemStats
	datetime    time.Time
	HeapAlloc   float64 `json:"memory-peak(MB)"`
	Time        string  `json:"microtime(ms)"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Children    []Event `json:"children,omitempty"`
}

func (oEvent *Event) New(sName string, sDescription string) {
	// Set memstats
	runtime.ReadMemStats(&oEvent.memStats)
	oEvent.HeapAlloc = float64(float64(oEvent.memStats.HeapAlloc) / float64(1000000))

	// Set datetime
	oEvent.datetime = time.Now()
	oEvent.Time = oEvent.datetime.UTC().Format("15:04:05")

	// Set name
	oEvent.Name = sName

	// Set description
	oEvent.Description = sDescription
}
