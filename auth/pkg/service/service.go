package service

import (
	"context"

	"github.com/GDGVIT/Project-Hades/model"
)

// AuthService describes the service.
type AuthService interface {
	// Add your methods here
	Login(ctx context.Context, email string, password string) (rs string, token string, err error)
	Signup(ctx context.Context, user model.User) (rs string, token string, err error)
	CreateOrg(ctx context.Context, data model.Organization) (rs string, err error)
	LoginOrg(ctx context.Context, data model.Organization) (rs string, err error)
	Invite(ctx context.Context, email string, org string) (rs string, err error)
	ShowInvites(ctx context.Context) (org string, admin string, err error)
	ShowProfile(ctx context.Context) (orgs []model.Organization, user []model.User, events []model.Event, err error)
}

type basicAuthService struct{}

/**
* @api {post} /api/v1/auth/login login as a user
* @apiName login as a user
* @apiGroup auth
*
* @apiParam {string} password password of the user
* @apiParam {string} email email of the user
*
*
* @apiParamExample {json} request-example
*{
*	"email":"test1@test.com",
*	"password":"test"
*}
*
* @apiParamExample {json} response-example
*{
*    "rs": "Done",
*    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQHRlc3QuY29tIiwicm9sZSI6IkRFRkFVTFQiLCJvcmdhbml6YXRpb24iOiIifQ.3Qj3eu8iwXL2v1Rb7qGEf5USQ-WVjRvYiLALWIbWZ5c",
*    "err": null
*}
**/
func (b *basicAuthService) Login(ctx context.Context, email string, password string) (rs string, token string, err error) {

	token, err = model.Login(email, password, "DEFAULT", "")

	if err != nil {
		return "Some error occurred", token, err
	}
	return "Done", token, nil
}

/**
* @api {post} /api/v1/auth/signup signup as a user
* @apiName signup as a user
* @apiGroup auth
*
* @apiParam {string} firstName first name of the user
* @apiParam {string} lastName last name of the user
* @apiParam {string} password password of the user
* @apiParam {string} email email of the user
* @apiParam {string} phoneNumber phoneNumber of the user
* @apiParam {string} linkedIn linkedIn URL of the user
* @apiParam {string} facebook facebook URL of the user
* @apiParam {string} linkedIn linkedIn URL of the user
* @apiParam {string} description description of the user
* @apiParam {string} createdAt when was the user created
*
*
* @apiParamExample {json} request-example
*
* {
* 	"user" : {
* 	"firstName": "test",
* 	"lastName": "test",
* 	"password": "test",
* 	"email": "test1@test.com",
* 	"phoneNumber": "998171818",
* 	"linkedIn": "test",
* 	"facebook": "test",
* 	"description": "test",
* 	"createdAt": "20-01-01"
* 	}
* }
*
*
*
* @apiParamExample {json} response-example
* {
*     "rs": "Done",
*     "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQHRlc3QuY29tIiwicm9sZSI6IkRFRkFVTFQiLCJvcmdhbml6YXRpb24iOiIifQ.3Qj3eu8iwXL2v1Rb7qGEf5USQ-WVjRvYiLALWIbWZ5c",
*     "err": null
* }
**/
func (b *basicAuthService) Signup(ctx context.Context, user model.User) (rs string, token string, err error) {
	if user.Email == "" {
		return "User email not found", "", nil
	}
	c := make(chan model.UserReturn)
	go user.Get(c)
	msg := <-c
	close(c)
	if err := msg.Err; err != nil {
		return msg.Message, "", err
	} else if msg.User.Email == "" {
		return msg.Message, "", nil
	}

	// generate JWT
	cc := make(chan model.TokenReturn)
	go model.TokenGen(user.Email, "DEFAULT", "", cc)
	msg2 := <-cc
	close(cc)
	if err := msg2.Err; err != nil {
		return msg2.Message, "", err
	}
	return msg2.Message, msg2.Token, nil
}

func (b *basicAuthService) CreateOrg(ctx context.Context, data model.Organization) (rs string, err error) {
	// TODO implement the business logic of CreateOrg
	return rs, err
}

func (b *basicAuthService) LoginOrg(ctx context.Context, data model.Organization) (rs string, err error) {
	// TODO implement the business logic of CreateOrg
	return rs, err
}
func (b *basicAuthService) Invite(ctx context.Context, email string, org string) (rs string, err error) {
	// TODO implement the business logic of Invite
	return rs, err
}
func (b *basicAuthService) ShowInvites(ctx context.Context) (org string, admin string, err error) {
	// TODO implement the business logic of ShowInvites
	return org, admin, err
}
func (b *basicAuthService) ShowProfile(ctx context.Context) (orgs []model.Organization, user []model.User, events []model.Event, err error) {
	// TODO implement the business logic of ShowProfile
	return orgs, user, events, err
}

// NewBasicAuthService returns a naive, stateless implementation of AuthService.
func NewBasicAuthService() AuthService {
	return &basicAuthService{}
}

// New returns a AuthService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthService {
	var svc AuthService = NewBasicAuthService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
