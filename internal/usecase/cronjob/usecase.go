package cronjob

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/resource/cronjob"
)

type UseCaseProvider interface {
	ListCronjob(ctx context.Context) ([]entity.Cronjob, error)
	GetCronjob(ctx context.Context, ID int64) (entity.Cronjob, error)
	AddCronjob(ctx context.Context, input CronjobRequest) (entity.Cronjob, error)
	UpdateCronjob(ctx context.Context, ID int64, input CronjobRequest) (entity.Cronjob, error)
	DeleteCronjob(ctx context.Context, ID int64) error
}

// UseCase types of useCase layer.
type UseCase struct {
	cronjob cronjob.ResourceProvider
}

// New initializes useCase layer.
func New(cronjob cronjob.ResourceProvider) UseCaseProvider {
	return &UseCase{
		cronjob: cronjob,
	}
}
