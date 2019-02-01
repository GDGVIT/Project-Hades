package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "github.com/GDGVIT/Project-Hades/attendance/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makePostAttendanceHandler creates the handler logic
func makePostAttendanceHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/post-attendance", http1.NewServer(endpoints.PostAttendanceEndpoint, decodePostAttendanceRequest, encodePostAttendanceResponse, options...))
}

// decodePostAttendanceResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostAttendanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.PostAttendanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostAttendanceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostAttendanceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostCouponHandler creates the handler logic
func makePostCouponHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/post-coupon", http1.NewServer(endpoints.PostCouponEndpoint, decodePostCouponRequest, encodePostCouponResponse, options...))
}

// decodePostCouponResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostCouponRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.PostCouponRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostCouponResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostCouponResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteAllCouponsHandler creates the handler logic
func makeDeleteAllCouponsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-all-coupons", http1.NewServer(endpoints.DeleteAllCouponsEndpoint, decodeDeleteAllCouponsRequest, encodeDeleteAllCouponsResponse, options...))
}

// decodeDeleteAllCouponsResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteAllCouponsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteAllCouponsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteAllCouponsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteAllCouponsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUnpostAttendanceHandler creates the handler logic
func makeUnpostAttendanceHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/unpost-attendance", http1.NewServer(endpoints.UnpostAttendanceEndpoint, decodeUnpostAttendanceRequest, encodeUnpostAttendanceResponse, options...))
}

// decodeUnpostAttendanceResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUnpostAttendanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UnpostAttendanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUnpostAttendanceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUnpostAttendanceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeViewCouponsHandler creates the handler logic
func makeViewCouponsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/view-coupons", http1.NewServer(endpoints.ViewCouponsEndpoint, decodeViewCouponsRequest, encodeViewCouponsResponse, options...))
}

// decodeViewCouponsResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeViewCouponsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ViewCouponsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeViewCouponsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeViewCouponsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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
