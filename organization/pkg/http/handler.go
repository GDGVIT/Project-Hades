package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpoint "github.com/GDGVIT/Project-Hades/organization/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
)

// makeLoginHandler creates the handler logic
func makeLoginHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/org/login", http1.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...))
}

// decodeLoginResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeLoginResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSignupHandler creates the handler logic
func makeSignupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/org/signup", http1.NewServer(endpoints.SignupEndpoint, decodeSignupRequest, encodeSignupResponse, options...))
}

// decodeSignupResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSignupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SignupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSignupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSignupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateOrgHandler creates the handler logic
func makeCreateOrgHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/org/create-org", http1.NewServer(endpoints.CreateOrgEndpoint, decodeCreateOrgRequest, encodeCreateOrgResponse, options...))
}

// decodeCreateOrgResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateOrgRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateOrgRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateOrgResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateOrgResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeLoginOrgHandler creates the handler logic
func makeLoginOrgHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/org/login-org", http1.NewServer(endpoints.LoginOrgEndpoint, decodeLoginOrgRequest, encodeLoginOrgResponse, options...))
}

// decodeLoginOrgResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeLoginOrgRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.LoginOrgRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeLoginOrgResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeLoginOrgResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddMembersHandler creates the handler logic
func makeAddMembersHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/org/add-members", http1.NewServer(endpoints.AddMembersEndpoint, decodeAddMembersRequest, encodeAddMembersResponse, options...))
}

// decodeAddMembersResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddMembersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.AddMembersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddMembersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddMembersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeBulkAddMembersHandler creates the handler logic
func makeBulkAddMembersHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/org/bulk-add-members", http1.NewServer(endpoints.BulkAddMembersEndpoint, decodeBulkAddMembersRequest, encodeBulkAddMembersResponse, options...))
}

// decodeBulkAddMembersResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeBulkAddMembersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.BulkAddMembersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeBulkAddMembersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeBulkAddMembersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeRemoveMembersHandler creates the handler logic
func makeRemoveMembersHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/org/remove-members", http1.NewServer(endpoints.RemoveMembersEndpoint, decodeRemoveMembersRequest, encodeRemoveMembersResponse, options...))
}

// decodeRemoveMembersResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRemoveMembersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.RemoveMembersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRemoveMembersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRemoveMembersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeBulkRemoveMembersHandler creates the handler logic
func makeBulkRemoveMembersHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/bapi/v1/org/ulk-remove-members", http1.NewServer(endpoints.BulkRemoveMembersEndpoint, decodeBulkRemoveMembersRequest, encodeBulkRemoveMembersResponse, options...))
}

// decodeBulkRemoveMembersResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeBulkRemoveMembersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.BulkRemoveMembersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeBulkRemoveMembersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeBulkRemoveMembersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeShowProfileHandler creates the handler logic
func makeShowProfileHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/org/show-profile", http1.NewServer(endpoints.ShowProfileEndpoint, decodeShowProfileRequest, encodeShowProfileResponse, options...))
}

// decodeShowProfileResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeShowProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ShowProfileRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeShowProfileResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeShowProfileResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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
