package main

import (
	"github.com/damianpugliese/go_rest_api_events_booking/cmd/api"
	"github.com/damianpugliese/go_rest_api_events_booking/internal/infrastructure/db"
)

func main() {
	db.InitDB()

	api.Start()
}