package cronjob

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/resource/cronjob"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/resource/log"
)

type UseCaseProvider interface {
	ListCronjob(ctx context.Context, UserId int64) ([]entity.Cronjob, error)
	GetCronjob(ctx context.Context, ID int64, UserId int64) (entity.Cronjob, error)
	AddCronjob(ctx context.Context, input CronjobRequest, UserId int64) (entity.Cronjob, error)
	UpdateCronjob(ctx context.Context, ID int64, input CronjobRequest, UserId int64) (entity.Cronjob, error)
	DeleteCronjob(ctx context.Context, ID int64, UserId int64) error
	RunAllCronjob(ctx context.Context) error
}

// UseCase types of useCase layer.
type UseCase struct {
	cronjob cronjob.ResourceProvider
	log     log.ResourceProvider
}

// New initializes useCase layer.
func New(cronjob cronjob.ResourceProvider, log log.ResourceProvider) UseCaseProvider {
	return &UseCase{
		cronjob: cronjob,
		log:     log,
	}
}
