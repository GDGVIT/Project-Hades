package main

import (
	service "github.com/GDGVIT/Project-Hades/events/cmd/service"
	"github.com/GDGVIT/Project-Hades/model"
)

func main() {
	conn := model.ConnectToDB()
	defer conn.Close()
	model.DBInit(conn)
	model.ConnectEnforcer()
	service.Run()
}
