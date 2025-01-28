package cronjob

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/resource/cronjob"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/resource/log"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/scheduler"
)

type UseCaseProvider interface {
	ListCronjob(ctx context.Context, UserId int64) ([]entity.Cronjob, error)
	GetCronjob(ctx context.Context, ID int64, UserId int64) (entity.Cronjob, error)
	AddCronjob(ctx context.Context, input CronjobRequest, UserId int64) (entity.Cronjob, error)
	UpdateCronjob(ctx context.Context, ID int64, input CronjobRequest, UserId int64) (entity.Cronjob, error)
	DeleteCronjob(ctx context.Context, ID int64, UserId int64) error
	RunAllCronjob(ctx context.Context) error
	RunCronjobManualy(ctx context.Context, ID int64, UserId int64) error
	GetLogByCronjob(ctx context.Context, ID int64) ([]entity.Log, error)
}

// UseCase types of useCase layer.
type UseCase struct {
	cronjob   cronjob.ResourceProvider
	log       log.ResourceProvider
	scheduler *scheduler.Scheduler
}

// New initializes useCase layer.
func New(cronjob cronjob.ResourceProvider, log log.ResourceProvider, scheduler *scheduler.Scheduler) UseCaseProvider {
	return &UseCase{
		cronjob:   cronjob,
		log:       log,
		scheduler: scheduler,
	}
}
