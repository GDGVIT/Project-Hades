package endpoint

import (
	"context"
	service "github.com/GDGVIT/Project-Hades/attendance/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// PostAttendanceRequest collects the request parameters for the PostAttendance method.
type PostAttendanceRequest struct {
	Reg       string `json:"reg"`
	Coupons   uint8  `json:"coupons"`
	EventName string `json:"event_name"`
}

// PostAttendanceResponse collects the response parameters for the PostAttendance method.
type PostAttendanceResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakePostAttendanceEndpoint returns an endpoint that invokes PostAttendance on the service.
func MakePostAttendanceEndpoint(s service.AttendanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostAttendanceRequest)
		rs, err := s.PostAttendance(ctx, req.Reg, req.Coupons, req.EventName)
		return PostAttendanceResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r PostAttendanceResponse) Failed() error {
	return r.Err
}

// PostCouponRequest collects the request parameters for the PostCoupon method.
type PostCouponRequest struct {
	Reg    string `json:"reg"`
	Coupon string `json:"coupon"`
}

// PostCouponResponse collects the response parameters for the PostCoupon method.
type PostCouponResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakePostCouponEndpoint returns an endpoint that invokes PostCoupon on the service.
func MakePostCouponEndpoint(s service.AttendanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostCouponRequest)
		rs, err := s.PostCoupon(ctx, req.Reg, req.Coupon)
		return PostCouponResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r PostCouponResponse) Failed() error {
	return r.Err
}

// DeleteCouponRequest collects the request parameters for the DeleteCoupon method.
type DeleteCouponRequest struct {
	Reg       string `json:"reg"`
	EventName string `json:"event_name"`
}

// DeleteCouponResponse collects the response parameters for the DeleteCoupon method.
type DeleteCouponResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteCouponEndpoint returns an endpoint that invokes DeleteCoupon on the service.
func MakeDeleteCouponEndpoint(s service.AttendanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCouponRequest)
		rs, err := s.DeleteCoupon(ctx, req.Reg, req.EventName)
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

// UnpostAttendanceRequest collects the request parameters for the UnpostAttendance method.
type UnpostAttendanceRequest struct {
	Reg       string `json:"reg"`
	EventName string `json:"event_name"`
}

// UnpostAttendanceResponse collects the response parameters for the UnpostAttendance method.
type UnpostAttendanceResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeUnpostAttendanceEndpoint returns an endpoint that invokes UnpostAttendance on the service.
func MakeUnpostAttendanceEndpoint(s service.AttendanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UnpostAttendanceRequest)
		rs, err := s.UnpostAttendance(ctx, req.Reg, req.EventName)
		return UnpostAttendanceResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r UnpostAttendanceResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// PostAttendance implements Service. Primarily useful in a client.
func (e Endpoints) PostAttendance(ctx context.Context, reg string, coupons uint8, eventName string) (rs string, err error) {
	request := PostAttendanceRequest{
		Coupons:   coupons,
		EventName: eventName,
		Reg:       reg,
	}
	response, err := e.PostAttendanceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostAttendanceResponse).Rs, response.(PostAttendanceResponse).Err
}

// PostCoupon implements Service. Primarily useful in a client.
func (e Endpoints) PostCoupon(ctx context.Context, reg string, coupon string) (rs string, err error) {
	request := PostCouponRequest{
		Coupon: coupon,
		Reg:    reg,
	}
	response, err := e.PostCouponEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostCouponResponse).Rs, response.(PostCouponResponse).Err
}

// DeleteCoupon implements Service. Primarily useful in a client.
func (e Endpoints) DeleteCoupon(ctx context.Context, reg string, eventName string) (rs string, err error) {
	request := DeleteCouponRequest{
		EventName: eventName,
		Reg:       reg,
	}
	response, err := e.DeleteCouponEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteCouponResponse).Rs, response.(DeleteCouponResponse).Err
}

// UnpostAttendance implements Service. Primarily useful in a client.
func (e Endpoints) UnpostAttendance(ctx context.Context, reg string, eventName string) (rs string, err error) {
	request := UnpostAttendanceRequest{
		EventName: eventName,
		Reg:       reg,
	}
	response, err := e.UnpostAttendanceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UnpostAttendanceResponse).Rs, response.(UnpostAttendanceResponse).Err
}
