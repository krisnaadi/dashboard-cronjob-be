package cronjob

import "github.com/krisnaadi/dashboard-cronjob-be/internal/usecase/cronjob"

type Handler struct {
	cronjob cronjob.UseCaseProvider
}

// New initializes handler layer.
func New(cronjob cronjob.UseCaseProvider) *Handler {
	return &Handler{
		cronjob: cronjob,
	}
}
