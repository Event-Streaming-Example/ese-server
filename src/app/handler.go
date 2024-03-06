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
	var newEventEntity models.EventEntity
	if err := c.BindJSON(&newEventEntity); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad data schema passed"})
		return
	}
	metaData := models.EventMetaData{
		ServerTimestamp: h.getCurrentTimeInMilli(),
	}
	event := models.Event{
		EventEntity:   newEventEntity,
		EventMetaData: metaData,
	}
	h.RedisClient.AddEvent(event)
	c.IndentedJSON(http.StatusCreated, event)
}

func (h *Handler) AddEvents(c *gin.Context) {
	AddEventsCounter.Inc()
	var events []models.Event
	var newEventEntities MultipleEventRequest
	if err := c.BindJSON(&newEventEntities); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad data schema passed"})
		return
	}
	for _, eventEntity := range newEventEntities.EventEntity {
		metaData := models.EventMetaData{
			ServerTimestamp: h.getCurrentTimeInMilli(),
		}
		event := models.Event{
			EventEntity:   eventEntity,
			EventMetaData: metaData,
		}
		events = append(events, event)
	}

	h.RedisClient.AddEvents(events)
	c.IndentedJSON(http.StatusCreated, events)
}

func (h *Handler) getCurrentTimeInMilli() int64 {
	return time.Now().UnixMilli()
}
