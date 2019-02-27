package service

import (
	"context"

	"github.com/GDGVIT/Project-Hades/model"
)

// CouponsService describes the service.
type CouponsService interface {
	CreateSchema(ctx context.Context, event string, coupons []model.Coupon) (rs string, err error)
	MarkPresent(ctx context.Context, attendance model.Attendance) (rs string, err error)
	RedeemCoupon(ctx context.Context, attendance model.Attendance, couponName string) (rs string, err error)
	DeleteCoupon(ctx context.Context, event string, coupon model.Coupon) (rs string, err error)
	DeleteSchema(ctx context.Context, event string) (rs string, err error)
	ViewSchema(ctx context.Context, event string) (rs []model.Coupon, err error)
}

type basicCouponsService struct{}

func (b *basicCouponsService) CreateSchema(ctx context.Context, event string, coupons []model.Coupon) (rs string, err error) {
	// TODO implement the business logic of CreateSchema
	return rs, err
}
func (b *basicCouponsService) MarkPresent(ctx context.Context, attendance model.Attendance) (rs string, err error) {
	// TODO implement the business logic of MarkPresent
	return rs, err
}
func (b *basicCouponsService) RedeemCoupon(ctx context.Context, attendance model.Attendance, couponName string) (rs string, err error) {
	// TODO implement the business logic of RedeemCoupon
	return rs, err
}
func (b *basicCouponsService) DeleteCoupon(ctx context.Context, event string, coupon model.Coupon) (rs string, err error) {
	// TODO implement the business logic of DeleteCoupon
	return rs, err
}
func (b *basicCouponsService) DeleteSchema(ctx context.Context, event string) (rs string, err error) {
	// TODO implement the business logic of DeleteSchema
	return rs, err
}
func (b *basicCouponsService) ViewSchema(ctx context.Context, event string) (rs []model.Coupon, err error) {
	// TODO implement the business logic of ViewSchema
	return rs, err
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
