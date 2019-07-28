package service

import (
	"context"
	"sync"

	"github.com/L04DB4L4NC3R/Project-Hades/model"
)

// CouponsService describes the service.
type CouponsService interface {
	CreateSchema(ctx context.Context, event string, coupons []model.Coupon) (rs string, err error)
	MarkPresent(ctx context.Context, attendance model.Attendance) (rs string, err error)
	RedeemCoupon(ctx context.Context, attendance model.Attendance, couponName string) (rs string, err error)
	DeleteCoupon(ctx context.Context, event string, coupon model.Coupon) (rs string, err error)
	DeleteSchema(ctx context.Context, event string, query model.Coupon) (rs string, err error)
	ViewSchema(ctx context.Context, event string) (rs []model.Coupon, err error)
}

type basicCouponsService struct{}

/**
*@api {post} /api/v1/coupons/create-schema create coupon schema
*@apiName create coupon schema
*@apiGroup coupons
*@apiPermission admin
*
*@apiParam {string} event name of the event
*@apiParam {string} name name of the coupon
*@apiParam {string} description description of the coupon
*@apiParam {int} day day of the event
*
*@apiParamExample {json} request-example
*
*{
*	"event":"DEVRELCONF",
*	"coupons": [{
*		"name":"lunch",
*		"description":"lunch",
*		"day":1
*	},{
*		"name":"bf",
*		"description":"bf",
*		"day":1
*	},{
*		"name":"lunch",
*		"description":"lunch",
*		"day":2
*	},{
*		"name":"dinner",
*		"description":"dinner",
*		"day":2
*	}]
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": "Successfully creted coupon schema for event DEVRELCONF",
*    "err": null
*}
*
**/
func (b *basicCouponsService) CreateSchema(ctx context.Context, event string, coupons []model.Coupon) (rs string, err error) {

	c := make(chan error)
	go model.CreateCouponSchema(event, coupons, c)
	if err := <-c; err != nil {
		return "Some error occurred", err
	}
	return "Successfully creted coupon schema for event " + event, nil
}

/**
*@api {post} /api/v1/coupons/mark-present mark attendee present
*@apiName mark attendee present
*@apiGroup coupons
*@apiPermission admin
*
*@apiParam {string} eventName name of the event
*@apiParam {string} email email of the attendee
*@apiParam {int} day day of the event
*
*@apiParamExample {json} request-example
*{
*	"attendance":{
*		"eventName":"DEVRELCONF",
*		"email":"a@a.COM",
*		"day":2
*	}
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": "Successfully marked present for the day",
*    "err": null
*}
**/
func (b *basicCouponsService) MarkPresent(ctx context.Context, attendance model.Attendance) (rs string, err error) {
	c := make(chan model.MessageReturn)
	go model.MarkPresent(attendance, c)
	msg := <-c
	if err := msg.Err; err != nil {
		return msg.Message, err
	}
	return msg.Message, nil
}

/**
*@api {post} /api/v1/coupons/redeem-coupon redeem a coupon
*@apiName redeem a coupon
*@apiGroup coupons
*@apiPermission admin
*
*@apiParam {string} eventName name of the event
*@apiParam {string} email email of the attendee
*@apiParam {int} day day of the event
*@apiParam {string} couponName name of the coupon
*
*@apiParamExample {json} request-example
*{
*	"attendance": {
*		"couponName":"dinner",
*		"day":2,
*		"email":"a@a.COM",
*		"eventName":"DEVRELCONF"
*	}
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": "No match found for this coupon",
*    "err": null
*}
**/
func (b *basicCouponsService) RedeemCoupon(ctx context.Context, attendance model.Attendance, couponName string) (rs string, err error) {
	c := make(chan model.MessageReturn)
	go model.RedeemCoupon(attendance, attendance.CouponName, c)
	msg := <-c
	if err := msg.Err; err != nil {
		return msg.Message, msg.Err
	}
	return msg.Message, nil
}

func (b *basicCouponsService) DeleteCoupon(ctx context.Context, event string, coupon model.Coupon) (rs string, err error) {
	// TODO implement the business logic of DeleteCoupon
	return rs, err
}

/**
*@api {post} /api/v1/coupons/delete-coupon delete coupon
*@apiName delete coupon
*@apiGroup coupons
*@apiPermission admin
*
*@apiParam {string} eventName name of the event
*@apiParam {string} name name of the coupon
*@apiParam {int} day day of the event
*@apiParam {string} description description of the coupon
*
*@apiParamExample {json} delete-specific-request
*{
*	"event":"DEVRELCONF",
*	"query":{
*		"name":"lunch",
*		"day":2,
*		"description":"lunch"
*	}
*}
*
*@apiParamExample {json} delete-specific-response
*{
*    "rs": "Deleted",
*    "err": null
*}
*
*@apiParamExample {json} delete-all-request
*{
*	"event":"DEVRELCONF"
*}
*
*@apiParamExample {json} delete-all-response
*{
*    "rs": "Deleted",
*    "err": null
*}
**/
func (b *basicCouponsService) DeleteSchema(ctx context.Context, event string, query model.Coupon) (rs string, err error) {
	c := make(chan model.MessageReturn)
	go model.DeleteCouponSchema(event, query, c)
	msg := <-c
	if err := msg.Err; err != nil {
		return msg.Message, msg.Err
	}
	return msg.Message, nil
}

/**
*@api {post} /api/v1/coupons/view-schema view coupon schema
*@apiName view coupon schema
*@apiGroup coupons
*@apiPermission admin
*
*@apiParam {string} event name of the event

*@apiParamExample {json} request-example
*{
*	"event":"DEVRELCONF"
*}
*
*@apiParamExample {json} response-example
*{
*    "rs": [
*        {
*            "name": "gargentaul",
*            "description": "dasd",
*            "day": 1
*        },
*        {
*            "name": "bf",
*            "description": "bf",
*            "day": 1
*        },
*        {
*            "name": "dinner",
*            "description": "dinner",
*            "day": 2
*        },
*        {
*            "name": "lunch",
*            "description": "lunch",
*            "day": 2
*        },
*        {
*            "name": "lunch",
*            "description": "lunch",
*            "day": 1
*        }
*    ],
*    "err": null
*}
**/
func (b *basicCouponsService) ViewSchema(ctx context.Context, event string) (rs []model.Coupon, err error) {
	mutex := &sync.Mutex{}
	c := make(chan model.CouponReturn)
	go model.ViewCouponSchema(event, mutex, c)
	msg := <-c
	if err := msg.Err; err != nil {
		return nil, err
	}
	return msg.Coupons, nil
}

// NewBasicCouponsService returns a naive, stateless implementation of CouponsService.
func NewBasicCouponsService() CouponsService {
	return &basicCouponsService{}
}

// New returns a CouponsService with all of the expected middleware wired in.
func New(middleware []Middleware) CouponsService {
	var svc CouponsService = NewBasicCouponsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
