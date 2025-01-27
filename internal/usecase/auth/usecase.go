package auth

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/resource/user"
)

type UseCaseProvider interface {
	GetAuthUser(ctx context.Context, ID int64) (entity.User, error)
	Login(ctx context.Context, input LoginRequest) (entity.User, error)
	Register(ctx context.Context, input RegisterRequest) (entity.User, error)
	GenerateToken(ctx context.Context, user entity.User) (string, error)
}

// UseCase types of useCase layer.
type UseCase struct {
	user user.ResourceProvider
}

// New initializes useCase layer.
func New(user user.ResourceProvider) UseCaseProvider {
	return &UseCase{
		user: user,
	}
}
