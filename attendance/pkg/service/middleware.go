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

func (l loggingMiddleware) PostAttendance(ctx context.Context, details model.Attendance) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "PostAttendance", "details", details, "rs", rs, "err", err)
	}()
	return l.next.PostAttendance(ctx, details)
}
func (l loggingMiddleware) PostCoupon(ctx context.Context, reg string, coupon string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "PostCoupon", "reg", reg, "coupon", coupon, "rs", rs, "err", err)
	}()
	return l.next.PostCoupon(ctx, reg, coupon)
}
func (l loggingMiddleware) DeleteCoupon(ctx context.Context, reg string, event string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "DeleteCoupon", "reg", reg, "event", event, "rs", rs, "err", err)
	}()
	return l.next.DeleteCoupon(ctx, reg, event)
}
func (l loggingMiddleware) UnpostAttendance(ctx context.Context, reg string, event string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "UnpostAttendance", "reg", reg, "event", event, "rs", rs, "err", err)
	}()
	return l.next.UnpostAttendance(ctx, reg, event)
}
func (l loggingMiddleware) ViewPresent(ctx context.Context, event string) (rs []model.Participant, err error) {
	defer func() {
		l.logger.Log("method", "ViewPresent", "event", event, "rs", rs, "err", err)
	}()
	return l.next.ViewPresent(ctx, event)
}
func (l loggingMiddleware) ViewAbsent(ctx context.Context, event string) (rs []model.Participant, err error) {
	defer func() {
		l.logger.Log("method", "ViewAbsent", "event", event, "rs", rs, "err", err)
	}()
	return l.next.ViewAbsent(ctx, event)
}
