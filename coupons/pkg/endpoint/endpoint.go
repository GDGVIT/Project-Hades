package endpoint

import (
	"context"

	service "github.com/GDGVIT/Project-Hades/coupons/pkg/service"
	model "github.com/GDGVIT/Project-Hades/model"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateSchemaRequest collects the request parameters for the CreateSchema method.
type CreateSchemaRequest struct {
	Event   string         `json:"event"`
	Coupons []model.Coupon `json:"coupons"`
}

// CreateSchemaResponse collects the response parameters for the CreateSchema method.
type CreateSchemaResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateSchemaEndpoint returns an endpoint that invokes CreateSchema on the service.
func MakeCreateSchemaEndpoint(s service.CouponsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateSchemaRequest)
		rs, err := s.CreateSchema(ctx, req.Event, req.Coupons)
		return CreateSchemaResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateSchemaResponse) Failed() error {
	return r.Err
}

// MarkPresentRequest collects the request parameters for the MarkPresent method.
type MarkPresentRequest struct {
	Attendance model.Attendance `json:"attendance"`
}

// MarkPresentResponse collects the response parameters for the MarkPresent method.
type MarkPresentResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeMarkPresentEndpoint returns an endpoint that invokes MarkPresent on the service.
func MakeMarkPresentEndpoint(s service.CouponsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MarkPresentRequest)
		rs, err := s.MarkPresent(ctx, req.Attendance)
		return MarkPresentResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r MarkPresentResponse) Failed() error {
	return r.Err
}

// RedeemCouponRequest collects the request parameters for the RedeemCoupon method.
type RedeemCouponRequest struct {
	Attendance model.Attendance `json:"attendance"`
	CouponName string           `json:"coupon_name"`
}

// RedeemCouponResponse collects the response parameters for the RedeemCoupon method.
type RedeemCouponResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeRedeemCouponEndpoint returns an endpoint that invokes RedeemCoupon on the service.
func MakeRedeemCouponEndpoint(s service.CouponsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RedeemCouponRequest)
		rs, err := s.RedeemCoupon(ctx, req.Attendance, req.CouponName)
		return RedeemCouponResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r RedeemCouponResponse) Failed() error {
	return r.Err
}

// DeleteCouponRequest collects the request parameters for the DeleteCoupon method.
type DeleteCouponRequest struct {
	Event  string       `json:"event"`
	Coupon model.Coupon `json:"coupon"`
}

// DeleteCouponResponse collects the response parameters for the DeleteCoupon method.
type DeleteCouponResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteCouponEndpoint returns an endpoint that invokes DeleteCoupon on the service.
func MakeDeleteCouponEndpoint(s service.CouponsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCouponRequest)
		rs, err := s.DeleteCoupon(ctx, req.Event, req.Coupon)
		return DeleteCouponResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteCouponResponse) Failed() error {
	return r.Err
}

// DeleteSchemaRequest collects the request parameters for the DeleteSchema method.
type DeleteSchemaRequest struct {
	Event string       `json:"event"`
	Query model.Coupon `json:"query"`
}

// DeleteSchemaResponse collects the response parameters for the DeleteSchema method.
type DeleteSchemaResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteSchemaEndpoint returns an endpoint that invokes DeleteSchema on the service.
func MakeDeleteSchemaEndpoint(s service.CouponsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteSchemaRequest)
		rs, err := s.DeleteSchema(ctx, req.Event, req.Query)
		return DeleteSchemaResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteSchemaResponse) Failed() error {
	return r.Err
}

// ViewSchemaRequest collects the request parameters for the ViewSchema method.
type ViewSchemaRequest struct {
	Event string `json:"event"`
}

// ViewSchemaResponse collects the response parameters for the ViewSchema method.
type ViewSchemaResponse struct {
	Rs  []model.Coupon `json:"rs"`
	Err error          `json:"err"`
}

// MakeViewSchemaEndpoint returns an endpoint that invokes ViewSchema on the service.
func MakeViewSchemaEndpoint(s service.CouponsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ViewSchemaRequest)
		rs, err := s.ViewSchema(ctx, req.Event)
		return ViewSchemaResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ViewSchemaResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateSchema implements Service. Primarily useful in a client.
func (e Endpoints) CreateSchema(ctx context.Context, event string, coupons []model.Coupon) (rs string, err error) {
	request := CreateSchemaRequest{
		Coupons: coupons,
		Event:   event,
	}
	response, err := e.CreateSchemaEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateSchemaResponse).Rs, response.(CreateSchemaResponse).Err
}

// MarkPresent implements Service. Primarily useful in a client.
func (e Endpoints) MarkPresent(ctx context.Context, attendance model.Attendance) (rs string, err error) {
	request := MarkPresentRequest{Attendance: attendance}
	response, err := e.MarkPresentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(MarkPresentResponse).Rs, response.(MarkPresentResponse).Err
}

// RedeemCoupon implements Service. Primarily useful in a client.
func (e Endpoints) RedeemCoupon(ctx context.Context, attendance model.Attendance, couponName string) (rs string, err error) {
	request := RedeemCouponRequest{
		Attendance: attendance,
		CouponName: couponName,
	}
	response, err := e.RedeemCouponEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RedeemCouponResponse).Rs, response.(RedeemCouponResponse).Err
}

// DeleteCoupon implements Service. Primarily useful in a client.
func (e Endpoints) DeleteCoupon(ctx context.Context, event string, coupon model.Coupon) (rs string, err error) {
	request := DeleteCouponRequest{
		Coupon: coupon,
		Event:  event,
	}
	response, err := e.DeleteCouponEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteCouponResponse).Rs, response.(DeleteCouponResponse).Err
}

// DeleteSchema implements Service. Primarily useful in a client.
func (e Endpoints) DeleteSchema(ctx context.Context, event string, query model.Coupon) (rs string, err error) {
	request := DeleteSchemaRequest{Event: event, Query: query}
	response, err := e.DeleteSchemaEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteSchemaResponse).Rs, response.(DeleteSchemaResponse).Err
}

// ViewSchema implements Service. Primarily useful in a client.
func (e Endpoints) ViewSchema(ctx context.Context, event string) (rs []model.Coupon, err error) {
	request := ViewSchemaRequest{Event: event}
	response, err := e.ViewSchemaEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ViewSchemaResponse).Rs, response.(ViewSchemaResponse).Err
}
