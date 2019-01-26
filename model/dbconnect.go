package model

import (
	"log"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func ConnectToDB(URI string) bolt.Conn {

	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo(URI)
	if err != nil {
		log.Fatalln("Error connecting to DB")
	}
	return conn
}
