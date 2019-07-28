package service

import (
	"context"

	model "github.com/GDGVIT/Project-Hades/model"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(SimpleProjectionService) SimpleProjectionService

type loggingMiddleware struct {
	logger log.Logger
	next   SimpleProjectionService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a SimpleProjectionService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next SimpleProjectionService) SimpleProjectionService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) ProjectAll(ctx context.Context, event string, query model.Query) (rs []model.Participant, err error) {
	defer func() {
		l.logger.Log("method", "ProjectAll", "event", event, "query", query, "rs", rs, "err", err)
	}()
	return l.next.ProjectAll(ctx, event, query)
}
func (l loggingMiddleware) ProjectPresent(ctx context.Context, event string, day int, query model.Query) (rs []model.Participant, err error) {
	defer func() {
		l.logger.Log("method", "ProjectPresent", "event", event, "day", day, "query", query, "rs", rs, "err", err)
	}()
	return l.next.ProjectPresent(ctx, event, day, query)
}
func (l loggingMiddleware) ProjectAbsent(ctx context.Context, event string, day int, query model.Query) (rs []model.Participant, err error) {
	defer func() {
		l.logger.Log("method", "ProjectAbsent", "event", event, "day", day, "query", query, "rs", rs, "err", err)
	}()
	return l.next.ProjectAbsent(ctx, event, day, query)
}
