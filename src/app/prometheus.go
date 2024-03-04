package app

import (
	"github.com/prometheus/client_golang/prometheus"
)

func createCounter(name string, desc string) prometheus.Counter {
	return prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: name,
			Help: desc,
		},
	)
}

var HealthCounter = createCounter("health_ping_count", "No. of requests handled by health handler")
var AddEventCounter = createCounter("add_event_ping_count", "No. of requests handled by add event handler")
var AddEventsCounter = createCounter("add_events_ping_count", "No. of requests handled by add events handler")
var GetEventsCounter = createCounter("get_events_ping_count", "No. of requests handled by get events handler")
