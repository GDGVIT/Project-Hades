package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/GDGVIT/Project-Hades/model"
)

// EventsService describes the service.
type EventsService interface {
	CreateEvent(ctx context.Context, event model.Event) (rs string, err error)
	ReadEvent(ctx context.Context, query model.Query) (rs []model.Event, err error)
	UpdateEvent(ctx context.Context, query model.Query) (rs string, err error)
	DeleteEvent(ctx context.Context, query model.Query) (rs string, err error)
}

type basicEventsService struct{}

/**
* @api {post} /api/v1/event/create-event create a new event
* @apiName create a new event
* @apiGroup events
* @apiPermission admin
* @apiSampleRequest http://localhost:8800/
*
* @apiParam {String} clubName Name of your club
* @apiParam {String} name Name of your event
* @apiParam {String} toDate ending date
* @apiParam {String} fromDate start date
* @apiParam {String} toTime start time
* @apiParam {String} fromTime ending time
* @apiParam {String} budget budget
* @apiParam {String} description event description
* @apiParam {String} [category] category of the event
* @apiParam {String} venue event venue
* @apiParam {String} [attendance] Name of your club
* @apiParam {String} expectedParticipants Expected Participants in the event
* @apiParam {Object} facultyCoordinator faculty coordinator details (Participant Object)
* @apiParam {Object} studentCoordinator student coordinator details (Participant Object)
* @apiParam {String} PROrequest PRO department requests
* @apiParam {String} campusEngineerRequest Campus engineer requests
* @apiParam {String} duration duration of event
*
* @apiParamExample {json} request-example
*
*{"event":{
*  "clubName":"GDG",
*  "name":"DEVRELCONF",
*  "toDate":"10TH OCTOBER",
*  "fromDate":"8TH OCTOBER",
*  "toTime":"10 PM",
*  "fromTime":"11 AM",
*  "budget":"200000",
*  "description":"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING",
*  "category":"TECHNICAL",
*  "venue":"ANNA AUDI",
*  "attendance":"4000",
*  "expectedParticipants":"4000",
*  "facultyCoordinator":{
*     "name":"NOORU MAA",
*     "registrationNumber":"17BBE1010",
*     "email":"SDADAS@A.COM",
*     "phoneNumber":"919191991911",
*     "gender":"M",
*     "eventsAttended":"ALL"
*  },
*  "studentCoordinator":{
*     "name":"NOOR",
*     "registrationNumber":"17BBE1010",
*     "email":"SDADAS@A.COM",
*     "phoneNumber":"919191991911",
*     "gender":"M",
*     "eventsAttended":"ALL"
*  },
*  "PROrequest":"SAJDOOSIJANDFSAKFDSAFD",
*  "campusEngineerRequest":"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD",
*  "duration":"16 hours"
*}
*}
*
*@apiParamExample {json} response-example
*{
*	rs:"created",
*	err:null
*}
*
 */
func (b *basicEventsService) CreateEvent(ctx context.Context, event model.Event) (rs string, err error) {

	// authorize user
	token, err := model.VerifyToken(ctx)
	if err != nil {
		fmt.Println("Hello world")
		return "Error authorizing user", err
	}
	if !model.Enforce(token.Email, event.ClubName, "member") && !model.Enforce(token.Email, event.ClubName, "admin") {
		return "Error authorizing user", nil
	}
	ce := make(chan error)
	go model.CreateEvent(event, ce)
	if err := <-ce; err != nil {
		return "some error occurred", err
	}

	data, err := json.Marshal(event)
	if err != nil {
		return "error occurred while unmarshaling json", err
	}
	go publishEvent("hades.event.CreateEvent", data)
	return "created", err
}

