package service

import (
	"context"
	model "github.com/GDGVIT/Project-Hades/model"
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

func (l loggingMiddleware) PostAttendance(ctx context.Context, query model.Attendance) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "PostAttendance", "query", query, "rs", rs, "err", err)
	}()
	return l.next.PostAttendance(ctx, query)
}
func (l loggingMiddleware) PostCoupon(ctx context.Context, coupon string, query model.Attendance) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "PostCoupon", "coupon", coupon, "query", query, "rs", rs, "err", err)
	}()
	return l.next.PostCoupon(ctx, coupon, query)
}
func (l loggingMiddleware) DeleteAllCoupons(ctx context.Context, query model.Attendance) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "DeleteAllCoupons", "query", query, "rs", rs, "err", err)
	}()
	return l.next.DeleteAllCoupons(ctx, query)
}
func (l loggingMiddleware) UnpostAttendance(ctx context.Context, query model.Attendance) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "UnpostAttendance", "query", query, "rs", rs, "err", err)
	}()
	return l.next.UnpostAttendance(ctx, query)
}
func (l loggingMiddleware) ViewCoupons(ctx context.Context, query model.Attendance) (rs []string, err error) {
	defer func() {
		l.logger.Log("method", "ViewCoupons", "query", query, "rs", rs, "err", err)
	}()
	return l.next.ViewCoupons(ctx, query)
}
