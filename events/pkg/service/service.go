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

/*
{
    "rs": {
        "clubName": "GDG",
        "name": "DEVRELCONF",
        "toDate": "10TH OCTOBER",
        "fromDate": "8TH OCTOBER",
        "toTime": "10 PM",
        "fromTime": "11 AM",
        "budget": "200000",
        "description": "TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING",
        "category": "TECHNICAL",
        "venue": "ANNA AUDI",
        "attendance": "4000",
        "expectedParticipants": "4000",
        "facultyCoordinator": {
            "name": "Murali S",
            "registrationNumber": "",
            "email": "SDADAS@A.COM",
            "phoneNumber": "919191991911",
            "gender": "M"
        },
        "studentCoordinator": {
            "name": "Dhruv sharma",
            "registrationNumber": "17BBE1010",
            "email": "SDADAS@A.COM",
            "phoneNumber": "919191991911",
            "gender": "M"
        },
        "guest": {
            "name": "angad sharma"",
            "email": "ASDSAD#ASD.COM",
            "phoneNumber": "11111111111",
            "gender": "F",
            "stake": "SOME MONAYYYY",
            "locationOfStay": "VIT campus"
        },
        "PROrequest": "SAJDOOSIJANDFSAKFDSAFD",
        "campusEngineerRequest": "SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD",
        "duration": "16 hours",
        "mainSponsor": {
            "name": "",
            "registrationNumber": "",
            "email": "",
            "phoneNumber": "",
            "gender": ""
        }
    },
    "err": null
}
*/
func (b *basicEventsService) ReadEvent(ctx context.Context, query model.Query) (rs model.Event, err error) {
	// create connection to DB
	conn := model.ConnectToDB(fmt.Sprintf("bolt://%s:%s@%s",
		model.DB_SECRET.DB_USERNAME, model.DB_SECRET.DB_PASSWORD,
		model.DB_SECRET.DB_ENDPOINT)) //("bolt://username:password@localhost:7687"
	fmt.Println(conn)

	defer conn.Close()

	ce := make(chan model.EventReturn)

	go model.ShowEventData(query, ce, conn)

	cb := <-ce
	if cb.Err != nil {
		return cb.Event, cb.Err
	}
	return cb.Event, err
}

/*
{
	"query":{
		"key":"clubName",
		"value":"GDG",
		"changeKey":"clubName",
		"changeValue":"codechef"
	}
}
*/
func (b *basicEventsService) UpdateEvent(ctx context.Context, query model.Query) (rs string, err error) {
	// create connection to DB
	conn := model.ConnectToDB(fmt.Sprintf("bolt://%s:%s@%s",
		model.DB_SECRET.DB_USERNAME, model.DB_SECRET.DB_PASSWORD,
		model.DB_SECRET.DB_ENDPOINT)) //("bolt://username:password@localhost:7687"
	fmt.Println(conn)

	defer conn.Close()

	ce := make(chan error)

	go model.UpdateEvent(query, ce, conn)
	if err := <-ce; err != nil {
		return "", err
	}
	return rs, err
}

/*
{
	"query":{
		"key":"clubName",
		"value":"GDG"
	}
}
*/
func (b *basicEventsService) DeleteEvent(ctx context.Context, query model.Query) (rs string, err error) {
	// create connection to DB
	conn := model.ConnectToDB(fmt.Sprintf("bolt://%s:%s@%s",
		model.DB_SECRET.DB_USERNAME, model.DB_SECRET.DB_PASSWORD,
		model.DB_SECRET.DB_ENDPOINT)) //("bolt://username:password@localhost:7687"
	fmt.Println(conn)

	defer conn.Close()

	ce := make(chan error)

	go model.DeleteEvent(query, ce, conn)
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
