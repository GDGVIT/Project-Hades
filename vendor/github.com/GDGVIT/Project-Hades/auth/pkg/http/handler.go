package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpoint "github.com/GDGVIT/Project-Hades/auth/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
)

// makeLoginHandler creates the handler logic
func makeLoginHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/auth/login", http1.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...))
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
	m.Handle("/api/v1/auth/signup", http1.NewServer(endpoints.SignupEndpoint, decodeSignupRequest, encodeSignupResponse, options...))
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
	m.Handle("/api/v1/auth/create-org", http1.NewServer(endpoints.CreateOrgEndpoint, decodeCreateOrgRequest, encodeCreateOrgResponse, options...))
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

// makeLoginOrgHandler Logins the handler logic
func makeLoginOrgHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/auth/login-org", http1.NewServer(endpoints.LoginOrgEndpoint, decodeLoginOrgRequest, encodeLoginOrgResponse, options...))
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

// makeInviteHandler creates the handler logic
func makeInviteHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/auth/invite", http1.NewServer(endpoints.InviteEndpoint, decodeInviteRequest, encodeInviteResponse, options...))
}

// decodeInviteResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeInviteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.InviteRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeInviteResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeInviteResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeShowInvitesHandler creates the handler logic
func makeShowInvitesHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/auth/show-invites", http1.NewServer(endpoints.ShowInvitesEndpoint, decodeShowInvitesRequest, encodeShowInvitesResponse, options...))
}

// decodeShowInvitesResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeShowInvitesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ShowInvitesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeShowInvitesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeShowInvitesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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
	m.Handle("/api/v1/auth/show-profile", http1.NewServer(endpoints.ShowProfileEndpoint, decodeShowProfileRequest, encodeShowProfileResponse, options...))
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
