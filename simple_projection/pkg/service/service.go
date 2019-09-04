package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/GDGVIT/Project-Hades/model"
)

// SimpleProjectionService describes the service.
type SimpleProjectionService interface {
	ProjectAll(ctx context.Context, event string, query model.Query) (rs []model.Participant, err error)
	ProjectPresent(ctx context.Context, event string, day int, query model.Query) (rs []model.Participant, err error)
	ProjectAbsent(ctx context.Context, event string, day int, query model.Query) (rs []model.Participant, err error)
}

type basicSimpleProjectionService struct{}

/**
*@api {get} /api/v1/simple-projection/project-all show participants of an event
*@apiName show participants of an event
*@apiGroup simple_projection
*@apiPermission admin
*@apiParam {String} [key] key to query the event by
*@apiParam {String} [value] value of the key
*@apiParam {String} event event name
*
*@apiParamExample {json} request-example
*{
*	"event":"DEVFEST 2019",
*	"query":{
*		"key":"gender",
*		"value":"F",
* 	"specific" : "DSC-VIT"
*	}
*}
*@apiParamExample {json} response-example
*{
*    "rs": [
*        {
*            "name": "das",
*            "registrationNumber": "17BCE2009",
*            "email": "x@axd.com",
*            "phoneNumber": "919191991911",
*            "gender": "F"
*        }
*    ],
*    "err": null
*}
*
*@apiParamExample {json} request-example-all
*{
*	"event":"DEVFEST 2019",
* "query": {
*			"key": ""
*			"value": "",
*			"specific" :"DSCVIT"
*	}
*}
*
*@apiParamExample {json} response-example-all
*{
*    "rs": [
*        {
*            "name": "das",
*            "registrationNumber": "17BCE2009",
*            "email": "x@axd.com",
*            "phoneNumber": "919191991911",
*            "gender": "F"
*        },
*        {
*            "name": "das",
*            "registrationNumber": "17BCE2009",
*            "email": "x@ax.com",
*            "phoneNumber": "919191991911",
*            "gender": "M"
*        },
*        {
*            "name": "das",
*            "registrationNumber": "17BCE2009",
*            "email": "x@x.com",
*            "phoneNumber": "919191991911",
*            "gender": "M"
*        }
*    ],
*    "err": null
*}
**/
func (b *basicSimpleProjectionService) ProjectAll(ctx context.Context, event string, query model.Query) (rs []model.Participant, err error) {
	// authorize user
	token, err := model.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(query, event)
	if model.Enforce(token.Email, query.Specific, "member") == true || model.Enforce(token.Email, query.Specific, "admin") == true {

		c := make(chan model.SafeParticipantReturn)
		mutex := &sync.Mutex{}
		go model.ViewAll(event, query, c, mutex)

		msg := <-c
		if err := msg.Err; err != nil {
			return nil, err
		}
		return msg.Participants, msg.Err

	}

	return nil, errors.New("Error authorizing user")
}

/**
*@api {get} /api/v1/simple-projection/project-present show present participants
*@apiName show present participants
*@apiGroup simple_projection
*@apiPermission admin
*@apiParam {String} [key] key to query the event by
*@apiParam {String} [value] value of the key
*@apiParam {String} event event name
*@apiParam {String} day day of the event
*
*@apiParamExample {json} request-example
*{
*	"event":"DEVFEST 2019",
*	"day":2,
*	"query":{
*		"key":"gender",
*		"value":"F",
* 	"specific" : "DSC-VIT"
*	}
*}
*@apiParamExample {json} response-example
*{
*    "rs": [
*        {
*            "name": "das",
*            "registrationNumber": "17BCE2009",
*            "email": "x@axd.com",
*            "phoneNumber": "919191991911",
*            "gender": "F"
*        }
*    ],
*    "err": null
*}
*
*@apiParamExample {json} request-example-all
*{
*	"event":"DEVFEST 2019",
*	"day":2
*}
*
*@apiParamExample {json} response-example-all
*{
*    "rs": [
*        {
*            "name": "das",
*            "registrationNumber": "17BCE2009",
*            "email": "x@axd.com",
*            "phoneNumber": "919191991911",
*            "gender": "F"
*        },
*        {
*            "name": "das",
*            "registrationNumber": "17BCE2009",
*            "email": "x@ax.com",
*            "phoneNumber": "919191991911",
*            "gender": "M"
*        },
*        {
*            "name": "das",
*            "registrationNumber": "17BCE2009",
*            "email": "x@x.com",
*            "phoneNumber": "919191991911",
*            "gender": "M"
*        }
*    ],
*    "err": null
*}
**/
func (b *basicSimpleProjectionService) ProjectPresent(ctx context.Context, event string, day int, query model.Query) (rs []model.Participant, err error) {
	// authorize user
	token, err := model.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}
	if model.Enforce(token.Email, query.Specific, "member") == true || model.Enforce(token.Email, query.Specific, "admin") == true {

		c := make(chan model.SafeParticipantReturn)
		mutex := &sync.Mutex{}
		go model.ViewPresent(event, query, day, c, mutex)

		msg := <-c
		if err := msg.Err; err != nil {
			return nil, err
		}
		return msg.Participants, msg.Err
		return rs, err

	}

	return nil, errors.New("Error authorizing user")
}

/**
*@api {get} /api/v1/simple-projection/project-absent show absent participants
*@apiName show absent participants
*@apiGroup simple_projection
*@apiPermission admin
*@apiParam {String} [key] key to query the event by
*@apiParam {String} [value] value of the key
*@apiParam {String} event event name
*@apiParam {String} day day of the event
*
*@apiParamExample {json} request-example
*{
*	"event":"DEVFEST 2019",
*	"day":2,
*	"query":{
*		"key":"gender",
*		"value":"F",
* 	"specific" : "DSC-VIT"
*	}
*}
*@apiParamExample {json} response-example
*{
*    "rs": [
*        {
*            "name": "das",
*            "registrationNumber": "17BCE2009",
*            "email": "x@axd.com",
*            "phoneNumber": "919191991911",
*            "gender": "F"
*        }
*    ],
*    "err": null
*}
*
*}
**/
func (b *basicSimpleProjectionService) ProjectAbsent(ctx context.Context, event string, day int, query model.Query) (rs []model.Participant, err error) {

	// authorize user
	token, err := model.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}
	if model.Enforce(token.Email, query.Specific, "member") == true || model.Enforce(token.Email, query.Specific, "admin") == true {

		c := make(chan model.SafeParticipantReturn)
		go model.ViewAbsent(event, query, day, c)

		msg := <-c
		if err := msg.Err; err != nil {
			return nil, err
		}
		return msg.Participants, msg.Err
		return rs, err

	}

	return nil, errors.New("Error authorizing user")
}

// NewBasicSimpleProjectionService returns a naive, stateless implementation of SimpleProjectionService.
func NewBasicSimpleProjectionService() SimpleProjectionService {
	return &basicSimpleProjectionService{}
}

// New returns a SimpleProjectionService with all of the expected middleware wired in.
func New(middleware []Middleware) SimpleProjectionService {
	var svc SimpleProjectionService = NewBasicSimpleProjectionService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
