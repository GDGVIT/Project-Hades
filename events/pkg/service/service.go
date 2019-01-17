package service

import (
	"context"

	model "github.com/angadsharma1016/omega_dbconfig"
)

// EventsService describes the service.
type EventsService interface {
	CreateEvent(ctx context.Context, s model.Event) (rs string, err error)
	ReadEvent(ctx context.Context, s model.Query) (rs string, err error)
	UpdateEvent(ctx context.Context, s model.Query) (rs string, err error)
	DeleteEvent(ctx context.Context, s model.Query) (rs string, err error)
}

type basicEventsService struct{}

func (b *basicEventsService) CreateEvent(ctx context.Context, s model.Event) (rs string, err error) {
	// TODO implement the business logic of CreateEvent
	return rs, err
}
func (b *basicEventsService) ReadEvent(ctx context.Context, s model.Query) (rs string, err error) {
	// TODO implement the business logic of ReadEvent
	return rs, err
}
func (b *basicEventsService) UpdateEvent(ctx context.Context, s model.Query) (rs string, err error) {
	// TODO implement the business logic of UpdateEvent
	return rs, err
}
func (b *basicEventsService) DeleteEvent(ctx context.Context, s model.Query) (rs string, err error) {
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
