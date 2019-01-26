package service

import (
	"context"
	"fmt"

	"github.com/GDGVIT/Project-Hades/model"
)

// EventsService describes the service.
type EventsService interface {
	// Add your methods here
	CreateEvent(ctx context.Context, event model.Event) (rs string, err error)
	ReadEvent(ctx context.Context, query model.Query) (rs model.Event, err error)
	UpdateEvent(ctx context.Context, query model.Query) (rs string, err error)
	DeleteEvent(ctx context.Context, query model.Query) (rs string, err error)
}

type basicEventsService struct{}

func (b *basicEventsService) CreateEvent(ctx context.Context, event model.Event) (rs string, err error) {

	// create connection to DB
	conn := model.ConnectToDB(fmt.Sprintf("bolt://%s:%s@%s",
		model.DB_SECRET.DB_USERNAME, model.DB_SECRET.DB_PASSWORD,
		model.DB_SECRET.DB_ENDPOINT)) //("bolt://username:password@localhost:7687"
	fmt.Println(conn)

	defer conn.Close()

	ce := make(chan error)
	go model.CreateEvent(event, ce, conn)
	if err := <-ce; err != nil {
		return "", err
	}
	return rs, err
}

func (b *basicEventsService) ReadEvent(ctx context.Context, query model.Query) (rs model.Event, err error) {
	// TODO implement the business logic of ReadEvent
	return rs, err
}
func (b *basicEventsService) UpdateEvent(ctx context.Context, query model.Query) (rs string, err error) {
	// TODO implement the business logic of UpdateEvent
	return rs, err
}
func (b *basicEventsService) DeleteEvent(ctx context.Context, query model.Query) (rs string, err error) {
	// TODO implement the business logic of DeleteEvent
	return rs, err
}

// NewBasicEventsService returns a naive, stateless implementation of EventsService.
func NewBasicEventsService() EventsService {
	return &basicEventsService{}
}

// New returns a EventsService with all of the expected middleware wired in.
func New(middleware []Middleware) EventsService {
	var svc EventsService = NewBasicEventsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
