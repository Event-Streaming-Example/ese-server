package data

type Event struct {
	EventType string                 `json:"eventType"`
	Timestamp int64                  `json:"timestamp"`
	IP        string                 `json:"ip"`
	Data      map[string]interface{} `json:"data"`
}

type MultipleEvents struct {
	Events []Event `json:"events"`
}

type EventIPLog struct {
	IP        string  `json:"ip"`
	EventLogs []Event `json:"eventLogs"`
}
