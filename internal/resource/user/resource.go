package user

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/repository/postgre/user"
)

type ResourceProvider interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
	GetUserByID(ctx context.Context, ID int64) (entity.User, error)
	AddUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, ID int64, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, ID int64) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

type Resource struct {
	user user.RepositoryProvider
}

func New(user user.RepositoryProvider) ResourceProvider {
	resource := Resource{}

	resource.user = user

	return &resource
}
