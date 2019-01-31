package endpoint

import (
	"context"

	service "github.com/GDGVIT/Project-Hades/attendance/pkg/service"
	"github.com/GDGVIT/Project-Hades/model"
	endpoint "github.com/go-kit/kit/endpoint"
)

// PostAttendanceRequest collects the request parameters for the PostAttendance method.
type PostAttendanceRequest struct {
	Reg       string `json:"reg"`
	Coupons   int    `json:"coupons"`
	EventName string `json:"event_name"`
	Day       int    `json:"day"`
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
		rs, err := s.PostAttendance(ctx, req.Reg, req.Coupons, req.EventName, req.Day)
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

// ViewPresentRequest collects the request parameters for the ViewPresent method.
type ViewPresentRequest struct {
	EventName string `json:"event_name"`
}

// ViewPresentResponse collects the response parameters for the ViewPresent method.
type ViewPresentResponse struct {
	Rs  []model.Participant `json:"rs"`
	Err error               `json:"err"`
}

// MakeViewPresentEndpoint returns an endpoint that invokes ViewPresent on the service.
func MakeViewPresentEndpoint(s service.AttendanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ViewPresentRequest)
		rs, err := s.ViewPresent(ctx, req.EventName)
		return ViewPresentResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ViewPresentResponse) Failed() error {
	return r.Err
}

// ViewAbsentRequest collects the request parameters for the ViewAbsent method.
type ViewAbsentRequest struct {
	EventName string `json:"event_name"`
}

// ViewAbsentResponse collects the response parameters for the ViewAbsent method.
type ViewAbsentResponse struct {
	Rs  []model.Participant `json:"rs"`
	Err error               `json:"err"`
}

// MakeViewAbsentEndpoint returns an endpoint that invokes ViewAbsent on the service.
func MakeViewAbsentEndpoint(s service.AttendanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ViewAbsentRequest)
		rs, err := s.ViewAbsent(ctx, req.EventName)
		return ViewAbsentResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ViewAbsentResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// PostAttendance implements Service. Primarily useful in a client.
func (e Endpoints) PostAttendance(ctx context.Context, reg string, coupons int, eventName string, day int) (rs string, err error) {
	request := PostAttendanceRequest{
		Coupons:   coupons,
		Day:       day,
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

// ViewPresent implements Service. Primarily useful in a client.
func (e Endpoints) ViewPresent(ctx context.Context, eventName string) (rs []model.Participant, err error) {
	request := ViewPresentRequest{EventName: eventName}
	response, err := e.ViewPresentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ViewPresentResponse).Rs, response.(ViewPresentResponse).Err
}

// ViewAbsent implements Service. Primarily useful in a client.
func (e Endpoints) ViewAbsent(ctx context.Context, eventName string) (rs []model.Participant, err error) {
	request := ViewAbsentRequest{EventName: eventName}
	response, err := e.ViewAbsentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ViewAbsentResponse).Rs, response.(ViewAbsentResponse).Err
}
