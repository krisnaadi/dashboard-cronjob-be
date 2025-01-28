package cronjob

import (
	"context"
	"errors"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
)

func (useCase *UseCase) ListCronjob(ctx context.Context, UserId int64) ([]entity.Cronjob, error) {

	categories, err := useCase.cronjob.GetCronjobs(ctx, UserId)
	if err != nil {
		logger.Trace(ctx, nil, err, "useCase.cronjob.GetCategories() error - ListCronjob")
		return categories, err
	}

	return categories, nil
}

func (useCase *UseCase) GetCronjob(ctx context.Context, ID int64, UserId int64) (entity.Cronjob, error) {

	cronjob, err := useCase.cronjob.GetCronjobByID(ctx, ID, UserId)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "useCase.cronjob.GetCronjobByID() error - GetCronjob")
		return cronjob, err
	}

	return cronjob, nil
}

func (useCase *UseCase) AddCronjob(ctx context.Context, input CronjobRequest, UserId int64) (entity.Cronjob, error) {

	cronjob := entity.Cronjob{
		Name:     input.Name,
		Task:     input.Task,
		Schedule: input.Schedule,
		Status:   input.Status,
		UserId:   UserId,
	}

	cronjob, err := useCase.cronjob.AddCronjob(ctx, cronjob)
	if err != nil {
		logger.Trace(ctx, cronjob, err, "useCase.cronjob.AddCronjob() error - AddCronjob")
		return cronjob, err
	}

	return cronjob, nil
}

func (useCase *UseCase) UpdateCronjob(ctx context.Context, ID int64, input CronjobRequest, UserId int64) (entity.Cronjob, error) {

	oldCronjob, err := useCase.cronjob.GetCronjobByID(ctx, ID, UserId)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "useCase.cronjob.GetCronjobByID() error - UpdateCronjob")
		return oldCronjob, err
	}

	if oldCronjob.ID == 0 {
		return oldCronjob, errors.New("data cronjob not found")
	}

	oldCronjob.Name = input.Name
	oldCronjob.Schedule = input.Schedule
	oldCronjob.Status = input.Status
	oldCronjob.Task = input.Task

	cronjob, err := useCase.cronjob.UpdateCronjob(ctx, ID, oldCronjob)
	if err != nil {
		logger.Trace(ctx, struct {
			ID      int64
			cronjob entity.Cronjob
		}{ID, oldCronjob}, err, "useCase.cronjob.UpdateCronjob() error - UpdateCronjob")
		return cronjob, err
	}

	return cronjob, nil
}

func (useCase *UseCase) DeleteCronjob(ctx context.Context, ID int64, UserId int64) error {

	cronjob, err := useCase.cronjob.GetCronjobByID(ctx, ID, UserId)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "useCase.cronjob.GetCronjobByID() error - DeleteCronjob")
		return err
	}

	if cronjob.ID == 0 {
		return errors.New("data cronjob not found")
	}

	err = useCase.cronjob.DeleteCronjob(ctx, cronjob.ID, UserId)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{cronjob.ID}, err, "useCase.cronjob.DeleteCronjob() error - DeleteCronjob")
		return err
	}

	return nil
}
