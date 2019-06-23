package endpoint

import (
	"context"

	model "github.com/GDGVIT/Project-Hades/model"
	service "github.com/GDGVIT/Project-Hades/organization/pkg/service"
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
func MakeLoginEndpoint(s service.OrganizationService) endpoint.Endpoint {
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
func MakeSignupEndpoint(s service.OrganizationService) endpoint.Endpoint {
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
func MakeCreateOrgEndpoint(s service.OrganizationService) endpoint.Endpoint {
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
	Name string `json:"name"`
	Role string `json:"role"`
}

// LoginOrgResponse collects the response parameters for the LoginOrg method.
type LoginOrgResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeLoginOrgEndpoint returns an endpoint that invokes LoginOrg on the service.
func MakeLoginOrgEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginOrgRequest)
		rs, err := s.LoginOrg(ctx, req.Name, req.Role)
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

// AddMembersRequest collects the request parameters for the AddMembers method.
type AddMembersRequest struct {
	Email string `json:"email"`
	Org   string `json:"org"`
}

// AddMembersResponse collects the response parameters for the AddMembers method.
type AddMembersResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeAddMembersEndpoint returns an endpoint that invokes AddMembers on the service.
func MakeAddMembersEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddMembersRequest)
		rs, err := s.AddMembers(ctx, req.Email, req.Org)
		return AddMembersResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r AddMembersResponse) Failed() error {
	return r.Err
}

// BulkAddMembersRequest collects the request parameters for the BulkAddMembers method.
type BulkAddMembersRequest struct {
	Emails []string `json:"emails"`
	Org    string   `json:"org"`
}

// BulkAddMembersResponse collects the response parameters for the BulkAddMembers method.
type BulkAddMembersResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeBulkAddMembersEndpoint returns an endpoint that invokes BulkAddMembers on the service.
func MakeBulkAddMembersEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BulkAddMembersRequest)
		rs, err := s.BulkAddMembers(ctx, req.Emails, req.Org)
		return BulkAddMembersResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r BulkAddMembersResponse) Failed() error {
	return r.Err
}

// RemoveMembersRequest collects the request parameters for the RemoveMembers method.
type RemoveMembersRequest struct {
	Email string `json:"email"`
	Org   string `json:"org"`
}

// RemoveMembersResponse collects the response parameters for the RemoveMembers method.
type RemoveMembersResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeRemoveMembersEndpoint returns an endpoint that invokes RemoveMembers on the service.
func MakeRemoveMembersEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveMembersRequest)
		rs, err := s.RemoveMembers(ctx, req.Email, req.Org)
		return RemoveMembersResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r RemoveMembersResponse) Failed() error {
	return r.Err
}

// BulkRemoveMembersRequest collects the request parameters for the BulkRemoveMembers method.
type BulkRemoveMembersRequest struct {
	Email string `json:"email"`
	Org   string `json:"org"`
}

// BulkRemoveMembersResponse collects the response parameters for the BulkRemoveMembers method.
type BulkRemoveMembersResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeBulkRemoveMembersEndpoint returns an endpoint that invokes BulkRemoveMembers on the service.
func MakeBulkRemoveMembersEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BulkRemoveMembersRequest)
		rs, err := s.BulkRemoveMembers(ctx, req.Email, req.Org)
		return BulkRemoveMembersResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r BulkRemoveMembersResponse) Failed() error {
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
func MakeShowProfileEndpoint(s service.OrganizationService) endpoint.Endpoint {
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

// LoginOrg implements Service. Primarily useful in a client.
func (e Endpoints) LoginOrg(ctx context.Context, name string, role string) (rs string, err error) {
	request := LoginOrgRequest{Name: name, Role: role}
	response, err := e.LoginOrgEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginOrgResponse).Rs, response.(LoginOrgResponse).Err
}

// AddMembers implements Service. Primarily useful in a client.
func (e Endpoints) AddMembers(ctx context.Context, email string, org string) (rs string, err error) {
	request := AddMembersRequest{
		Email: email,
		Org:   org,
	}
	response, err := e.AddMembersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddMembersResponse).Rs, response.(AddMembersResponse).Err
}

// BulkAddMembers implements Service. Primarily useful in a client.
func (e Endpoints) BulkAddMembers(ctx context.Context, emails []string, org string) (rs string, err error) {
	request := BulkAddMembersRequest{
		Emails: emails,
		Org:    org,
	}
	response, err := e.BulkAddMembersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(BulkAddMembersResponse).Rs, response.(BulkAddMembersResponse).Err
}

// RemoveMembers implements Service. Primarily useful in a client.
func (e Endpoints) RemoveMembers(ctx context.Context, email string, org string) (rs string, err error) {
	request := RemoveMembersRequest{
		Email: email,
		Org:   org,
	}
	response, err := e.RemoveMembersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RemoveMembersResponse).Rs, response.(RemoveMembersResponse).Err
}

// BulkRemoveMembers implements Service. Primarily useful in a client.
func (e Endpoints) BulkRemoveMembers(ctx context.Context, email string, org string) (rs string, err error) {
	request := BulkRemoveMembersRequest{
		Email: email,
		Org:   org,
	}
	response, err := e.BulkRemoveMembersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(BulkRemoveMembersResponse).Rs, response.(BulkRemoveMembersResponse).Err
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
