package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GDGVIT/Project-Hades/model"
	"github.com/GDGVIT/Project-Hades/organization/endpoints"
)

func main() {
	conn := model.ConnectToDB()
	defer conn.Close()
	model.DBInit(conn)
	model.ConnectEnforcer()
	mux := endpoints.Init()
	fmt.Println("Listening on port 8087....")
	log.Fatal(http.ListenAndServe(":8087", mux))
}
