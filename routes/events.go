package routes

import (
	"github.com/damianpugliese/go_rest_api_events_booking/handlers"
	"github.com/damianpugliese/go_rest_api_events_booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterEventsRoutes(server *gin.Engine) {
	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:id", handlers.GetEvent)

	protected := server.Group("/")
	protected.Use(middlewares.Authenticate)
	protected.POST("/events", handlers.CreateEvent)
	protected.PUT("/events/:id", handlers.UpdateEvent)
	protected.DELETE("/events/:id", handlers.DeleteEvent)
	protected.POST("/events/:id/register", handlers.RegisterForEvent)
	protected.DELETE("/events/:id/register", handlers.CancelRegistration)
}