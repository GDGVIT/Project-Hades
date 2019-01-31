package service

import (
	"context"

	"github.com/GDGVIT/Project-Hades/model"
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

func (l loggingMiddleware) PostAttendance(ctx context.Context, reg string, coupons int, eventName string, day int) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "PostAttendance", "reg", reg, "coupons", coupons, "eventName", eventName, "day", day, "rs", rs, "err", err)
	}()
	return l.next.PostAttendance(ctx, reg, coupons, eventName, day)
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
func (l loggingMiddleware) ViewPresent(ctx context.Context, eventName string) (rs []model.Participant, err error) {
	defer func() {
		l.logger.Log("method", "ViewPresent", "eventName", eventName, "rs", rs, "err", err)
	}()
	return l.next.ViewPresent(ctx, eventName)
}
func (l loggingMiddleware) ViewAbsent(ctx context.Context, eventName string) (rs []model.Participant, err error) {
	defer func() {
		l.logger.Log("method", "ViewAbsent", "eventName", eventName, "rs", rs, "err", err)
	}()
	return l.next.ViewAbsent(ctx, eventName)
}
