package log

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
)

func (resource *Resource) GetLogs(ctx context.Context, JobId int64) ([]entity.Log, error) {
	logs, err := resource.log.GetLogs(ctx, JobId)
	if err != nil {
		logger.Trace(ctx, nil, err, "resource.log.GetLogs() error - GetLogs")
		return logs, err
	}

	return logs, nil
}

func (resource *Resource) AddLog(ctx context.Context, log entity.Log) (entity.Log, error) {
	log, err := resource.log.InsertLog(ctx, log)
	if err != nil {
		logger.Trace(ctx, log, err, "resource.log.InsertLog() error - AddLog")
		return log, err
	}

	return log, nil
}
