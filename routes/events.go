package routes

import (
	"github.com/damianpugliese/go_rest_api_events_booking/handlers"
	"github.com/damianpugliese/go_rest_api_events_booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterEventsRoutes(server *gin.Engine) {
	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:id", handlers.GetEvent)
	server.POST("/events", middlewares.Authenticate, handlers.CreateEvent)
	server.PUT("/events/:id", middlewares.Authenticate, handlers.UpdateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate, handlers.DeleteEvent)
}