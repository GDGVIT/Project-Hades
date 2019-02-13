package controller

import (
	"log"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/api/v1/exporter/excel", HandleExcel)
	http.HandleFunc("/api/v1/exporter/json", HandleJson)
	log.Println("Exporter listening on port 8804")
	log.Fatal(http.ListenAndServe(":8084", nil))

}
