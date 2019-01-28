package model

import (
	"fmt"
	"log"
	"os"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/joho/godotenv"
)

var con bolt.Conn

func DBInit(c bolt.Conn) {
	con = c
}

func ConnectToDB() bolt.Conn {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	URI := fmt.Sprintf("bolt://%s:%s@%s",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ENDPOINT"))

	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo(URI)
	if err != nil {
		log.Fatalln("Error connecting to DB")
	}
	return conn
}
