package api

import (
	"github.com/damianpugliese/go_rest_api_events_booking/cmd/api/routes"
	"github.com/gin-gonic/gin"
)

func Start() {
	server := gin.Default()

	routes.RegisterEventsRoutes(server)
	routes.RegisterUsersRoutes(server)

	server.Run(":8080")
}