package models

type Event struct {
	EventEntity   EventEntity   `json:"entity"`
	EventMetaData EventMetaData `json:"meta_data"`
}
