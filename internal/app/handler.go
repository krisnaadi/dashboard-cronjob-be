package app

import (
	"github.com/krisnaadi/dashboard-cronjob-be/internal/handler/auth"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/handler/cronjob"
)

// Handlers types of handler layer.
type Handlers struct {
	Cronjob cronjob.Handler
	Auth    auth.Handler
}

// New initializes handler layer.
func NewHandler(useCase *UseCases) *Handlers {
	return &Handlers{
		Cronjob: *cronjob.New(useCase.cronjob),
		Auth:    *auth.New(useCase.auth),
	}
}
