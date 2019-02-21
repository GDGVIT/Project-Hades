package endpoint

import (
	"context"
	service "github.com/GDGVIT/Project-Hades/guests/pkg/service"
	model "github.com/GDGVIT/Project-Hades/model"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateGuestRequest collects the request parameters for the CreateGuest method.
type CreateGuestRequest struct {
	Event string      `json:"event"`
	Guest model.Guest `json:"guest"`
}

// CreateGuestResponse collects the response parameters for the CreateGuest method.
type CreateGuestResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateGuestEndpoint returns an endpoint that invokes CreateGuest on the service.
func MakeCreateGuestEndpoint(s service.GuestsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateGuestRequest)
		rs, err := s.CreateGuest(ctx, req.Event, req.Guest)
		return CreateGuestResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateGuestResponse) Failed() error {
	return r.Err
}

// ReadGuestRequest collects the request parameters for the ReadGuest method.
type ReadGuestRequest struct {
	Query model.Query `json:"query"`
}

// ReadGuestResponse collects the response parameters for the ReadGuest method.
type ReadGuestResponse struct {
	Rs  []model.Guest `json:"rs"`
	Err error         `json:"err"`
}

// MakeReadGuestEndpoint returns an endpoint that invokes ReadGuest on the service.
func MakeReadGuestEndpoint(s service.GuestsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadGuestRequest)
		rs, err := s.ReadGuest(ctx, req.Query)
		return ReadGuestResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ReadGuestResponse) Failed() error {
	return r.Err
}

// UpdateGuestRequest collects the request parameters for the UpdateGuest method.
type UpdateGuestRequest struct {
	Query model.Query `json:"query"`
}

// UpdateGuestResponse collects the response parameters for the UpdateGuest method.
type UpdateGuestResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeUpdateGuestEndpoint returns an endpoint that invokes UpdateGuest on the service.
func MakeUpdateGuestEndpoint(s service.GuestsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateGuestRequest)
		rs, err := s.UpdateGuest(ctx, req.Query)
		return UpdateGuestResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateGuestResponse) Failed() error {
	return r.Err
}

// DeleteGuestRequest collects the request parameters for the DeleteGuest method.
type DeleteGuestRequest struct {
	Query model.Query `json:"query"`
}

// DeleteGuestResponse collects the response parameters for the DeleteGuest method.
type DeleteGuestResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteGuestEndpoint returns an endpoint that invokes DeleteGuest on the service.
func MakeDeleteGuestEndpoint(s service.GuestsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteGuestRequest)
		rs, err := s.DeleteGuest(ctx, req.Query)
		return DeleteGuestResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteGuestResponse) Failed() error {
	return r.Err
}

// CreateSponsorRequest collects the request parameters for the CreateSponsor method.
type CreateSponsorRequest struct {
	Event   string            `json:"event"`
	Sponsor model.Participant `json:"sponsor"`
}

// CreateSponsorResponse collects the response parameters for the CreateSponsor method.
type CreateSponsorResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateSponsorEndpoint returns an endpoint that invokes CreateSponsor on the service.
func MakeCreateSponsorEndpoint(s service.GuestsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateSponsorRequest)
		rs, err := s.CreateSponsor(ctx, req.Event, req.Sponsor)
		return CreateSponsorResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateSponsorResponse) Failed() error {
	return r.Err
}

// ReadSponsorRequest collects the request parameters for the ReadSponsor method.
type ReadSponsorRequest struct {
	Query model.Query `json:"query"`
}

// ReadSponsorResponse collects the response parameters for the ReadSponsor method.
type ReadSponsorResponse struct {
	Rs  []model.Participant `json:"rs"`
	Err error               `json:"err"`
}

// MakeReadSponsorEndpoint returns an endpoint that invokes ReadSponsor on the service.
func MakeReadSponsorEndpoint(s service.GuestsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadSponsorRequest)
		rs, err := s.ReadSponsor(ctx, req.Query)
		return ReadSponsorResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ReadSponsorResponse) Failed() error {
	return r.Err
}

