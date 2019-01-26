package endpoint

import (
	"context"

	model "github.com/GDGVIT/Project-Hades/model"
	service "github.com/GDGVIT/Project-Hades/participants/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateAttendeeRequest collects the request parameters for the CreateAttendee method.
type CreateAttendeeRequest struct {
	Details model.Attendee `json:"details"`
}

// CreateAttendeeResponse collects the response parameters for the CreateAttendee method.
type CreateAttendeeResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateAttendeeEndpoint returns an endpoint that invokes CreateAttendee on the service.
func MakeCreateAttendeeEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAttendeeRequest)
		rs, err := s.CreateAttendee(ctx, req.Details)
		return CreateAttendeeResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateAttendeeResponse) Failed() error {
	return r.Err
}

// ReadAttendeeRequest collects the request parameters for the ReadAttendee method.
type ReadAttendeeRequest struct {
	Query model.Query `json:"query"`
}

// ReadAttendeeResponse collects the response parameters for the ReadAttendee method.
type ReadAttendeeResponse struct {
	Rs  []model.Attendee `json:"rs"`
	Err error            `json:"err"`
}

// MakeReadAttendeeEndpoint returns an endpoint that invokes ReadAttendee on the service.
func MakeReadAttendeeEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadAttendeeRequest)
		rs, err := s.ReadAttendee(ctx, req.Query)
		return ReadAttendeeResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ReadAttendeeResponse) Failed() error {
	return r.Err
}

// UpdateAttendeeRequest collects the request parameters for the UpdateAttendee method.
type UpdateAttendeeRequest struct {
	Query model.Query `json:"query"`
}

// UpdateAttendeeResponse collects the response parameters for the UpdateAttendee method.
type UpdateAttendeeResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeUpdateAttendeeEndpoint returns an endpoint that invokes UpdateAttendee on the service.
func MakeUpdateAttendeeEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAttendeeRequest)
		rs, err := s.UpdateAttendee(ctx, req.Query)
		return UpdateAttendeeResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateAttendeeResponse) Failed() error {
	return r.Err
}

// DeleteAttendeeRequest collects the request parameters for the DeleteAttendee method.
type DeleteAttendeeRequest struct {
	Query model.Query `json:"query"`
}

// DeleteAttendeeResponse collects the response parameters for the DeleteAttendee method.
type DeleteAttendeeResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteAttendeeEndpoint returns an endpoint that invokes DeleteAttendee on the service.
func MakeDeleteAttendeeEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAttendeeRequest)
		rs, err := s.DeleteAttendee(ctx, req.Query)
		return DeleteAttendeeResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteAttendeeResponse) Failed() error {
	return r.Err
}

// DeleteAllAttendeeRequest collects the request parameters for the DeleteAllAttendee method.
type DeleteAllAttendeeRequest struct {
	Query model.Query `json:"query"`
}

// DeleteAllAttendeeResponse collects the response parameters for the DeleteAllAttendee method.
type DeleteAllAttendeeResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteAllAttendeeEndpoint returns an endpoint that invokes DeleteAllAttendee on the service.
func MakeDeleteAllAttendeeEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAllAttendeeRequest)
		rs, err := s.DeleteAllAttendee(ctx, req.Query)
		return DeleteAllAttendeeResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteAllAttendeeResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateAttendee implements Service. Primarily useful in a client.
func (e Endpoints) CreateAttendee(ctx context.Context, details model.Attendee) (rs string, err error) {
	request := CreateAttendeeRequest{
		Details: details,
	}
	response, err := e.CreateAttendeeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateAttendeeResponse).Rs, response.(CreateAttendeeResponse).Err
}

// ReadAttendee implements Service. Primarily useful in a client.
func (e Endpoints) ReadAttendee(ctx context.Context, query model.Query) (rs []model.Attendee, err error) {
	request := ReadAttendeeRequest{Query: query}
	response, err := e.ReadAttendeeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadAttendeeResponse).Rs, response.(ReadAttendeeResponse).Err
}

// UpdateAttendee implements Service. Primarily useful in a client.
func (e Endpoints) UpdateAttendee(ctx context.Context, query model.Query) (rs string, err error) {
	request := UpdateAttendeeRequest{Query: query}
	response, err := e.UpdateAttendeeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateAttendeeResponse).Rs, response.(UpdateAttendeeResponse).Err
}

// DeleteAttendee implements Service. Primarily useful in a client.
func (e Endpoints) DeleteAttendee(ctx context.Context, query model.Query) (rs string, err error) {
	request := DeleteAttendeeRequest{Query: query}
	response, err := e.DeleteAttendeeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteAttendeeResponse).Rs, response.(DeleteAttendeeResponse).Err
}

// DeleteAllAttendee implements Service. Primarily useful in a client.
func (e Endpoints) DeleteAllAttendee(ctx context.Context, query model.Query) (rs string, err error) {
	request := DeleteAllAttendeeRequest{Query: query}
	response, err := e.DeleteAllAttendeeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteAllAttendeeResponse).Rs, response.(DeleteAllAttendeeResponse).Err
}
