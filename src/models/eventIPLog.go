package models

type EventIpLog struct {
	Ip        string  `json:"ip"`
	EventLogs []Event `json:"event_logs"`
}
