package application

import (
	"net/http"

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
	var newEvent data.Event
	if err := c.BindJSON(&newEvent); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad data schema passed"})
		return
	}
	controller.Storage.AddEvent(newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func (controller *Controller) AddEvents(c *gin.Context) {
	addEventsCounter.Inc()
	var newEvents data.MultipleEvents
	if err := c.BindJSON(&newEvents); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad data schema passed"})
		return
	}
	controller.Storage.AddEvents(newEvents.Events)
	c.IndentedJSON(http.StatusOK, newEvents)
}
