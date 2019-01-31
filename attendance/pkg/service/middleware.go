package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AttendanceService) AttendanceService

type loggingMiddleware struct {
	logger log.Logger
	next   AttendanceService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AttendanceService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AttendanceService) AttendanceService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) PostAttendance(ctx context.Context, reg string, coupons uint8, eventName string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "PostAttendance", "reg", reg, "coupons", coupons, "eventName", eventName, "rs", rs, "err", err)
	}()
	return l.next.PostAttendance(ctx, reg, coupons, eventName)
}
func (l loggingMiddleware) PostCoupon(ctx context.Context, reg string, coupon string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "PostCoupon", "reg", reg, "coupon", coupon, "rs", rs, "err", err)
	}()
	return l.next.PostCoupon(ctx, reg, coupon)
}
func (l loggingMiddleware) DeleteCoupon(ctx context.Context, reg string, eventName string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "DeleteCoupon", "reg", reg, "eventName", eventName, "rs", rs, "err", err)
	}()
	return l.next.DeleteCoupon(ctx, reg, eventName)
}
func (l loggingMiddleware) UnpostAttendance(ctx context.Context, reg string, eventName string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "UnpostAttendance", "reg", reg, "eventName", eventName, "rs", rs, "err", err)
	}()
	return l.next.UnpostAttendance(ctx, reg, eventName)
}
