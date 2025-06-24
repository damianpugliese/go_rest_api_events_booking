package routes

import (
	"github.com/damianpugliese/go_rest_api_events_booking/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterEventsRoutes(server *gin.Engine) {
	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:id", handlers.GetEvent)
	server.POST("/events", handlers.CreateEvent)
	server.PUT("/events/:id", handlers.UpdateEvent)
	server.DELETE("/events/:id", handlers.DeleteEvent)
}