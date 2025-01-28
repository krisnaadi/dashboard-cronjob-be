package app

import (
	"github.com/krisnaadi/dashboard-cronjob-be/internal/usecase/auth"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/usecase/cronjob"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/scheduler"
)

type UseCases struct {
	cronjob cronjob.UseCaseProvider
	auth    auth.UseCaseProvider
}

// NewUseCase initializes useCase layer.
func NewUseCase(resources *Resources, scheduler *scheduler.Scheduler) *UseCases {
	return &UseCases{
		cronjob: cronjob.New(resources.cronjob, resources.log, scheduler),
		auth:    auth.New(resources.user),
	}
}
