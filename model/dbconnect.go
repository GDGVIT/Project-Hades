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

	// cmd := exec.Command("sleep", "10")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// err = cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }

	conn, err := bolt.NewDriver().OpenNeo(os.Getenv("PROD_URI"))
	if err != nil {
		log.Fatalln("Error connecting to DB")
	}
	return conn
}
