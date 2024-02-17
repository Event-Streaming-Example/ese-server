package data

type EventEntity struct {
	EventType string                 `json:"event_type"`
	Timestamp int64                  `json:"timestamp"`
	IP        string                 `json:"ip"`
	Data      map[string]interface{} `json:"data"`
}

type EventMetaData struct {
	ServerTimestamp int64 `json:"server_timestamp"`
}

type Event struct {
	EventEntity   EventEntity   `json:"entity"`
	EventMetaData EventMetaData `json:"meta_data"`
}

type EventIPLog struct {
	IP        string  `json:"ip"`
	EventLogs []Event `json:"event_logs"`
}
