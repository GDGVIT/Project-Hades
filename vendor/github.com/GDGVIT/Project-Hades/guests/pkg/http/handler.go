package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpoint "github.com/GDGVIT/Project-Hades/guests/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
)

// makeCreateGuestHandler creates the handler logic
func makeCreateGuestHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/guests/create-guest", http1.NewServer(endpoints.CreateGuestEndpoint, decodeCreateGuestRequest, encodeCreateGuestResponse, options...))
}

// decodeCreateGuestResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateGuestRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateGuestRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateGuestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateGuestResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadGuestHandler creates the handler logic
func makeReadGuestHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/guests/read-guest", http1.NewServer(endpoints.ReadGuestEndpoint, decodeReadGuestRequest, encodeReadGuestResponse, options...))
}

// decodeReadGuestResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadGuestRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ReadGuestRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReadGuestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadGuestResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateGuestHandler creates the handler logic
func makeUpdateGuestHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/guests/update-guest", http1.NewServer(endpoints.UpdateGuestEndpoint, decodeUpdateGuestRequest, encodeUpdateGuestResponse, options...))
}

// decodeUpdateGuestResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateGuestRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateGuestRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateGuestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateGuestResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteGuestHandler creates the handler logic
func makeDeleteGuestHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/guests/delete-guest", http1.NewServer(endpoints.DeleteGuestEndpoint, decodeDeleteGuestRequest, encodeDeleteGuestResponse, options...))
}

// decodeDeleteGuestResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteGuestRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteGuestRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteGuestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteGuestResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateSponsorHandler creates the handler logic
func makeCreateSponsorHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/guests/create-sponsor", http1.NewServer(endpoints.CreateSponsorEndpoint, decodeCreateSponsorRequest, encodeCreateSponsorResponse, options...))
}

// decodeCreateSponsorResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateSponsorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateSponsorRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateSponsorResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateSponsorResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadSponsorHandler creates the handler logic
func makeReadSponsorHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/guests/read-sponsor", http1.NewServer(endpoints.ReadSponsorEndpoint, decodeReadSponsorRequest, encodeReadSponsorResponse, options...))
}

// decodeReadSponsorResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadSponsorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ReadSponsorRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReadSponsorResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadSponsorResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateSponsorHandler creates the handler logic
func makeUpdateSponsorHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/guests/update-sponsor", http1.NewServer(endpoints.UpdateSponsorEndpoint, decodeUpdateSponsorRequest, encodeUpdateSponsorResponse, options...))
}

// decodeUpdateSponsorResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateSponsorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateSponsorRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateSponsorResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateSponsorResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteSponsorHandler creates the handler logic
func makeDeleteSponsorHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/guests/delete-sponsor", http1.NewServer(endpoints.DeleteSponsorEndpoint, decodeDeleteSponsorRequest, encodeDeleteSponsorResponse, options...))
}

// decodeDeleteSponsorResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteSponsorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteSponsorRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteSponsorResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteSponsorResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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
