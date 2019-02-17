package service

import (
	"log"

	nats "github.com/nats-io/go-nats"
)

func publishEvent(event string, data []byte) {

	// connect to NATS
	natsConn, err := nats.Connect("nats:4222")
	if err != nil {
		log.Printf("Error connecting to NATS: %v", err)
	}
	defer natsConn.Close()

	// publish event
	err = natsConn.Publish(event, data)
	if err != nil {
		log.Printf("Error publishing to NATS: %v", err)
	}
	log.Printf("Event published with subject %s", event)

}
