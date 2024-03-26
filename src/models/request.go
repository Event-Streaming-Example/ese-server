package models

type PushEvent struct {
	Type    string `json:"type"`
	SubType string `json:"sub_type"`
}

type PushEventRequest struct {
	Event     PushEvent              `json:"event"`
	Timestamp int64                  `json:"timestamp"`
	IP        string                 `json:"ip"`
	Data      map[string]interface{} `json:"data"`
}

type PushMultipleEventRequest struct {
	EventEntity []PushEventRequest `json:"events"`
}
