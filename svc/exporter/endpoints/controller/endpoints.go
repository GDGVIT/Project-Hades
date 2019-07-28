package controller

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func RegisterRoutes() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/exporter/excel", HandleExcel)
	mux.HandleFunc("/api/v1/exporter/json", HandleJson)
	corsMux := cors.Default().Handler(mux)
	log.Println("Exporter listening on port 8804")
	log.Fatal(http.ListenAndServe(":8084", corsMux))

}
