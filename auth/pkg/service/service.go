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

func (b *basicAuthService) Login(ctx context.Context, email string, password string) (rs string, token string, err error) {
	// TODO implement the business logic of Login
	return rs, token, err
}
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
	return msg.Message, token, nil
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
