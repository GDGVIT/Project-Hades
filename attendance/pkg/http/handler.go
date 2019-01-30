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

// makeDeleteCouponHandler creates the handler logic
func makeDeleteCouponHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-coupon", http1.NewServer(endpoints.DeleteCouponEndpoint, decodeDeleteCouponRequest, encodeDeleteCouponResponse, options...))
}

// decodeDeleteCouponResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteCouponRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteCouponRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteCouponResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteCouponResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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

// makeViewPresentHandler creates the handler logic
func makeViewPresentHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/view-present", http1.NewServer(endpoints.ViewPresentEndpoint, decodeViewPresentRequest, encodeViewPresentResponse, options...))
}

// decodeViewPresentResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeViewPresentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ViewPresentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeViewPresentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeViewPresentResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeViewAbsentHandler creates the handler logic
func makeViewAbsentHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/view-absent", http1.NewServer(endpoints.ViewAbsentEndpoint, decodeViewAbsentRequest, encodeViewAbsentResponse, options...))
}

// decodeViewAbsentResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeViewAbsentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ViewAbsentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeViewAbsentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeViewAbsentResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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
