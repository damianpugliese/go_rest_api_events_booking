package main

import (
	"github.com/damianpugliese/go_rest_api_events_booking/db"
	"github.com/damianpugliese/go_rest_api_events_booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterEventsRoutes(server)
	routes.RegisterUsersRoutes(server)

	server.Run(":8080")
}