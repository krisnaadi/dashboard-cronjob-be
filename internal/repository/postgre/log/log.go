package log

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
)

func (repository *Repository) GetLogByID(ctx context.Context, ID int64) (entity.Log, error) {
	var log entity.Log

	err := repository.db.Where("id = ?", ID).Find(&log).Error
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "repository.db.Where().Find() error - GetLogByID")
		return entity.Log{}, err
	}

	return log, nil
}

func (repository *Repository) GetLogs(ctx context.Context) ([]entity.Log, error) {
	var logs []entity.Log
	err := repository.db.Order("id asc").Find(&logs).Error

	if err != nil {
		logger.Trace(ctx, nil, err, "repository.db.Order().Find() error - GetLogs")
		return nil, err
	}

	return logs, nil
}

func (repository *Repository) InsertLog(ctx context.Context, log entity.Log) (entity.Log, error) {
	err := repository.db.Create(&log).Error
	if err != nil {
		logger.Trace(ctx, log, err, "repository.db.Create() error - InsertLog")
		return entity.Log{}, err
	}

	return log, nil
}

func (repository *Repository) UpdateLog(ctx context.Context, log entity.Log) (entity.Log, error) {
	err := repository.db.Save(&log).Error
	if err != nil {
		logger.Trace(ctx, log, err, "repository.db.Save() error - UpdateLog")
		return entity.Log{}, err
	}

	return log, nil
}

func (repository *Repository) DeleteLog(ctx context.Context, ID int64) error {
	err := repository.db.Delete(&entity.Log{}, ID).Error
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "repository.db.Delete() error - DeleteLog")
		return err
	}

	return nil
}
