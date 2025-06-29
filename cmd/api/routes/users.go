package routes

import (
	"github.com/damianpugliese/go_rest_api_events_booking/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUsersRoutes(server *gin.Engine) {
	server.POST("/signup", handlers.Signup)
	server.POST("/login", handlers.Login)
}