package app

import (
	"github.com/krisnaadi/dashboard-cronjob-be/internal/middleware/httpwrapper"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/middleware/logrequest"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/middleware/panichandler"
)

type Middleware struct {
	HttpWrapper  *httpwrapper.Middleware
	LogRequest   *logrequest.Middleware
	PanicHandler *panichandler.Middleware
}

// NewMiddleware initializes middleware
func NewMiddleware(useCase *UseCases) *Middleware {
	return &Middleware{
		HttpWrapper:  httpwrapper.New(),
		LogRequest:   logrequest.New(),
		PanicHandler: panichandler.New(),
	}
}
