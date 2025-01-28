package cronjob

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"

	"gorm.io/gorm"
)

type RepositoryProvider interface {
	GetCronjobs(ctx context.Context, UserId int64) ([]entity.Cronjob, error)
	GetCronjobByID(ctx context.Context, ID int64, UserId int64) (entity.Cronjob, error)
	InsertCronjob(ctx context.Context, cronjob entity.Cronjob) (entity.Cronjob, error)
	UpdateCronjob(ctx context.Context, cronjob entity.Cronjob) (entity.Cronjob, error)
	DeleteCronjob(ctx context.Context, ID int64, UserId int64) error
	GetAllActiveCronjob(ctx context.Context) ([]entity.Cronjob, error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryProvider {
	return &Repository{
		db: db,
	}
}
