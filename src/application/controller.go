package application

import (
	"net/http"
	"time"

	"ese/server/data"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Storage *data.RedisClient
	Address string
}

func (controller *Controller) Health(c *gin.Context) {
	healthCounter.Inc()
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Server running @" + controller.Address})
}

func (usecases *Controller) GetEventLogs(c *gin.Context) {
	getEventsCounter.Inc()
	data := usecases.Storage.GetAllEvents()
	c.IndentedJSON(http.StatusOK, data)
}

func (controller *Controller) AddEvent(c *gin.Context) {
	addEventCounter.Inc()
	var newEventEntity data.EventEntity
	if err := c.BindJSON(&newEventEntity); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad data schema passed"})
		return
	}
	metaData := data.EventMetaData{
		ServerTimestamp: time.Now().Unix(),
	}
	event := data.Event{
		EventEntity:   newEventEntity,
		EventMetaData: metaData,
	}
	controller.Storage.AddEvent(event)
	c.IndentedJSON(http.StatusCreated, event)
}

func (controller *Controller) AddEvents(c *gin.Context) {
	addEventsCounter.Inc()
	var events []data.Event
	var newEventEntities data.MultipleEventEntities
	if err := c.BindJSON(&newEventEntities); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad data schema passed"})
		return
	}
	for _, eventEntity := range newEventEntities.EventEntity {
		metaData := data.EventMetaData{
			ServerTimestamp: time.Now().Unix(),
		}
		event := data.Event{
			EventEntity:   eventEntity,
			EventMetaData: metaData,
		}
		events = append(events, event)
	}

	controller.Storage.AddEvents(events)
	c.IndentedJSON(http.StatusOK, events)
}
