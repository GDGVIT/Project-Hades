package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpoint "github.com/GDGVIT/Project-Hades/events/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
)

// makeCreateEventHandler creates the handler logic
func makeCreateEventHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/event/create-event", http1.NewServer(endpoints.CreateEventEndpoint, decodeCreateEventRequest, encodeCreateEventResponse, options...))
}

// decodeCreateEventResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateEventRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateEventResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateEventResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadEventHandler creates the handler logic
func makeReadEventHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/event/read-event", http1.NewServer(endpoints.ReadEventEndpoint, decodeReadEventRequest, encodeReadEventResponse, options...))
}

// decodeReadEventResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ReadEventRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReadEventResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadEventResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateEventHandler creates the handler logic
func makeUpdateEventHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/event/update-event", http1.NewServer(endpoints.UpdateEventEndpoint, decodeUpdateEventRequest, encodeUpdateEventResponse, options...))
}

// decodeUpdateEventResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateEventRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateEventResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateEventResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteEventHandler creates the handler logic
func makeDeleteEventHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/event/delete-event", http1.NewServer(endpoints.DeleteEventEndpoint, decodeDeleteEventRequest, encodeDeleteEventResponse, options...))
}

// decodeDeleteEventResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteEventRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteEventResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteEventResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
