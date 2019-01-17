package endpoint

import (
	service "angadsharma1016/omega_microservices/events/pkg/service"
	"context"
	omegadbconfig "github.com/angadsharma1016/omega_dbconfig"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateEventRequest collects the request parameters for the CreateEvent method.
type CreateEventRequest struct {
	S omegadbconfig.Event `json:"s"`
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
		rs, err := s.CreateEvent(ctx, req.S)
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
	S omegadbconfig.Query `json:"s"`
}

// ReadEventResponse collects the response parameters for the ReadEvent method.
type ReadEventResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeReadEventEndpoint returns an endpoint that invokes ReadEvent on the service.
func MakeReadEventEndpoint(s service.EventsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadEventRequest)
		rs, err := s.ReadEvent(ctx, req.S)
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
	S omegadbconfig.Query `json:"s"`
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
		rs, err := s.UpdateEvent(ctx, req.S)
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
	S omegadbconfig.Query `json:"s"`
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
		rs, err := s.DeleteEvent(ctx, req.S)
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
func (e Endpoints) CreateEvent(ctx context.Context, s omegadbconfig.Event) (rs string, err error) {
	request := CreateEventRequest{S: s}
	response, err := e.CreateEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateEventResponse).Rs, response.(CreateEventResponse).Err
}

// ReadEvent implements Service. Primarily useful in a client.
func (e Endpoints) ReadEvent(ctx context.Context, s omegadbconfig.Query) (rs string, err error) {
	request := ReadEventRequest{S: s}
	response, err := e.ReadEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadEventResponse).Rs, response.(ReadEventResponse).Err
}

// UpdateEvent implements Service. Primarily useful in a client.
func (e Endpoints) UpdateEvent(ctx context.Context, s omegadbconfig.Query) (rs string, err error) {
	request := UpdateEventRequest{S: s}
	response, err := e.UpdateEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateEventResponse).Rs, response.(UpdateEventResponse).Err
}

// DeleteEvent implements Service. Primarily useful in a client.
func (e Endpoints) DeleteEvent(ctx context.Context, s omegadbconfig.Query) (rs string, err error) {
	request := DeleteEventRequest{S: s}
	response, err := e.DeleteEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteEventResponse).Rs, response.(DeleteEventResponse).Err
}
