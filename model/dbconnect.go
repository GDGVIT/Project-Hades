package model

import (
	"fmt"
	"log"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func ConnectToDB() bolt.Conn {

	URI := fmt.Sprintf("bolt://%s:%s@%s",
		DB_SECRET.DB_USERNAME, DB_SECRET.DB_PASSWORD,
		DB_SECRET.DB_ENDPOINT)

	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo(URI)
	if err != nil {
		log.Fatalln("Error connecting to DB")
	}
	return conn
}
