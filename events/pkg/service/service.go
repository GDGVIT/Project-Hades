package service

import (
	"context"

	"github.com/GDGVIT/Project-Hades/model"
)

// EventsService describes the service.
type EventsService interface {
	CreateEvent(ctx context.Context, event model.Event) (rs string, err error)
	ReadEvent(ctx context.Context, query model.Query) (rs model.Event, err error)
	UpdateEvent(ctx context.Context, query model.Query) (rs string, err error)
	DeleteEvent(ctx context.Context, query model.Query) (rs string, err error)
}

type basicEventsService struct{}

/**
* @api {post} /create-event create a new event
* @apiName create a new event
* @apiGroup admin
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
* @apiParam {Object} guest guest details (Guest object)
* @apiParam {String} PROrequest PRO department requests
* @apiParam {String} campusEngineerRequest Campus engineer requests
* @apiParam {String} duration duration of event
* @apiParam {Object} mainSponsor sponsor details (Participant Object)
*
* @apiParamExample {json} request-example
*
*"event":{
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
*     "name":"NOORU BAAP",
*     "registrationNumber":"17BBE1010",
*     "email":"SDADAS@A.COM",
*     "phoneNumber":"919191991911",
*     "gender":"M",
*     "eventsAttended":"ALL"
*  },
*  "guest":{
*     "name":"ALLAHH DAAS",
*     "email":"ASDSAD#ASD.COM",
*     "phoneNumber":"11111111111",
*     "gender":"F",
*     "stake":"SOME MONAYYYY",
*     "locationOfStay":"TERA GHAR"
*  },
*  "PROrequest":"SAJDOOSIJANDFSAKFDSAFD",
*  "campusEngineerRequest":"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD",
*  "duration":"16 hours",
*  "mainSponsor":{
*     "name":"ALLAHH DAAS",
*     "email":"ASDSAD#ASD.COM",
*     "phoneNumber":"11111111111",
*     "gender":"F",
*     "stake":"SOME MONAYYYY",
*     "locationOfStay":"TERA GHAR"
*  }
*
*
*
 */
func (b *basicEventsService) CreateEvent(ctx context.Context, event model.Event) (rs string, err error) {

	ce := make(chan error)
	go model.CreateEvent(event, ce)
	if err := <-ce; err != nil {
		return "", err
	}
	return rs, err
}

/**
*@api {post} /read-event read an event
*@apiName read an event
*@apiGroup admin
*@apiParam {String} key key to query the event by
*@apiParam {String} value value of the key
*
*@apiParamExample {json} request-example
*    "query":{
*		"key":"clubName",
*		"value":"GDG"
*	}
*
*@apiParamExample {json} response-example
*{
*    "rs": {
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
*            "name": "Murali S",
*            "registrationNumber": "",
*            "email": "SDADAS@A.COM",
*            "phoneNumber": "919191991911",
*            "gender": "M"
*        },
*        "studentCoordinator": {
*            "name": "Dhruv sharma",
*            "registrationNumber": "17BBE1010",
*            "email": "SDADAS@A.COM",
*            "phoneNumber": "919191991911",
*            "gender": "M"
*        },
*        "guest": {
*            "name": "angad sharma"",
*            "email": "ASDSAD#ASD.COM",
*            "phoneNumber": "11111111111",
*            "gender": "F",
*            "stake": "SOME MONAYYYY",
*            "locationOfStay": "VIT campus"
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
*        }
*    },
*    "err": null
*}
**/
func (b *basicEventsService) ReadEvent(ctx context.Context, query model.Query) (rs model.Event, err error) {

	ce := make(chan model.EventReturn)

	go model.ShowEventData(query, ce)

	cb := <-ce
	if cb.Err != nil {
		return cb.Event, cb.Err
	}
	return cb.Event, err
}

/**
*@api {post} /update-event update an event
*@apiName update an event
*@apiGroup admin
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
*	}
*}
*@apiParamExample {json} response-example
*{
*	rs:"",
*	err:null
*}
**/
func (b *basicEventsService) UpdateEvent(ctx context.Context, query model.Query) (rs string, err error) {

	ce := make(chan error)

	go model.UpdateEvent(query, ce)
	if err := <-ce; err != nil {
		return "", err
	}
	return rs, err
}

/**
*@api {post} /delete-event delete an event
*@apiName delete an event
*@apiGroup admin
*@apiParam {String} key key to query the event by
*@apiParam {String} value value of the key
*
*@apiParamExample {json} request-example
*{
*	"query":{
*		"key":"clubName",
*		"value":"GDG"
*	}
*}
*@apiParamExample {json} response-example
*{
*	 rs:"",
*	 err:null
* }
*
**/
func (b *basicEventsService) DeleteEvent(ctx context.Context, query model.Query) (rs string, err error) {

	ce := make(chan error)

	go model.DeleteEvent(query, ce)
	if err := <-ce; err != nil {
		return "", err
	}
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
