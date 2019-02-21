package service

import (
	"context"
	"sync"

	"github.com/GDGVIT/Project-Hades/model"
)

// GuestsService describes the service.
type GuestsService interface {
	// Add your methods here
	CreateGuest(ctx context.Context, event string, guest model.Guest) (rs string, err error)
	ReadGuest(ctx context.Context, query model.Query) (rs []model.Guest, err error)
	UpdateGuest(ctx context.Context, query model.Query) (rs string, err error)
	DeleteGuest(ctx context.Context, query model.Query) (rs string, err error)

	CreateSponsor(ctx context.Context, event string, sponsor model.Participant) (rs string, err error)
	ReadSponsor(ctx context.Context, query model.Query) (rs []model.Participant, err error)
	UpdateSponsor(ctx context.Context, query model.Query) (rs string, err error)
	DeleteSponsor(ctx context.Context, query model.Query) (rs string, err error)
}

type basicGuestsService struct{}

func (b *basicGuestsService) CreateGuest(ctx context.Context, event string, guest model.Guest) (rs string, err error) {

	c := make(chan error)
	mutex := &sync.Mutex{}
	go model.CreateGuest(event, guest, c, mutex)
	if err := <-c; err != nil {
		return "Some error occurred while creating guest", err
	}
	return "Guest created", nil
}

func (b *basicGuestsService) ReadGuest(ctx context.Context, query model.Query) (rs []model.Guest, err error) {
	c := make(chan model.GuestReturn)
	go model.ReadGuest(query, c)

	data := <-c
	if err := data.Err; err != nil {
		return nil, err
	}
	return data.Guests, nil
}
func (b *basicGuestsService) UpdateGuest(ctx context.Context, query model.Query) (rs string, err error) {
	// TODO implement the business logic of UpdateGuest
	return rs, err
}
func (b *basicGuestsService) DeleteGuest(ctx context.Context, query model.Query) (rs string, err error) {
	c := make(chan error)
	go model.RmGuest(query, c)
	if err := <-c; err != nil {
		return "Error occurred while deleting guest", err
	}
	return "Guest deleted", nil
}
func (b *basicGuestsService) CreateSponsor(ctx context.Context, event string, sponsor model.Participant) (rs string, err error) {
	// TODO implement the business logic of CreateSponsor
	return rs, err
}
func (b *basicGuestsService) ReadSponsor(ctx context.Context, query model.Query) (rs []model.Participant, err error) {
	// TODO implement the business logic of ReadSponsor
	return rs, err
}
func (b *basicGuestsService) UpdateSponsor(ctx context.Context, query model.Query) (rs string, err error) {
	// TODO implement the business logic of UpdateSponsor
	return rs, err
}
func (b *basicGuestsService) DeleteSponsor(ctx context.Context, query model.Query) (rs string, err error) {
	// TODO implement the business logic of DeleteSponsor
	return rs, err
}

// NewBasicGuestsService returns a naive, stateless implementation of GuestsService.
func NewBasicGuestsService() GuestsService {
	return &basicGuestsService{}
}

// New returns a GuestsService with all of the expected middleware wired in.
func New(middleware []Middleware) GuestsService {
	var svc GuestsService = NewBasicGuestsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
