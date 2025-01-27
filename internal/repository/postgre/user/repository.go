package user

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"

	"gorm.io/gorm"
)

type RepositoryProvider interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
	GetUserByID(ctx context.Context, ID int64) (entity.User, error)
	InsertUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, ID int64) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryProvider {
	return &Repository{
		db: db,
	}
}
