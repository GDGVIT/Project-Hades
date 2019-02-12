package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpoint "github.com/GDGVIT/Project-Hades/participants/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
)

// makeCreateAttendeeHandler creates the handler logic
func makeCreateAttendeeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/participants/create-attendee", http1.NewServer(endpoints.CreateAttendeeEndpoint, decodeCreateAttendeeRequest, encodeCreateAttendeeResponse, options...))
}

// decodeCreateAttendeeResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateAttendeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateAttendeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateAttendeeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateAttendeeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadAttendeeHandler creates the handler logic
func makeReadAttendeeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/participants/read-attendee", http1.NewServer(endpoints.ReadAttendeeEndpoint, decodeReadAttendeeRequest, encodeReadAttendeeResponse, options...))
}

// decodeReadAttendeeResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadAttendeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ReadAttendeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReadAttendeeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadAttendeeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeRmAttendeeHandler creates the handler logic
func makeRmAttendeeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/participants/rm-attendee", http1.NewServer(endpoints.RmAttendeeEndpoint, decodeRmAttendeeRequest, encodeRmAttendeeResponse, options...))
}

// decodeRmAttendeeResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRmAttendeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.RmAttendeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRmAttendeeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRmAttendeeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteAttendeeHandler creates the handler logic
func makeDeleteAttendeeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/participants/delete-attendee", http1.NewServer(endpoints.DeleteAttendeeEndpoint, decodeDeleteAttendeeRequest, encodeDeleteAttendeeResponse, options...))
}

// decodeDeleteAttendeeResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteAttendeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteAttendeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteAttendeeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteAttendeeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteAllAttendeeHandler creates the handler logic
func makeDeleteAllAttendeeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/participants/delete-all-attendee", http1.NewServer(endpoints.DeleteAllAttendeeEndpoint, decodeDeleteAllAttendeeRequest, encodeDeleteAllAttendeeResponse, options...))
}

// decodeDeleteAllAttendeeResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteAllAttendeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteAllAttendeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteAllAttendeeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteAllAttendeeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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
