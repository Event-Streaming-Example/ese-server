package models

type EventIPLog struct {
	IP        string  `json:"ip"`
	EventLogs []Event `json:"event_logs"`
}
