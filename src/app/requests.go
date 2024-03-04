package app

import "ese.server/models"

type MultipleEventRequest struct {
	EventEntity []models.EventEntity `json:"events"`
}
