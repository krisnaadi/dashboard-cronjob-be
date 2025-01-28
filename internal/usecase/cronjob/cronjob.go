package cronjob

import (
	"context"
	"errors"
	"time"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/execcommand"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/scheduler"
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

func (useCase *UseCase) RunAllCronjob(ctx context.Context) error {
	cronjobs, err := useCase.cronjob.GetAllActiveCronjob(ctx)
	if err != nil {
		logger.Trace(ctx, nil, err, "useCase.cronjob.GetAllActiveCronjob() error - RunAllCronjob")
		return err
	}
	scheduler := scheduler.New()
	for _, cronjob := range cronjobs {
		scheduler.AddJob(cronjob.Schedule, runTask(ctx, cronjob, useCase))
	}

	return nil
}

func runTask(ctx context.Context, cronjob entity.Cronjob, useCase *UseCase) error {
	start := time.Now()
	log := entity.Log{
		JobId:         cronjob.ID,
		ExecutionTime: start,
	}
	err := execcommand.Shellout(cronjob.Task)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{cronjob.ID}, err, "execcommand.Shellout() error - RunCronjob")
		log.Status = false
		log.ErrorMessage = err.Error()
		log.Duration = time.Since(start).Milliseconds()
		_, err = useCase.log.AddLog(ctx, log)
		if err != nil {
			logger.Trace(ctx, log, err, "useCase.log.AddLog() error - RunCronjob")
			return err
		}
		return err
	}
	log.Status = true
	log.Duration = time.Since(start).Milliseconds()
	_, err = useCase.log.AddLog(ctx, log)
	if err != nil {
		logger.Trace(ctx, log, err, "useCase.log.AddLog() error - RunCronjob")
		return err
	}
	return nil
}
