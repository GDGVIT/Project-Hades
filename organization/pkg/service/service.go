package service

import (
	"context"

	"github.com/GDGVIT/Project-Hades/model"
)

// OrganizationService describes the service.
type OrganizationService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Login(ctx context.Context, email string, password string) (rs string, token string, err error)
	Signup(ctx context.Context, user model.User) (rs string, token string, err error)
	CreateOrg(ctx context.Context, data model.Organization) (rs string, err error)
	LoginOrg(ctx context.Context, name string, token string) (rs string, err error)
	AddMembers(ctx context.Context, email string, org string) (rs string, err error)
	BulkAddMembers(ctx context.Context, emails []string, org string) (rs string, err error)
	RemoveMembers(ctx context.Context, email string, org string) (rs string, err error)
	BulkRemoveMembers(ctx context.Context, email string, org string) (rs string, err error)
	ShowProfile(ctx context.Context) (orgs []model.Organization, user []model.User, events []model.Event, err error)
}

type basicOrganizationService struct{}

/**
* @api {post} /api/v1/org/login login as a user
* @apiName login as a user
* @apiGroup organization
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
func (b *basicOrganizationService) Login(ctx context.Context, email string, password string) (rs string, token string, err error) {
	token, err = model.Login(email, password, "DEFAULT", "")
	if err != nil {
		return "Some error occurred", token, err
	}
	return "Done!!", token, nil
}

/**
* @api {post} /api/v1/org/signup signup as a user
* @apiName signup as a user
* @apiGroup organization
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
func (b *basicOrganizationService) Signup(ctx context.Context, user model.User) (rs string, token string, err error) {
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

/**
* @api {post} /api/v1/org/create-org create an organization
* @apiName create an organization
* @apiGroup organization
*
* @apiPermission user
* @apiParam {string} name name of the org
* @apiParam {string} location location of the org
* @apiParam {string} description description of the org
* @apiParam {string} tag tag of the org
* @apiParam {string} website website of the org
*
* @apiParamExample {json} request-example
*
*{
*	"data":{
*	"name":"DSC-VIT",
*	"location":"India",
*	"description":"Developer Student Clubs",
*	"tag":"technical",
*	"website":"https://dsv-vit-vellore.com"
*}
*}
* }
*
**/
func (b *basicOrganizationService) CreateOrg(ctx context.Context, data model.Organization) (rs string, err error) {
	tk, err := model.VerifyToken(ctx)
	if err != nil {
		return "", err
	}

	if data.Name == "" {
		return "ORG name needed", nil
	}
	if err = model.CreateNewOrg(data); err != nil {
		return "", err
	}
	if err := model.AddPolicy(tk.Email, data.Name, "admin"); err != nil {
		return "Error creating policy", err
	}
	if er := model.AddPolicy(tk.Email, data.Name, "member"); er != nil {
		return "Error creating policy", err
	}
	return "Created organization", nil
}

/**
* @api {post} /api/v1/org/login-org login to the org workspace (for privellage escalation)
* @apiName login to the org workspace (for privellage escalation)
* @apiGroup organization
* @apiPermission organization member
*
* @apiParam {string} password password of the user
* @apiParam {string} email email of the user
*
*
* @apiParamExample {json} request-example
*{
*	"name":"GDGVIT",
*	"role":"admin"
*}
*
* @apiParamExample {json} response-example
*{
*    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQHRlc3QuY29tIiwicm9sZSI6IkRFRkFVTFQiLCJvcmdhbml6YXRpb24iOiIifQ.3Qj3eu8iwXL2v1Rb7qGEf5USQ-WVjRvYiLALWIbWZ5c",
*    "err": null
*}
**/
func (b *basicOrganizationService) LoginOrg(ctx context.Context, name string, role string) (rs string, err error) {
	if role != "member" && role != "admin" {
		return "Invalid role. Only <member|admin> allowed", nil
	}
	token, err := model.VerifyToken(ctx)
	if err != nil {
		return "unable to verufy user", err
	}

	// check if user is authorized for the role
	if !model.Enforce(token.Email, name, role) {
		return "failed to authenticate user", nil
	}
	// generate and return token
	cc := make(chan model.TokenReturn)
	go model.TokenGen(token.Email, role, name, cc)
	tk := <-cc
	return tk.Token, tk.Err
}

/**
* @api {post} /api/v1/org/add-members invite a user org
* @apiName invite a user to an org
* @apiGroup organization
* @apiPermission organization admin
*
* @apiParam {string} email email of the user
* @apiParam {string} org name of the organization
*
* @apiParamExample {json} request-example
*
*{
*	"data":{
*	"name":"DSC-VIT",
*	"location":"India",
*	"description":"Developer Student Clubs",
*	"tag":"technical",
*	"website":"https://dsv-vit-vellore.com"
*}
*}
* }
 */
func (b *basicOrganizationService) AddMembers(ctx context.Context, email string, org string) (rs string, err error) {
	tk, err := model.VerifyToken(ctx)
	if err != nil || !model.Enforce(tk.Email, tk.Organization, "admin") {
		return "error authorizing user", err
	}
	if err := model.InviteUserToOrg(email, org); err != nil {
		return "", err
	}
	if err := model.AddPolicy(email, org, "member"); err != nil {
		return "error adding policy", err
	}

	return "successful", nil
}
func (b *basicOrganizationService) BulkAddMembers(ctx context.Context, emails []string, org string) (rs string, err error) {
	// TODO implement the business logic of BulkAddMembers
	return rs, err
}
func (b *basicOrganizationService) RemoveMembers(ctx context.Context, email string, org string) (rs string, err error) {
	// TODO implement the business logic of RemoveMembers
	return rs, err
}
func (b *basicOrganizationService) BulkRemoveMembers(ctx context.Context, email string, org string) (rs string, err error) {
	// TODO implement the business logic of BulkRemoveMembers
	return rs, err
}
func (b *basicOrganizationService) ShowProfile(ctx context.Context) (orgs []model.Organization, user []model.User, events []model.Event, err error) {
	// TODO implement the business logic of ShowProfile
	return orgs, user, events, err
}

// NewBasicOrganizationService returns a naive, stateless implementation of OrganizationService.
func NewBasicOrganizationService() OrganizationService {
	return &basicOrganizationService{}
}

// New returns a OrganizationService with all of the expected middleware wired in.
func New(middleware []Middleware) OrganizationService {
	var svc OrganizationService = NewBasicOrganizationService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
