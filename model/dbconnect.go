package model

import (
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

	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo(os.Getenv("URI"))
	if err != nil {
		log.Fatalln("Error connecting to DB")
	}
	return conn
}
