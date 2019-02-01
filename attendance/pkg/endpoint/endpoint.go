package endpoint

import (
	"context"
	service "github.com/GDGVIT/Project-Hades/attendance/pkg/service"
	model "github.com/GDGVIT/Project-Hades/model"
	endpoint "github.com/go-kit/kit/endpoint"
)

// PostAttendanceRequest collects the request parameters for the PostAttendance method.
type PostAttendanceRequest struct {
	Query model.Attendance `json:"query"`
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
		rs, err := s.PostAttendance(ctx, req.Query)
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
	Coupon string           `json:"coupon"`
	Query  model.Attendance `json:"query"`
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
		rs, err := s.PostCoupon(ctx, req.Coupon, req.Query)
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

// DeleteAllCouponsRequest collects the request parameters for the DeleteAllCoupons method.
type DeleteAllCouponsRequest struct {
	Query model.Attendance `json:"query"`
}

// DeleteAllCouponsResponse collects the response parameters for the DeleteAllCoupons method.
type DeleteAllCouponsResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteAllCouponsEndpoint returns an endpoint that invokes DeleteAllCoupons on the service.
func MakeDeleteAllCouponsEndpoint(s service.AttendanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAllCouponsRequest)
		rs, err := s.DeleteAllCoupons(ctx, req.Query)
		return DeleteAllCouponsResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteAllCouponsResponse) Failed() error {
	return r.Err
}

// UnpostAttendanceRequest collects the request parameters for the UnpostAttendance method.
type UnpostAttendanceRequest struct {
	Query model.Attendance `json:"query"`
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
		rs, err := s.UnpostAttendance(ctx, req.Query)
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

// ViewCouponsRequest collects the request parameters for the ViewCoupons method.
type ViewCouponsRequest struct {
	Query model.Attendance `json:"query"`
}

// ViewCouponsResponse collects the response parameters for the ViewCoupons method.
type ViewCouponsResponse struct {
	Rs  []string `json:"rs"`
	Err error    `json:"err"`
}

// MakeViewCouponsEndpoint returns an endpoint that invokes ViewCoupons on the service.
func MakeViewCouponsEndpoint(s service.AttendanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ViewCouponsRequest)
		rs, err := s.ViewCoupons(ctx, req.Query)
		return ViewCouponsResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ViewCouponsResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// PostAttendance implements Service. Primarily useful in a client.
func (e Endpoints) PostAttendance(ctx context.Context, query model.Attendance) (rs string, err error) {
	request := PostAttendanceRequest{Query: query}
	response, err := e.PostAttendanceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostAttendanceResponse).Rs, response.(PostAttendanceResponse).Err
}

// PostCoupon implements Service. Primarily useful in a client.
func (e Endpoints) PostCoupon(ctx context.Context, coupon string, query model.Attendance) (rs string, err error) {
	request := PostCouponRequest{
		Coupon: coupon,
		Query:  query,
	}
	response, err := e.PostCouponEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostCouponResponse).Rs, response.(PostCouponResponse).Err
}

// DeleteAllCoupons implements Service. Primarily useful in a client.
func (e Endpoints) DeleteAllCoupons(ctx context.Context, query model.Attendance) (rs string, err error) {
	request := DeleteAllCouponsRequest{Query: query}
	response, err := e.DeleteAllCouponsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteAllCouponsResponse).Rs, response.(DeleteAllCouponsResponse).Err
}

// UnpostAttendance implements Service. Primarily useful in a client.
func (e Endpoints) UnpostAttendance(ctx context.Context, query model.Attendance) (rs string, err error) {
	request := UnpostAttendanceRequest{Query: query}
	response, err := e.UnpostAttendanceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UnpostAttendanceResponse).Rs, response.(UnpostAttendanceResponse).Err
}

// ViewCoupons implements Service. Primarily useful in a client.
func (e Endpoints) ViewCoupons(ctx context.Context, query model.Attendance) (rs []string, err error) {
	request := ViewCouponsRequest{Query: query}
	response, err := e.ViewCouponsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ViewCouponsResponse).Rs, response.(ViewCouponsResponse).Err
}
