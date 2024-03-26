package models

type Event struct {
	EventType       string                 `json:"type"`
	EventSubType    string                 `json:"sub_type"`
	Ip              string                 `json:"ip"`
	ClientTimestamp int64                  `json:"client_ts"`
	ServerTimestamp int64                  `json:"server_ts"`
	Data            map[string]interface{} `json:"data"`
}
