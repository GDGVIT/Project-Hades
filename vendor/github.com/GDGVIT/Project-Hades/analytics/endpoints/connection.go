package endpoints

import (
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"

	db "github.com/GDGVIT/Project-Hades/analytics/modelfuncs"
	nats "github.com/nats-io/go-nats"
)

type Server struct {
}

func (s *Server) serve(port string, mux *http.Handler) {
	log.Printf("Listening on port %s", port)
	http.ListenAndServe(port, *mux)
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

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/analytics", readFromDB())
	mux.HandleFunc("/api/v1/analytics/all", readAllFromDB())

	CORSmux := cors.Default().Handler(mux)
	natsConn, _ := s.eventSubscribe()
	defer natsConn.Close()
	s.serve(":8085", &CORSmux)
}
