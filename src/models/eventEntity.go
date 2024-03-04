package models

type EventEntity struct {
	EventType string                 `json:"event_type"`
	Timestamp int64                  `json:"timestamp"`
	IP        string                 `json:"ip"`
	Data      map[string]interface{} `json:"data"`
}
