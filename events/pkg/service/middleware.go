package service

import (
	"context"
	omegadbconfig "github.com/angadsharma1016/omega_dbconfig"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(EventsService) EventsService

type loggingMiddleware struct {
	logger log.Logger
	next   EventsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a EventsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next EventsService) EventsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateEvent(ctx context.Context, s omegadbconfig.Event) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "CreateEvent", "s", s, "rs", rs, "err", err)
	}()
	return l.next.CreateEvent(ctx, s)
}
func (l loggingMiddleware) ReadEvent(ctx context.Context, s omegadbconfig.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "ReadEvent", "s", s, "rs", rs, "err", err)
	}()
	return l.next.ReadEvent(ctx, s)
}
func (l loggingMiddleware) UpdateEvent(ctx context.Context, s omegadbconfig.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "UpdateEvent", "s", s, "rs", rs, "err", err)
	}()
	return l.next.UpdateEvent(ctx, s)
}
func (l loggingMiddleware) DeleteEvent(ctx context.Context, s omegadbconfig.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "DeleteEvent", "s", s, "rs", rs, "err", err)
	}()
	return l.next.DeleteEvent(ctx, s)
}
