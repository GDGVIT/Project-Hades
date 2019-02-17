package main

import (
	"github.com/GDGVIT/Project-Hades/analytics/db"
	"github.com/GDGVIT/Project-Hades/analytics/endpoints"
)

func main() {
	var server *endpoints.Server

	// immutable event storage
	conn := db.DBconnect()
	defer conn.Close()

	server.Run()
}
