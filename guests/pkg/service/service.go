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

/**
*@api {post} /api/v1/guests/create-guest create a guest
*@apiName create a guest
*@apiGroup guest
*@apiPermission admin
*
*@apiParam {string} event name of the event
*@apiParam {string} name name of the guest
*@apiParam {string} email email of the guest
*@apiParam {string} phoneNumber phone number of the guest
*@apiParam {string} gender gender of the guest
*@apiParam {string} stake stake of the guest
*@apiParam {string} locationOfStay where does the guest stay? (origin)

*
*@apiParamExample {json} request-example
*
*{
*	"event":"DEVRELCONF",
*	"guest": {
*		"name":"angad",
*		"email":"sdaasd@a.com",
*		"phoneNumber":"9999999999",
*		"gender":"M",
*		"stake":"speaker",
*		"locationOfStay":"Bangalore"
*	}
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": "Guest created",
*    "err": null
*}
*
**/
func (b *basicGuestsService) CreateGuest(ctx context.Context, event string, guest model.Guest) (rs string, err error) {

	c := make(chan error)
	mutex := &sync.Mutex{}
	go model.CreateGuest(event, guest, c, mutex)
	if err := <-c; err != nil {
		return "Some error occurred while creating guest", err
	}
	return "Guest created", nil
}

/**
*@api {post} /api/v1/guests/read-guest read a guest
*@apiName read a guest
*@apiGroup guest
*@apiPermission admin
*
*@apiParam {string} key key of the event field
*@apiParam {string} value value of the event field
*
*@apiParamExample {json} request-example
*
*{
*	"query": {
*		"key":"name",
*		"value":"DEVRELCONF"
*	}
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": [
*        {
*            "name": "dafds",
*            "email": "sdadsaadasd@a.com",
*            "phoneNumber": "9699999999",
*            "gender": "M",
*            "stake": "speaker",
*            "locationOfStay": "Bangalore"
*        },
*        {
*            "name": "angad",
*            "email": "sdaasd@a.com",
*            "phoneNumber": "9999999999",
*            "gender": "M",
*            "stake": "speaker",
*            "locationOfStay": "Bangalore"
*        }
*    ],
*    "err": null
*}
*
**/
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

/**
*@api {post} /api/v1/guests/delete-guest delete a guest
*@apiName delete a guest
*@apiGroup guest
*@apiPermission admin
*
*@apiParam {string} key key of the event field
*@apiParam {string} value value of the event field
*@apiParam {string} changeKey Name of the guest to be deleted
*
*@apiParamExample {json} request-example
*
*{
*	"query": {
*		"key":"name",
*		"value":"DEVRELCONF",
*		"changeKey":"angad"
*	}
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": "Guest deleted",
*    "err": null
*}
*
**/
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
