package application

import (
	"net/http"

	"ese/server/data"

	"github.com/gin-gonic/gin"
)

type Usecases struct {
	Storage *data.RedisClient
}

func (controller *Usecases) Health(c *gin.Context) {
	healthCounter.Inc()
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Server running"})
}

func (usecases *Usecases) GetEventLogs(c *gin.Context) {
	data := usecases.Storage.GetAllEvents()
	c.IndentedJSON(http.StatusOK, data)
}

func (controller *Usecases) AddEvent(c *gin.Context) {
	var newEvent data.Event
	if err := c.BindJSON(&newEvent); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad data schema passed"})
		return
	}
	increamentEventCounterIfExists(newEvent, controller.Storage)
	controller.Storage.AddEvent(newEvent)
	addEventCounter.Inc()
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func (controller *Usecases) AddEvents(c *gin.Context) {
	var newEvents data.MultipleEvents
	if err := c.BindJSON(&newEvents); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad data schema passed"})
		return
	}
	for _, event := range newEvents.Events {
		increamentEventCounterIfExists(event, controller.Storage)
	}
	controller.Storage.AddEvents(newEvents.Events)
	addEventsCounter.Inc()
	c.IndentedJSON(http.StatusOK, newEvents)
}

func increamentEventCounterIfExists(event data.Event, storage *data.RedisClient) {
	if storage.EventExists(event.IP) == 0 {
		eventsCounter.Inc()
	}
}
