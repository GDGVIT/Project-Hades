package service

import (
	"context"
	"log"
	"sync"

	"github.com/GDGVIT/Project-Hades/model"
)

// ParticipantsService describes the service.
type ParticipantsService interface {
	CreateAttendee(ctx context.Context, details model.Attendee) (rs string, err error)
	ReadAttendee(ctx context.Context, query model.Query) (rs []model.Participant, err error)
	UpdateAttendee(ctx context.Context, query model.Query) (rs string, err error)
	DeleteAttendee(ctx context.Context, query model.Query) (rs string, err error)
	DeleteAllAttendee(ctx context.Context, query model.Query) (rs string, err error)
}

type basicParticipantsService struct{}

/**
{
	"details":{
      "name":"angad sharma",
      "registrationNumber":"17BBE1010",
      "email":"SDADAS@A.COM",
      "phoneNumber":"919191991911",
      "gender":"M",
      "eventsAttended":"ALL",
      "eventName":"DEVSOC"
   }
}
**/
func (b *basicParticipantsService) CreateAttendee(ctx context.Context, details model.Attendee) (rs string, err error) {

	conn := model.ConnectToDB()
	defer conn.Close()

	c := make(chan error)
	var mutex = &sync.Mutex{}

	go model.CreateAttendee(details.EventName, model.Participant{
		Name:               details.Name,
		RegistrationNumber: details.RegistrationNumber,
		Email:              details.Email,
		PhoneNumber:        details.PhoneNumber,
		Gender:             details.Gender,
	}, c, mutex, conn)

	if err := <-c; err != nil {
		return "", err
	}

	return rs, err
}

/**
{
	"query":{
		"key":"name",
		"Value":"angad sharma"
	}
}

{
    "rs": [
        {
            "name": "angad sharma",
            "registrationNumber": "17BBE1010",
            "email": "SDADAS@A.COM",
            "phoneNumber": "919191991911",
            "gender": "M"
        }
    ],
    "err": null
}
**/
func (b *basicParticipantsService) ReadAttendee(ctx context.Context, query model.Query) (rs []model.Participant, err error) {

	conn := model.ConnectToDB()
	defer conn.Close()

	c := make(chan model.ParticipantReturn)
	var mutex = &sync.Mutex{}

	go model.ReadAttendee(query, c, mutex, conn)

	cb := <-c

	if err := cb.Err; err != nil {
		return cb.Attendees, err
	}

	return cb.Attendees, nil
}

func (b *basicParticipantsService) UpdateAttendee(ctx context.Context, query model.Query) (rs string, err error) {
	conn := model.ConnectToDB()
	defer conn.Close()

	c := make(chan error)

	go model.UpdateAttendee(query, c, conn)

	if err := <-c; err != nil {
		log.Println("Error updating attendees")
		return rs, err
	}

	return rs, nil
}
func (b *basicParticipantsService) DeleteAttendee(ctx context.Context, query model.Query) (rs string, err error) {
	conn := model.ConnectToDB()
	defer conn.Close()

	c := make(chan error)

	go model.DeleteAttendee(query, c, conn)

	if err := <-c; err != nil {
		log.Println("Error deleting attendees")
		return rs, err
	}

	return rs, nil
}

func (b *basicParticipantsService) DeleteAllAttendee(ctx context.Context, query model.Query) (rs string, err error) {
	// TODO implement the business logic of DeleteAllAttendee
	return rs, err
}

// NewBasicParticipantsService returns a naive, stateless implementation of ParticipantsService.
func NewBasicParticipantsService() ParticipantsService {
	return &basicParticipantsService{}
}

// New returns a ParticipantsService with all of the expected middleware wired in.
func New(middleware []Middleware) ParticipantsService {
	var svc ParticipantsService = NewBasicParticipantsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
