package auth

import "github.com/krisnaadi/dashboard-cronjob-be/internal/usecase/auth"

type Handler struct {
	auth auth.UseCaseProvider
}

// New initializes handler layer.
func New(auth auth.UseCaseProvider) *Handler {
	return &Handler{
		auth: auth,
	}
}
