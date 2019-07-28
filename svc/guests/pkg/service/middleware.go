package service

import (
	"context"
	model "github.com/GDGVIT/Project-Hades/model"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(GuestsService) GuestsService

type loggingMiddleware struct {
	logger log.Logger
	next   GuestsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a GuestsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next GuestsService) GuestsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateGuest(ctx context.Context, event string, guest model.Guest) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "CreateGuest", "event", event, "guest", guest, "rs", rs, "err", err)
	}()
	return l.next.CreateGuest(ctx, event, guest)
}
func (l loggingMiddleware) ReadGuest(ctx context.Context, query model.Query) (rs []model.Guest, err error) {
	defer func() {
		l.logger.Log("method", "ReadGuest", "query", query, "rs", rs, "err", err)
	}()
	return l.next.ReadGuest(ctx, query)
}
func (l loggingMiddleware) UpdateGuest(ctx context.Context, query model.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "UpdateGuest", "query", query, "rs", rs, "err", err)
	}()
	return l.next.UpdateGuest(ctx, query)
}
func (l loggingMiddleware) DeleteGuest(ctx context.Context, query model.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "DeleteGuest", "query", query, "rs", rs, "err", err)
	}()
	return l.next.DeleteGuest(ctx, query)
}
func (l loggingMiddleware) CreateSponsor(ctx context.Context, event string, sponsor model.Participant) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "CreateSponsor", "event", event, "sponsor", sponsor, "rs", rs, "err", err)
	}()
	return l.next.CreateSponsor(ctx, event, sponsor)
}
func (l loggingMiddleware) ReadSponsor(ctx context.Context, query model.Query) (rs []model.Participant, err error) {
	defer func() {
		l.logger.Log("method", "ReadSponsor", "query", query, "rs", rs, "err", err)
	}()
	return l.next.ReadSponsor(ctx, query)
}
func (l loggingMiddleware) UpdateSponsor(ctx context.Context, query model.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "UpdateSponsor", "query", query, "rs", rs, "err", err)
	}()
	return l.next.UpdateSponsor(ctx, query)
}
func (l loggingMiddleware) DeleteSponsor(ctx context.Context, query model.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "DeleteSponsor", "query", query, "rs", rs, "err", err)
	}()
	return l.next.DeleteSponsor(ctx, query)
}
