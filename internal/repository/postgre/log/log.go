package log

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
)

func (repository *Repository) GetLogs(ctx context.Context, JobId int64) ([]entity.Log, error) {
	var logs []entity.Log
	err := repository.db.Where("job_id = ?", JobId).Order("id asc").Find(&logs).Error

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