/**
*@api {get} /api/v1/event/read-event read an event
*@apiName read an event
*@apiGroup events
*@apiPermission admin
*@apiParam {String} key key to query the event by
*@apiParam {String} value value of the key
*@apiParam {String} [specific] search by name of the event
*@apiParamExample {json} request-example
*    {"query":{
*		"key":"clubName",
*		"value":"GDG",
*		"specific":"DEVFEST 2019"
*		"organization":"CodeChef-VIT"
*	}}
*
*@apiParamExample {json} response-example
*{
*    "rs": [{
*        "clubName": "GDG",
*        "name": "DEVRELCONF",
*        "toDate": "10TH OCTOBER",
*        "fromDate": "8TH OCTOBER",
*        "toTime": "10 PM",
*        "fromTime": "11 AM",
*        "budget": "200000",
*        "description": "TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING",
*        "category": "TECHNICAL",
*        "venue": "ANNA AUDI",
*        "attendance": "4000",
*        "expectedParticipants": "4000",
*        "facultyCoordinator": {
*            "name": "NOORU MAA",
*            "registrationNumber": "17BBE1010",
*            "email": "SDADAS@A.COM",
*            "phoneNumber": "919191991911",
*            "gender": "M"
*        },
*        "studentCoordinator": {
*            "name": "NOORU BAAP",
*            "registrationNumber": "17BBE1010",
*            "email": "SDADAS@A.COM",
*            "phoneNumber": "919191991911",
*            "gender": "M"
*        },
*        "PROrequest": "SAJDOOSIJANDFSAKFDSAFD",
*        "campusEngineerRequest": "SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD",
*        "duration": "16 hours",
*        "mainSponsor": {
*            "name": "",
*            "registrationNumber": "",
*            "email": "",
*            "phoneNumber": "",
*            "gender": ""
*        },
*        "status": "false"
*
*    }],
*    "err": null
*}
**/
func (b *basicEventsService) ReadEvent(ctx context.Context, query model.Query) (rs []model.Event, err error) {

	// authorize user
	token, err := model.VerifyToken(ctx)
	if err != nil {
		fmt.Println("Hello world")
		return nil, err
	}
	if !model.Enforce(token.Email, query.Organization, "member") && !model.Enforce(token.Email, query.Organization, "admin") {
		return nil, nil
	}
	ce := make(chan model.EventReturn)

	go model.ShowEventData(query, ce)

	cb := <-ce
	if cb.Err != nil {
		return cb.Event, cb.Err
	}
	return cb.Event, err
}

/**
*@api {put} /api/v1/event/update-event update an event
*@apiName update an event
*@apiGroup events
*@apiPermission admin
*@apiParam {String} key key to query the event by
*@apiParam {String} value value of the key
*@apiParam {String} changeKey key of the value which needs to be altered
*@apiParam {String} changeValue the new value
*
*@apiParamExample {json} request-example
*{
*	"query":{
*		"key":"clubName",
*		"value":"GDG",
*		"changeKey":"clubName",
*		"changeValue":"codechef"
*		"organization":"CodeChef-VIT"
*	}
*}
*@apiParamExample {json} response-example
*{
*	rs:"updated",
*	err:null
*}
**/
func (b *basicEventsService) UpdateEvent(ctx context.Context, query model.Query) (rs string, err error) {

	// authorize user
	token, err := model.VerifyToken(ctx)
	if err != nil {
		fmt.Println("Hello world")
		return "Error authorizing user", err
	}
	if !model.Enforce(token.Email, query.Organization, "member") && !model.Enforce(token.Email, query.Organization, "admin") {
		return "Error authorizing user", nil
	}

	ce := make(chan error)

	go model.UpdateEvent(query, ce)
	if err := <-ce; err != nil {
		return "some error occurred", err
	}
	return "updated", err
}

/**
*@api {delete} /api/v1/event/delete-event delete an event
*@apiName delete an event
*@apiGroup events
*@apiPermission admin
*@apiParam {String} key key to query the event by
*@apiParam {String} value value of the key
*
*@apiParamExample {json} request-example
*{
*	"query":{
*		"key":"clubName",
*		"value":"GDG"
*		"organization":"CodeChef-VIT"
*	}
*}
*@apiParamExample {json} response-example
*{
*	 rs:"deleted",
*	 err:null
* }
*
**/
func (b *basicEventsService) DeleteEvent(ctx context.Context, query model.Query) (rs string, err error) {
	// authorize user
	token, err := model.VerifyToken(ctx)
	if err != nil {
		fmt.Println("Hello world")
		return "Error authorizing user", err
	}

	if !model.Enforce(token.Email, query.Organization, "member") && !model.Enforce(token.Email, query.Organization, "admin") {
		return "Error authorizing user", nil
	}

	ce := make(chan error)

	go model.DeleteEvent(query, ce)
	if err := <-ce; err != nil {
		return "some error occurred", err
	}
	return "deleted", err
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
