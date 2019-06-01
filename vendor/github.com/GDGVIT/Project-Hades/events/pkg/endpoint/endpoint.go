package endpoint

import (
	"context"

	service "github.com/GDGVIT/Project-Hades/events/pkg/service"
	model "github.com/GDGVIT/Project-Hades/model"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateEventRequest collects the request parameters for the CreateEvent method.
type CreateEventRequest struct {
	Event model.Event `json:"event"`
}

// CreateEventResponse collects the response parameters for the CreateEvent method.
type CreateEventResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateEventEndpoint returns an endpoint that invokes CreateEvent on the service.
func MakeCreateEventEndpoint(s service.EventsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateEventRequest)
		rs, err := s.CreateEvent(ctx, req.Event)
		return CreateEventResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateEventResponse) Failed() error {
	return r.Err
}

// ReadEventRequest collects the request parameters for the ReadEvent method.
type ReadEventRequest struct {
	Query model.Query `json:"query"`
}

// ReadEventResponse collects the response parameters for the ReadEvent method.
type ReadEventResponse struct {
	Rs  []model.Event `json:"rs"`
	Err error         `json:"err"`
}

// MakeReadEventEndpoint returns an endpoint that invokes ReadEvent on the service.
func MakeReadEventEndpoint(s service.EventsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadEventRequest)
		rs, err := s.ReadEvent(ctx, req.Query)
		return ReadEventResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ReadEventResponse) Failed() error {
	return r.Err
}

// UpdateEventRequest collects the request parameters for the UpdateEvent method.
type UpdateEventRequest struct {
	Query model.Query `json:"query"`
}

// UpdateEventResponse collects the response parameters for the UpdateEvent method.
type UpdateEventResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeUpdateEventEndpoint returns an endpoint that invokes UpdateEvent on the service.
func MakeUpdateEventEndpoint(s service.EventsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateEventRequest)
		rs, err := s.UpdateEvent(ctx, req.Query)
		return UpdateEventResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateEventResponse) Failed() error {
	return r.Err
}

// DeleteEventRequest collects the request parameters for the DeleteEvent method.
type DeleteEventRequest struct {
	Query model.Query `json:"query"`
}

// DeleteEventResponse collects the response parameters for the DeleteEvent method.
type DeleteEventResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteEventEndpoint returns an endpoint that invokes DeleteEvent on the service.
func MakeDeleteEventEndpoint(s service.EventsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteEventRequest)
		rs, err := s.DeleteEvent(ctx, req.Query)
		return DeleteEventResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteEventResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateEvent implements Service. Primarily useful in a client.
func (e Endpoints) CreateEvent(ctx context.Context, event model.Event) (rs string, err error) {
	request := CreateEventRequest{Event: event}
	response, err := e.CreateEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateEventResponse).Rs, response.(CreateEventResponse).Err
}

// ReadEvent implements Service. Primarily useful in a client.
func (e Endpoints) ReadEvent(ctx context.Context, query model.Query) (rs []model.Event, err error) {
	request := ReadEventRequest{Query: query}
	response, err := e.ReadEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadEventResponse).Rs, response.(ReadEventResponse).Err
}

// UpdateEvent implements Service. Primarily useful in a client.
func (e Endpoints) UpdateEvent(ctx context.Context, query model.Query) (rs string, err error) {
	request := UpdateEventRequest{Query: query}
	response, err := e.UpdateEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateEventResponse).Rs, response.(UpdateEventResponse).Err
}

// DeleteEvent implements Service. Primarily useful in a client.
func (e Endpoints) DeleteEvent(ctx context.Context, query model.Query) (rs string, err error) {
	request := DeleteEventRequest{Query: query}
	response, err := e.DeleteEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteEventResponse).Rs, response.(DeleteEventResponse).Err
}
