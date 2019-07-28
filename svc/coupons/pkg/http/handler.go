package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpoint "github.com/GDGVIT/Project-Hades/coupons/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
)

// makeCreateSchemaHandler creates the handler logic
func makeCreateSchemaHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/coupons/create-schema", http1.NewServer(endpoints.CreateSchemaEndpoint, decodeCreateSchemaRequest, encodeCreateSchemaResponse, options...))
}

// decodeCreateSchemaResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateSchemaRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateSchemaRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateSchemaResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateSchemaResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeMarkPresentHandler creates the handler logic
func makeMarkPresentHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/coupons/mark-present", http1.NewServer(endpoints.MarkPresentEndpoint, decodeMarkPresentRequest, encodeMarkPresentResponse, options...))
}

// decodeMarkPresentResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeMarkPresentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.MarkPresentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeMarkPresentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeMarkPresentResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeRedeemCouponHandler creates the handler logic
func makeRedeemCouponHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/coupons/redeem-coupon", http1.NewServer(endpoints.RedeemCouponEndpoint, decodeRedeemCouponRequest, encodeRedeemCouponResponse, options...))
}

// decodeRedeemCouponResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRedeemCouponRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.RedeemCouponRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRedeemCouponResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRedeemCouponResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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
	m.Handle("/api/v1/coupons/delete-coupon", http1.NewServer(endpoints.DeleteCouponEndpoint, decodeDeleteCouponRequest, encodeDeleteCouponResponse, options...))
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

// makeDeleteSchemaHandler creates the handler logic
func makeDeleteSchemaHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/coupons/delete-schema", http1.NewServer(endpoints.DeleteSchemaEndpoint, decodeDeleteSchemaRequest, encodeDeleteSchemaResponse, options...))
}

// decodeDeleteSchemaResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteSchemaRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteSchemaRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteSchemaResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteSchemaResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeViewSchemaHandler creates the handler logic
func makeViewSchemaHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/api/v1/coupons/view-schema", http1.NewServer(endpoints.ViewSchemaEndpoint, decodeViewSchemaRequest, encodeViewSchemaResponse, options...))
}

// decodeViewSchemaResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeViewSchemaRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ViewSchemaRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeViewSchemaResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeViewSchemaResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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
