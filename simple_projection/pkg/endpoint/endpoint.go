package endpoint

import (
	"context"

	model "github.com/GDGVIT/Project-Hades/model"
	service "github.com/GDGVIT/Project-Hades/simple_projection/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// ProjectAllRequest collects the request parameters for the ProjectAll method.
type ProjectAllRequest struct {
	Event string      `json:"event"`
	Query model.Query `json:"query"`
}

// ProjectAllResponse collects the response parameters for the ProjectAll method.
type ProjectAllResponse struct {
	Rs  []model.Participant `json:"rs"`
	Err error               `json:"err"`
}

// MakeProjectAllEndpoint returns an endpoint that invokes ProjectAll on the service.
func MakeProjectAllEndpoint(s service.SimpleProjectionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ProjectAllRequest)
		rs, err := s.ProjectAll(ctx, req.Event, req.Query)
		return ProjectAllResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ProjectAllResponse) Failed() error {
	return r.Err
}

// ProjectPresentRequest collects the request parameters for the ProjectPresent method.
type ProjectPresentRequest struct {
	Event string      `json:"event"`
	Day   int         `json:"day"`
	Query model.Query `json:"query"`
}

// ProjectPresentResponse collects the response parameters for the ProjectPresent method.
type ProjectPresentResponse struct {
	Rs  []model.Participant `json:"rs"`
	Err error               `json:"err"`
}

// MakeProjectPresentEndpoint returns an endpoint that invokes ProjectPresent on the service.
func MakeProjectPresentEndpoint(s service.SimpleProjectionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ProjectPresentRequest)
		rs, err := s.ProjectPresent(ctx, req.Event, req.Day, req.Query)
		return ProjectPresentResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ProjectPresentResponse) Failed() error {
	return r.Err
}

// ProjectAbsentRequest collects the request parameters for the ProjectAbsent method.
type ProjectAbsentRequest struct {
	Event string      `json:"event"`
	Day   int         `json:"day"`
	Query model.Query `json:"query"`
}

// ProjectAbsentResponse collects the response parameters for the ProjectAbsent method.
type ProjectAbsentResponse struct {
	Rs  []model.Participant `json:"rs"`
	Err error               `json:"err"`
}

// MakeProjectAbsentEndpoint returns an endpoint that invokes ProjectAbsent on the service.
func MakeProjectAbsentEndpoint(s service.SimpleProjectionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ProjectAbsentRequest)
		rs, err := s.ProjectAbsent(ctx, req.Event, req.Day, req.Query)
		return ProjectAbsentResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ProjectAbsentResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// ProjectAll implements Service. Primarily useful in a client.
func (e Endpoints) ProjectAll(ctx context.Context, event string, query model.Query) (rs []model.Participant, err error) {
	request := ProjectAllRequest{
		Event: event,
		Query: query,
	}
	response, err := e.ProjectAllEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ProjectAllResponse).Rs, response.(ProjectAllResponse).Err
}

// ProjectPresent implements Service. Primarily useful in a client.
func (e Endpoints) ProjectPresent(ctx context.Context, event string, day int, query model.Query) (rs []model.Participant, err error) {
	request := ProjectPresentRequest{
		Day:   day,
		Event: event,
		Query: query,
	}
	response, err := e.ProjectPresentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ProjectPresentResponse).Rs, response.(ProjectPresentResponse).Err
}

// ProjectAbsent implements Service. Primarily useful in a client.
func (e Endpoints) ProjectAbsent(ctx context.Context, event string, day int, query model.Query) (rs []model.Participant, err error) {
	request := ProjectAbsentRequest{
		Day:   day,
		Event: event,
		Query: query,
	}
	response, err := e.ProjectAbsentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ProjectAbsentResponse).Rs, response.(ProjectAbsentResponse).Err
}
