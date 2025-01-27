package app

import "github.com/krisnaadi/dashboard-cronjob-be/internal/usecase/cronjob"

type UseCases struct {
	cronjob cronjob.UseCaseProvider
}

// NewUseCase initializes useCase layer.
func NewUseCase(resources *Resources) *UseCases {
	return &UseCases{
		cronjob: cronjob.New(resources.cronjob),
	}
}
