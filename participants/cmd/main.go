package main

import (
	"github.com/GDGVIT/Project-Hades/model"
	service "github.com/GDGVIT/Project-Hades/participants/cmd/service"
)

var X = 2

func main() {
	conn := model.ConnectToDB()
	defer conn.Close()
	model.DBInit(conn)
	service.Run()
}
