package endpoints

import (
	"log"
	"net/http"
	"time"

	"github.com/GDGVIT/Project-Hades/analytics/db"
	nats "github.com/nats-io/go-nats"
)

type Server struct {
}

func (s *Server) serve(port string) {
	log.Printf("Listening on port %s", port)
	http.ListenAndServe(port, nil)
}

func (s *Server) eventSubscribe() (*nats.Conn, error) {

	// connect to NATS
	natsConn, err := nats.Connect("nats:4222")
	if err != nil {
		log.Printf("Error connecting to NATS: %v", err)
		return nil, err
	}
	log.Println("Connected to NATS")

	// subscribe to all Hades events
	natsConn.Subscribe("hades.>", func(msg *nats.Msg) {
		log.Printf("Got a hit on %s", msg.Subject)
		go db.CreateLogs(msg.Subject, time.Now().String(), msg.Data)
	})
	return natsConn, nil
}

func (s *Server) Run() {
	http.HandleFunc("/api/v1/analytics", index())
	natsConn, _ := s.eventSubscribe()
	defer natsConn.Close()
	s.serve(":8085")
}
