package app

import (
	"net/http"
	"time"

	"ese.server/models"
	"ese.server/redis"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	RedisClient *redis.RedisClient
	Address     string
}

func (h *Handler) GetHealth(c *gin.Context) {
	HealthCounter.Inc()
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Server running @" + h.Address})
}

func (h *Handler) GetEvents(c *gin.Context) {
	GetEventsCounter.Inc()
	data := h.RedisClient.GetAllEvents()
	c.IndentedJSON(http.StatusOK, data)
}

func (h *Handler) AddEvent(c *gin.Context) {
	AddEventCounter.Inc()
	var newEventEntity models.PushEventRequest
	if err := c.BindJSON(&newEventEntity); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad data schema passed"})
		return
	}
	event := h.convertPushEventRequestToEvent(newEventEntity)
	h.RedisClient.AddEvent(event)
	c.IndentedJSON(http.StatusCreated, event)
}

func (h *Handler) AddEvents(c *gin.Context) {
	AddEventsCounter.Inc()
	var events []models.Event
	var newEventEntities models.PushMultipleEventRequest
	if err := c.BindJSON(&newEventEntities); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad data schema passed"})
		return
	}
	for _, eventEntity := range newEventEntities.EventEntity {
		event := h.convertPushEventRequestToEvent(eventEntity)
		events = append(events, event)
	}

	h.RedisClient.AddEvents(events)
	c.IndentedJSON(http.StatusCreated, events)
}

func (h *Handler) getCurrentTimeInMilli() int64 {
	return time.Now().UnixMilli()
}

func (h *Handler) convertPushEventRequestToEvent(pushEvent models.PushEventRequest) models.Event {
	return models.Event{
		EventType:       pushEvent.Event.Type,
		EventSubType:    pushEvent.Event.SubType,
		Ip:              pushEvent.IP,
		ClientTimestamp: pushEvent.Timestamp,
		ServerTimestamp: h.getCurrentTimeInMilli(),
		Data:            pushEvent.Data,
	}
}
