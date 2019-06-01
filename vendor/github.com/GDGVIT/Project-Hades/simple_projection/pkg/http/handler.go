package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpoint "github.com/GDGVIT/Project-Hades/simple_projection/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
)

// makeProjectAllHandler creates the handler logic
func makeProjectAllHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/simple-projection/project-all", http1.NewServer(endpoints.ProjectAllEndpoint, decodeProjectAllRequest, encodeProjectAllResponse, options...))
}

// decodeProjectAllResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeProjectAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ProjectAllRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeProjectAllResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeProjectAllResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeProjectPresentHandler creates the handler logic
func makeProjectPresentHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/simple-projection/project-present", http1.NewServer(endpoints.ProjectPresentEndpoint, decodeProjectPresentRequest, encodeProjectPresentResponse, options...))
}

// decodeProjectPresentResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeProjectPresentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ProjectPresentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeProjectPresentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeProjectPresentResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeProjectAbsentHandler creates the handler logic
func makeProjectAbsentHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/simple-projection/project-absent", http1.NewServer(endpoints.ProjectAbsentEndpoint, decodeProjectAbsentRequest, encodeProjectAbsentResponse, options...))
}

// decodeProjectAbsentResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeProjectAbsentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ProjectAbsentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeProjectAbsentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeProjectAbsentResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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
