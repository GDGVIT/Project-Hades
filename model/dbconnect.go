package model

import (
	"log"
	"os"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

var con bolt.Conn

func DBInit(c bolt.Conn) {
	con = c
}

func ConnectToDB() bolt.Conn {

	conn, err := bolt.NewDriver().OpenNeo(os.Getenv("PROD_URI"))
	if err != nil {
		log.Fatalln("Error connecting to DB")
	}
	return conn
}
