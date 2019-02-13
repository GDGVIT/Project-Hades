package main

import (
	"log"
	"os"

	"github.com/GDGVIT/Project-Hades/exporter/endpoints/controller"
)

func main() {
	log.Println(os.Getenv("PROD_URI"))
	controller.Startup()
}
