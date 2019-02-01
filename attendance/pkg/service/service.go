package service

import (
	"context"
	"fmt"
	"log"

	"github.com/GDGVIT/Project-Hades/model"
)

// AttendanceService describes the service.
type AttendanceService interface {
	PostAttendance(ctx context.Context, query model.Attendance) (rs string, err error)
	PostCoupon(ctx context.Context, coupon string, query model.Attendance) (rs string, err error)
	DeleteAllCoupons(ctx context.Context, query model.Attendance) (rs string, err error)
	UnpostAttendance(ctx context.Context, query model.Attendance) (rs string, err error)
	ViewCoupons(ctx context.Context, query model.Attendance) (rs []string, err error)
	// ViewPresent(ctx context.Context, query model.Attendance) (rs []model.Participant, err error)
	// ViewAbsent(ctx context.Context, query model.Attendance) (rs []model.Participant, err error)
}

type basicAttendanceService struct{}

/**
*@api {post} /post-attendance mark attendance
*@apiName mark attendance
*@apiGroup attendance
*@apiPermission admin
*
*@apiParam {string} eventName name of the event
*@apiParam {string} email email of the participant
*@apiParam {int} day day of the event
*@apiParam {int} coupons number of coupons for that day
*
*@apiParamExample {json} request-example
*
*{
*
*	"query":{
*		"eventName":"DEVFEST",
*		"day":2,
*		"coupons":5,
*		"email":"angad.sharma2017@vitstudent.ac.in"
*	}
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": "Successfully marked present for the day",
*    "err": null
*}
*
**/
func (b *basicAttendanceService) PostAttendance(ctx context.Context, query model.Attendance) (rs string, err error) {
	c := make(chan model.MessageReturn)
	log.Println(query)
	go model.MarkPresent(query, c)

	ret := <-c
	if err := ret.Err; err != nil {
		return ret.Message, err
	}
	return ret.Message, ret.Err
}

/**
*@api {post} /post-coupon redeem a coupon
*@apiName redeem a coupon
*@apiGroup attendance
*@apiPermission admin
*
*@apiParam {string} eventName name of the event
*@apiParam {string} email email of the participant
*@apiParam {int} day day of the event
*@apiParam {int} coupon the coupon that waas redeemed
*
*@apiParamExample {json} request-example
*
*{
*	"coupon":"$2a$05$uFMPHxzB3IuOYJCZf2TlWutbZWBjLG650ieEbNls4iFBFOYYi.RQK",
*	"query":{
*		"eventName":"DEVFEST 2019",
*		"day":2,
*		"email": "a@a.com"
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
func (b *basicAttendanceService) PostCoupon(ctx context.Context, coupon string, query model.Attendance) (rs string, err error) {
	c := make(chan model.MessageReturn)
	go model.PostCoupon(coupon, query, c)

	ret := <-c
	if err := ret.Err; err != nil {
		return ret.Message, err
	}
	return ret.Message, ret.Err
}
func (b *basicAttendanceService) DeleteAllCoupons(ctx context.Context, query model.Attendance) (rs string, err error) {
	// TODO implement the business logic of DeleteAllCoupons
	return rs, err
}
func (b *basicAttendanceService) UnpostAttendance(ctx context.Context, query model.Attendance) (rs string, err error) {
	// TODO implement the business logic of UnpostAttendance
	return rs, err
}

/**
*@api {post} /post-attendance view coupons
*@apiName view coupons
*@apiGroup attendance
*@apiPermission admin
*
*@apiParam {string} eventName name of the event
*@apiParam {string} email email of the participant
*@apiParam {int} day day of the event
*
*@apiParamExample {json} request-example
*
*{
*
*	"query":{
*		"eventName":"DEVFEST",
*		"day":2,
*		"email":"angad.sharma2017@vitstudent.ac.in"
*	}
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": [
*        "$2a$05$Q6p2bV07hR2Kp02jPbo9UOUtV.FZhRfpQDHApdN.xp/5XmSYS0PGC",
*        "$2a$05$MPED5ZYA4Q9QwK.mZCaskew/hmM/HOvuN2Vx042QIENEoaUAFS6zW",
*        "$2a$05$/iDvNw/11F/xCWMbbRjgMOWWZ1ICJB7crmW2mr0BNI5vElXqPHQL6",
*        "$2a$05$IhZM3URq0VMCT1KO/OBeJOPJiG/XZ5y.AW9nLZVgpewGCuBSap/zC",
*        "$2a$05$YN95rNFZB9Y7o0UOyQHUlu.7iJ6nZSFf0SNWrHjtboQB3LgHkHcLa",
*        "$2a$05$5rAbmkgHkO0VVrcgT8LuwOpty0STOKaTKaEuSYbQtu.YqyM/5jOsa",
*        "$2a$05$KmGC/i4VnZFb9h5789dIcu0jz/v6.HKUXPKlUgBKD7HzpUtsLkgD.",
*        "$2a$05$4QxBr9uahOE6bW6JMHeXBuuF4P3uRIZJCIebaB91oaj9mjjVGdmIW",
*        "$2a$05$IHWoPSkldlbC9YafMTY2w..jw5uCY1zpGCxbMQPFRh/vM4gHCMoFC",
*        "$2a$05$QY6oN1pXcRjU.7XIwcgi/u9/D3iz/RXI2od/oqy5zzNY16tNDvN3W"
*    ],
*    "err": null
*}
*
**/

func (b *basicAttendanceService) ViewCoupons(ctx context.Context, query model.Attendance) (rs []string, err error) {
	rs = model.ViewCoupon(query)
	if rs == nil {
		return nil, fmt.Errorf("No coupons found")
	}
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
