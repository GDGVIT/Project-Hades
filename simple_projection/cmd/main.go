package main

import (
	"github.com/GDGVIT/Project-Hades/model"
	service "github.com/GDGVIT/Project-Hades/simple_projection/cmd/service"
)

func main() {
	conn := model.ConnectToDB()
	defer conn.Close()
	model.DBInit(conn)
	model.ConnectEnforcer()
	service.Run()
}
