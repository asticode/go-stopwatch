package stopwatch

import (
	"errors"
	"fmt"
	"github.com/asticode/go-toolbox/network"
	"net/http"
)

type Stopwatch struct {
	Configuration Configuration `json:"-"`
	id            string
	enabler       Enabler
	events        []Event
}

type Json struct {
	Events   []Event        `json:"events"`
	Timeline []TimelineItem `json:"timeline"`
}

type TimelineItem struct {
	Delta      float64 `json:"delta(ms)"`
	EventStart Event   `json:"event_start"`
	EventStop  Event   `json:"event_stop"`
}

func NewStopwatch(oConfiguration Configuration) (*Stopwatch, error) {
	oStopwatch := Stopwatch{
		Configuration: oConfiguration,
	}
	oErr := oStopwatch.LoadConfiguration()
	return &oStopwatch, oErr
}

func (oStopwatch *Stopwatch) LoadConfiguration() error {
	// Set prefix
	if oStopwatch.Configuration.Id == "" {
		return errors.New("Stopwatch Id is required")
	}
	oStopwatch.id = oStopwatch.Configuration.Id

	// Set enabler
	oStopwatch.enabler = oStopwatch.Configuration.Enabler

	// Return
	return nil
}

func (oStopwatch *Stopwatch) Push(sName string, sDescription string) {
	// Create event
	oEvent := Event{}
	oEvent.New(oStopwatch.addIdToName(sName), sDescription)

	// Add event
	oStopwatch.events = append(oStopwatch.events, oEvent)
}

func (oStopwatch Stopwatch) addIdToName(sName string) string {
	if oStopwatch.id != "" {
		return fmt.Sprintf("%s: %s", oStopwatch.id, sName)
	}
	return sName
}

func (oStopwatch Stopwatch) IsEnabled(sIpAddress string, oHeader http.Header) bool {
	// Check ip addresses
	if len(oStopwatch.enabler.IpAddresses) > 0 && !network.ValidateIPAddress(sIpAddress, oStopwatch.enabler.IpAddresses) {
		return false
	}

	// Check headers
	for sHeaderName, sHeaderValue := range oStopwatch.enabler.Headers {
		if oHeader.Get(sHeaderName) != sHeaderValue {
			return false
		}
	}

	// Return true by default
	return true
}

func (oStopwatch Stopwatch) Json() Json {
	// Initialize.
	oJson := Json{
		Events:   oStopwatch.events,
		Timeline: []TimelineItem{},
	}
	oEventLast := Event{}
	iDelta := float64(-1)

	// Loop through events.
	for _, oEvent := range oStopwatch.events {
		// Initialize
		oTimelineItem := TimelineItem{}

		// Set delta
		if iDelta == -1 {
			iDelta = 0
		} else {
			iDelta = float64(float64(oEvent.datetime.Sub(oEventLast.datetime).Nanoseconds()) / float64(1000000))
		}

		// Set timeline item
		oTimelineItem.Delta = iDelta
		oTimelineItem.EventStart = oEventLast
		oTimelineItem.EventStop = oEvent

		// Add to timeline
		oJson.Timeline = append(oJson.Timeline, oTimelineItem)

		// Set event last
		oEventLast = oEvent
	}

	// Return
	return oJson
}
