package cronjob

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/repository/postgre/cronjob"
)

type ResourceProvider interface {
	GetCronjobs(ctx context.Context, UserId int64) ([]entity.Cronjob, error)
	GetCronjobByID(ctx context.Context, ID int64, UserId int64) (entity.Cronjob, error)
	AddCronjob(ctx context.Context, cronjob entity.Cronjob) (entity.Cronjob, error)
	UpdateCronjob(ctx context.Context, ID int64, cronjob entity.Cronjob) (entity.Cronjob, error)
	DeleteCronjob(ctx context.Context, ID int64, UserId int64) error
	GetAllActiveCronjob(ctx context.Context) ([]entity.Cronjob, error)
}

type Resource struct {
	cronjob cronjob.RepositoryProvider
}

func New(cronjob cronjob.RepositoryProvider) ResourceProvider {
	resource := Resource{}

	resource.cronjob = cronjob

	return &resource
}
