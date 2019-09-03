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
	ReadAttendee(ctx context.Context, query model.Query) (rs []model.Attendee, err error)
	RmAttendee(ctx context.Context, query model.Query) (rs string, err error)
	DeleteAttendee(ctx context.Context, query model.Query) (rs string, err error)
	DeleteAllAttendee(ctx context.Context, query model.Query) (rs string, err error)
}

type basicParticipantsService struct{}

/**
*@api {post} /api/v1/participants/create-attendee create an attendee
*@apiName create an attendee
*@apiGroup participants
*@apiPermission admin
*@apiParam {String} name name of the participant
*@apiParam {String} registrationNumber registration number of the participant
*@apiParam {String} email email of the participant
*@apiParam {String} phoneNumber phoneNumber of the participant
*@apiParam {String} gender gender of the participant
*@apiParam {String} eventName name of the event registering for
*
*@apiParamExample {json} request-example
*{
*	"details":{
*      "name":"angad sharma",
*      "registrationNumber":"17BBE1010",
*      "email":"SDADAS@A.COM",
*      "phoneNumber":"919191991911",
*      "gender":"M",
*      "eventsAttended":"ALL",
*      "eventName":"DEVSOC"
*   }
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": "created",
*    "err": null
*}
*
**/
func (b *basicParticipantsService) CreateAttendee(ctx context.Context, details model.Attendee) (rs string, err error) {

	c := make(chan error)
	var mutex = &sync.Mutex{}

	go model.CreateAttendee(details.EventName, model.Participant{
		Name:               details.Name,
		RegistrationNumber: details.RegistrationNumber,
		Email:              details.Email,
		PhoneNumber:        details.PhoneNumber,
		Gender:             details.Gender,
	}, c, mutex)

	if err := <-c; err != nil {
		return "some error occurred", err
	}

	return "created", nil
}

/**
*@api {get} /api/v1/participants/read-attendee read an attendee
*@apiName read an attendee
*@apiPermission super-admin
*@apiGroup participants
*@apiParam {String} key key to query the attendee by
*@apiParam {String} value value of the key
*
*@apiParamExample {json} request-example
*{
*	"query":{
*		"key":"name",
*		"Value":"angad sharma"
*	}
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": [
*        {
*            "name": "angad sharma",
*            "registrationNumber": "17BBE1010",
*            "email": "SDADAS@A.COM",
*            "phoneNumber": "919191991911",
*            "gender": "M",
			 "attended":"absent"
*        }
*    ],
*    "err": null
*}
**/
func (b *basicParticipantsService) ReadAttendee(ctx context.Context, query model.Query) (rs []model.Attendee, err error) {

	c := make(chan model.ParticipantReturn)
	var mutex = &sync.Mutex{}

	go model.ReadAttendee(query, c, mutex)

	cb := <-c

	if err := cb.Err; err != nil {
		return cb.Attendees, err
	}

	return cb.Attendees, nil
}

/**
*@api {delete} /api/v1/participants/rm-attendee remove attendee from an event
*@apiName remove attendee from an event
*@apiGroup participants
*@apiPermission admin
*@apiParam {String} key key to query the attendee by
*@apiParam {String} value value of the key
*@apiParam {String} changeKey Name of the club
*@apiParam {String} changeValue Name of the event
*
*@apiParamExample {json} request-example
*{
*	"query":{
*		"key":"name",
*		"Value":"dhruv sharma",
*		"changeKey":"DSC VIT",
*		"changeValue":"DEVFEST 2019"
*	}
*}
*
*@apiParamExample {json} response-example
*
*{
*    "rs": "updated",
*    "err": null
*}
**/
func (b *basicParticipantsService) RmAttendee(ctx context.Context, query model.Query) (rs string, err error) {
	c := make(chan error)

	go model.RmAttendee(query, c)

	if err := <-c; err != nil {
		log.Println("Error updating attendees")
		return "some error occurred", err
	}

	return "updated", nil
}

/**
*@api {delete} /api/v1/participants/delete-attendee delete an attendee
*@apiName delete an attendee
*@apiGroup participants
*@apiPermission super-admin
*@apiParam {String} key key to query the attendee by
*@apiParam {String} value value of the key
*
*@apiParamExample {json} request-example
*{
*	"query":{
*		"key":"name",
*		"Value":"dhruv sharma"
*	}
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": "deleted",
*    "err": null
*}
**/
func (b *basicParticipantsService) DeleteAttendee(ctx context.Context, query model.Query) (rs string, err error) {
	c := make(chan error)

	go model.DeleteAttendee(query, c)

	if err := <-c; err != nil {
		return "some error occurred", err
	}

	return "deleted", nil
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
