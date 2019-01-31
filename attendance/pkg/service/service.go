package service

import (
	"context"
)

// AttendanceService describes the service.
type AttendanceService interface {
	PostAttendance(ctx context.Context, reg string, coupons uint8, eventName string) (rs string, err error)
	PostCoupon(ctx context.Context, reg string, coupon string) (rs string, err error)
	DeleteCoupon(ctx context.Context, reg string, eventName string) (rs string, err error)
	UnpostAttendance(ctx context.Context, reg string, eventName string) (rs string, err error)
	// ViewPresent(ctx context.Context, eventName string) (rs []model.Participant, err error)
	// ViewAbsent(ctx context.Context, eventName string) (rs []model.Participant, err error)
}

type basicAttendanceService struct{}

func (b *basicAttendanceService) PostAttendance(ctx context.Context, reg string, coupons uint8, eventName string) (rs string, err error) {
	// TODO implement the business logic of PostAttendance
	return rs, err
}
func (b *basicAttendanceService) PostCoupon(ctx context.Context, reg string, coupon string) (rs string, err error) {
	// TODO implement the business logic of PostCoupon
	return rs, err
}
func (b *basicAttendanceService) DeleteCoupon(ctx context.Context, reg string, eventName string) (rs string, err error) {
	// TODO implement the business logic of DeleteCoupon
	return rs, err
}
func (b *basicAttendanceService) UnpostAttendance(ctx context.Context, reg string, eventName string) (rs string, err error) {
	// TODO implement the business logic of UnpostAttendance
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