// UpdateSponsorRequest collects the request parameters for the UpdateSponsor method.
type UpdateSponsorRequest struct {
	Query model.Query `json:"query"`
}

// UpdateSponsorResponse collects the response parameters for the UpdateSponsor method.
type UpdateSponsorResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeUpdateSponsorEndpoint returns an endpoint that invokes UpdateSponsor on the service.
func MakeUpdateSponsorEndpoint(s service.GuestsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateSponsorRequest)
		rs, err := s.UpdateSponsor(ctx, req.Query)
		return UpdateSponsorResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateSponsorResponse) Failed() error {
	return r.Err
}

// DeleteSponsorRequest collects the request parameters for the DeleteSponsor method.
type DeleteSponsorRequest struct {
	Query model.Query `json:"query"`
}

// DeleteSponsorResponse collects the response parameters for the DeleteSponsor method.
type DeleteSponsorResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteSponsorEndpoint returns an endpoint that invokes DeleteSponsor on the service.
func MakeDeleteSponsorEndpoint(s service.GuestsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteSponsorRequest)
		rs, err := s.DeleteSponsor(ctx, req.Query)
		return DeleteSponsorResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteSponsorResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateGuest implements Service. Primarily useful in a client.
func (e Endpoints) CreateGuest(ctx context.Context, event string, guest model.Guest) (rs string, err error) {
	request := CreateGuestRequest{
		Event: event,
		Guest: guest,
	}
	response, err := e.CreateGuestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateGuestResponse).Rs, response.(CreateGuestResponse).Err
}

// ReadGuest implements Service. Primarily useful in a client.
func (e Endpoints) ReadGuest(ctx context.Context, query model.Query) (rs []model.Guest, err error) {
	request := ReadGuestRequest{Query: query}
	response, err := e.ReadGuestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadGuestResponse).Rs, response.(ReadGuestResponse).Err
}

// UpdateGuest implements Service. Primarily useful in a client.
func (e Endpoints) UpdateGuest(ctx context.Context, query model.Query) (rs string, err error) {
	request := UpdateGuestRequest{Query: query}
	response, err := e.UpdateGuestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateGuestResponse).Rs, response.(UpdateGuestResponse).Err
}

// DeleteGuest implements Service. Primarily useful in a client.
func (e Endpoints) DeleteGuest(ctx context.Context, query model.Query) (rs string, err error) {
	request := DeleteGuestRequest{Query: query}
	response, err := e.DeleteGuestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteGuestResponse).Rs, response.(DeleteGuestResponse).Err
}

// CreateSponsor implements Service. Primarily useful in a client.
func (e Endpoints) CreateSponsor(ctx context.Context, event string, sponsor model.Participant) (rs string, err error) {
	request := CreateSponsorRequest{
		Event:   event,
		Sponsor: sponsor,
	}
	response, err := e.CreateSponsorEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateSponsorResponse).Rs, response.(CreateSponsorResponse).Err
}

// ReadSponsor implements Service. Primarily useful in a client.
func (e Endpoints) ReadSponsor(ctx context.Context, query model.Query) (rs []model.Participant, err error) {
	request := ReadSponsorRequest{Query: query}
	response, err := e.ReadSponsorEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadSponsorResponse).Rs, response.(ReadSponsorResponse).Err
}

// UpdateSponsor implements Service. Primarily useful in a client.
func (e Endpoints) UpdateSponsor(ctx context.Context, query model.Query) (rs string, err error) {
	request := UpdateSponsorRequest{Query: query}
	response, err := e.UpdateSponsorEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateSponsorResponse).Rs, response.(UpdateSponsorResponse).Err
}

// DeleteSponsor implements Service. Primarily useful in a client.
func (e Endpoints) DeleteSponsor(ctx context.Context, query model.Query) (rs string, err error) {
	request := DeleteSponsorRequest{Query: query}
	response, err := e.DeleteSponsorEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteSponsorResponse).Rs, response.(DeleteSponsorResponse).Err
}
