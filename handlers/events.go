package handlers

import (
	"net/http"
	"strconv"

	"github.com/damianpugliese/go_rest_api_events_booking/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the events", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

func GetEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not request data", "error": err.Error()})
		return
	}

	userId := c.GetInt64("userId")

	event.UserID = userId

	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func UpdateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	_, err = models.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event", "error": err.Error()})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not request data", "error": err.Error()})
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}

func DeleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event", "error": err.Error()})
		return
	}

	err = event.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}