package endpoint

import (
	"context"

	service "github.com/GDGVIT/Project-Hades/auth/pkg/service"
	model "github.com/GDGVIT/Project-Hades/model"
	endpoint "github.com/go-kit/kit/endpoint"
)

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	Rs    string `json:"rs"`
	Token string `json:"token"`
	Err   error  `json:"err"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		rs, token, err := s.Login(ctx, req.Email, req.Password)
		return LoginResponse{
			Err:   err,
			Rs:    rs,
			Token: token,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.Err
}

// SignupRequest collects the request parameters for the Signup method.
type SignupRequest struct {
	User model.User `json:"user"`
}

// SignupResponse collects the response parameters for the Signup method.
type SignupResponse struct {
	Rs    string `json:"rs"`
	Token string `json:"token"`
	Err   error  `json:"err"`
}

// MakeSignupEndpoint returns an endpoint that invokes Signup on the service.
func MakeSignupEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignupRequest)
		rs, token, err := s.Signup(ctx, req.User)
		return SignupResponse{
			Err:   err,
			Rs:    rs,
			Token: token,
		}, nil
	}
}

// Failed implements Failer.
func (r SignupResponse) Failed() error {
	return r.Err
}

// CreateOrgRequest collects the request parameters for the CreateOrg method.
type CreateOrgRequest struct {
	Data model.Organization `json:"data"`
}

// CreateOrgResponse collects the response parameters for the CreateOrg method.
type CreateOrgResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateOrgEndpoint returns an endpoint that invokes CreateOrg on the service.
func MakeCreateOrgEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateOrgRequest)
		rs, err := s.CreateOrg(ctx, req.Data)
		return CreateOrgResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateOrgResponse) Failed() error {
	return r.Err
}

// LoginOrgRequest collects the request parameters for the LoginOrg method.
type LoginOrgRequest struct {
	Data model.Organization `json:"data"`
}

// LoginOrgResponse collects the response parameters for the LoginOrg method.
type LoginOrgResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeLoginOrgEndpoint returns an endpoint that invokes LoginOrg on the service.
func MakeLoginOrgEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginOrgRequest)
		rs, err := s.LoginOrg(ctx, req.Data)
		return LoginOrgResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginOrgResponse) Failed() error {
	return r.Err
}

// InviteRequest collects the request parameters for the Invite method.
type InviteRequest struct {
	Email string `json:"email"`
	Org   string `json:"org"`
}

// InviteResponse collects the response parameters for the Invite method.
type InviteResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeInviteEndpoint returns an endpoint that invokes Invite on the service.
func MakeInviteEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InviteRequest)
		rs, err := s.Invite(ctx, req.Email, req.Org)
		return InviteResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r InviteResponse) Failed() error {
	return r.Err
}

// ShowInvitesRequest collects the request parameters for the ShowInvites method.
type ShowInvitesRequest struct{}

// ShowInvitesResponse collects the response parameters for the ShowInvites method.
type ShowInvitesResponse struct {
	Org   string `json:"org"`
	Admin string `json:"admin"`
	Err   error  `json:"err"`
}

// MakeShowInvitesEndpoint returns an endpoint that invokes ShowInvites on the service.
func MakeShowInvitesEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		org, admin, err := s.ShowInvites(ctx)
		return ShowInvitesResponse{
			Admin: admin,
			Err:   err,
			Org:   org,
		}, nil
	}
}

// Failed implements Failer.
func (r ShowInvitesResponse) Failed() error {
	return r.Err
}

// ShowProfileRequest collects the request parameters for the ShowProfile method.
type ShowProfileRequest struct{}

// ShowProfileResponse collects the response parameters for the ShowProfile method.
type ShowProfileResponse struct {
	Orgs   []model.Organization `json:"orgs"`
	User   []model.User         `json:"user"`
	Events []model.Event        `json:"events"`
	Err    error                `json:"err"`
}

// MakeShowProfileEndpoint returns an endpoint that invokes ShowProfile on the service.
func MakeShowProfileEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		orgs, user, events, err := s.ShowProfile(ctx)
		return ShowProfileResponse{
			Err:    err,
			Events: events,
			Orgs:   orgs,
			User:   user,
		}, nil
	}
}

// Failed implements Failer.
func (r ShowProfileResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Login implements Service. Primarily useful in a client.
func (e Endpoints) Login(ctx context.Context, email string, password string) (rs string, token string, err error) {
	request := LoginRequest{
		Email:    email,
		Password: password,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).Rs, response.(LoginResponse).Token, response.(LoginResponse).Err
}

// Signup implements Service. Primarily useful in a client.
func (e Endpoints) Signup(ctx context.Context, user model.User) (rs string, token string, err error) {
	request := SignupRequest{User: user}
	response, err := e.SignupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SignupResponse).Rs, response.(SignupResponse).Token, response.(SignupResponse).Err
}

// CreateOrg implements Service. Primarily useful in a client.
func (e Endpoints) CreateOrg(ctx context.Context, data model.Organization) (rs string, err error) {
	request := CreateOrgRequest{Data: data}
	response, err := e.CreateOrgEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateOrgResponse).Rs, response.(CreateOrgResponse).Err
}

// Invite implements Service. Primarily useful in a client.
func (e Endpoints) Invite(ctx context.Context, email string, org string) (rs string, err error) {
	request := InviteRequest{
		Email: email,
		Org:   org,
	}
	response, err := e.InviteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(InviteResponse).Rs, response.(InviteResponse).Err
}

// ShowInvites implements Service. Primarily useful in a client.
func (e Endpoints) ShowInvites(ctx context.Context) (org string, admin string, err error) {
	request := ShowInvitesRequest{}
	response, err := e.ShowInvitesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ShowInvitesResponse).Org, response.(ShowInvitesResponse).Admin, response.(ShowInvitesResponse).Err
}

// ShowProfile implements Service. Primarily useful in a client.
func (e Endpoints) ShowProfile(ctx context.Context) (orgs []model.Organization, user []model.User, events []model.Event, err error) {
	request := ShowProfileRequest{}
	response, err := e.ShowProfileEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ShowProfileResponse).Orgs, response.(ShowProfileResponse).User, response.(ShowProfileResponse).Events, response.(ShowProfileResponse).Err
}

// LoginOrg implements Service. Primarily useful in a client.
func (e Endpoints) LoginOrg(ctx context.Context, data model.Organization) (rs string, err error) {
	request := LoginOrgRequest{Data: data}
	response, err := e.LoginOrgEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginOrgResponse).Rs, response.(LoginOrgResponse).Err
}
