package cronjob

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
)

func (resource *Resource) GetCronjobByID(ctx context.Context, ID int64, UserId int64) (entity.Cronjob, error) {
	cronjob, err := resource.cronjob.GetCronjobByID(ctx, ID, UserId)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "resource.cronjob.GetCronjobByID() error - GetCronjobByID")
		return cronjob, err
	}

	return cronjob, nil
}

func (resource *Resource) GetCronjobs(ctx context.Context, UserId int64) ([]entity.Cronjob, error) {
	cronjobs, err := resource.cronjob.GetCronjobs(ctx, UserId)
	if err != nil {
		logger.Trace(ctx, nil, err, "resource.cronjob.GetCronjobs() error - GetCronjobs")
		return cronjobs, err
	}

	return cronjobs, nil
}

func (resource *Resource) AddCronjob(ctx context.Context, cronjob entity.Cronjob) (entity.Cronjob, error) {
	cronjob, err := resource.cronjob.InsertCronjob(ctx, cronjob)
	if err != nil {
		logger.Trace(ctx, cronjob, err, "resource.cronjob.InsertCronjob() error - AddCronjob")
		return cronjob, err
	}

	return cronjob, nil
}

func (resource *Resource) UpdateCronjob(ctx context.Context, ID int64, cronjob entity.Cronjob) (entity.Cronjob, error) {
	newCronjob, err := resource.cronjob.UpdateCronjob(ctx, cronjob)
	if err != nil {
		logger.Trace(ctx, cronjob, err, "resource.cronjob.UpdateCronjob() error - UpdateCronjob")
		return newCronjob, err
	}

	return newCronjob, nil
}

func (resource *Resource) DeleteCronjob(ctx context.Context, ID int64, UserId int64) error {

	err := resource.cronjob.DeleteCronjob(ctx, ID, UserId)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "resource.cronjob.DeleteCronjob() error - DeleteCronjob")
		return err
	}

	return nil
}

func (resource *Resource) GetAllActiveCronjob(ctx context.Context) ([]entity.Cronjob, error) {
	cronjobs, err := resource.cronjob.GetAllActiveCronjob(ctx)
	if err != nil {
		logger.Trace(ctx, nil, err, "resource.cronjob.GetAllActiveCronjob() error - GetAllActiveCronjob")
		return cronjobs, err
	}

	return cronjobs, nil
}
