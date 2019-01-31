package service

import (
	"context"

	"github.com/GDGVIT/Project-Hades/model"
)

// AttendanceService describes the service.
type AttendanceService interface {
	PostAttendance(ctx context.Context, details model.Attendance) (rs string, err error)
	PostCoupon(ctx context.Context, email string, coupon string) (rs string, err error)
	DeleteCoupon(ctx context.Context, email string, event string) (rs string, err error)
	UnpostAttendance(ctx context.Context, email string, event string) (rs string, err error)
	ViewPresent(ctx context.Context, event string) (rs []model.Participant, err error)
	ViewAbsent(ctx context.Context, event string) (rs []model.Participant, err error)
}

type basicAttendanceService struct{}

/**
*@api {post} /post-attendance
*@apiName mark attendance
*@apiGroup attendance
*@apiPermission admin
*
*@apiParam {string} eventName name of the event
*@apiParam {string} registrationNumber registration number of the participant
*@apiParam {int} day day of the event
*@apiParam {int} coupons number of coupons for that day
*
*@apiParamExample {json} request-example
*
*{
*
*	"details":{
*		"eventName":"DEVFEST",
*		"day":2,
*		"coupons":5,
*		"email":"angad.sharma2017@vitstudent.ac.in"
*	}
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": "done",
*    "err": null
*}
*
**/
func (b *basicAttendanceService) PostAttendance(ctx context.Context, details model.Attendance) (rs string, err error) {
	c := make(chan error)
	go model.MarkPresent(details.EventName, details.Email, details.Coupons, details.Day, c)
	if err := <-c; err != nil {
		return "Error marking present.", err
	}
	return "done", err
}
func (b *basicAttendanceService) PostCoupon(ctx context.Context, email string, coupon string) (rs string, err error) {
	// TODO implement the business logic of PostCoupon
	return rs, err
}
func (b *basicAttendanceService) DeleteCoupon(ctx context.Context, email string, event string) (rs string, err error) {
	// TODO implement the business logic of DeleteCoupon
	return rs, err
}
func (b *basicAttendanceService) UnpostAttendance(ctx context.Context, email string, event string) (rs string, err error) {
	// TODO implement the business logic of UnpostAttendance
	return rs, err
}
func (b *basicAttendanceService) ViewPresent(ctx context.Context, event string) (rs []model.Participant, err error) {
	// TODO implement the business logic of ViewPresent
	return rs, err
}
func (b *basicAttendanceService) ViewAbsent(ctx context.Context, event string) (rs []model.Participant, err error) {
	// TODO implement the business logic of ViewAbsent
	return rs, err
}

// NewBasicAttendanceService returns a naive, stateless implementation of AttendanceService.
func NewBasicAttendanceService() AttendanceService {
	return &basicAttendanceService{}
}

// New returns a AttendanceService with all of the expected middleware wired in.
func New(middleware []Middleware) AttendanceService {
	var svc AttendanceService = NewBasicAttendanceService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
