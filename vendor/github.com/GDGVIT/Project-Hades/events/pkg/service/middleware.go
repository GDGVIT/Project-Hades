package service

import (
	"context"

	model "github.com/GDGVIT/Project-Hades/model"
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

func (l loggingMiddleware) CreateEvent(ctx context.Context, event model.Event) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "CreateEvent", "event", event, "rs", rs, "err", err)
	}()
	return l.next.CreateEvent(ctx, event)
}
func (l loggingMiddleware) ReadEvent(ctx context.Context, query model.Query) (rs []model.Event, err error) {
	defer func() {
		l.logger.Log("method", "ReadEvent", "query", query, "rs", rs, "err", err)
	}()
	return l.next.ReadEvent(ctx, query)
}
func (l loggingMiddleware) UpdateEvent(ctx context.Context, query model.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "UpdateEvent", "query", query, "rs", rs, "err", err)
	}()
	return l.next.UpdateEvent(ctx, query)
}
func (l loggingMiddleware) DeleteEvent(ctx context.Context, query model.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "DeleteEvent", "query", query, "rs", rs, "err", err)
	}()
	return l.next.DeleteEvent(ctx, query)
}
