package main

import (
	service "github.com/GDGVIT/Project-Hades/guests/cmd/service"
	"github.com/GDGVIT/Project-Hades/model"
)

func main() {
	conn := model.ConnectToDB()
	defer conn.Close()
	model.DBInit(conn)
	service.Run()
}
