package application

import (
	"github.com/prometheus/client_golang/prometheus"
)

func CreateCounter(name string, desc string) prometheus.Counter {
	return prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: name,
			Help: desc,
		},
	)
}

var healthCounter = CreateCounter("health_ping_count", "No. of requests handled by health handler")
var addEventCounter = CreateCounter("add_event_ping_count", "No. of requests handled by add event handler")
var addEventsCounter = CreateCounter("add_events_ping_count", "No. of requests handled by add events handler")
var eventsCounter = CreateCounter("events_counter", "No. of events ")
