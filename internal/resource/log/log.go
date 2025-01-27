package log

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
)

func (resource *Resource) GetLogByID(ctx context.Context, ID int64) (entity.Log, error) {
	log, err := resource.log.GetLogByID(ctx, ID)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "resource.log.GetLogByID() error - GetLogByID")
		return log, err
	}

	return log, nil
}

func (resource *Resource) GetLogs(ctx context.Context) ([]entity.Log, error) {
	logs, err := resource.log.GetLogs(ctx)
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

func (resource *Resource) UpdateLog(ctx context.Context, ID int64, log entity.Log) (entity.Log, error) {
	newLog, err := resource.log.UpdateLog(ctx, log)
	if err != nil {
		logger.Trace(ctx, log, err, "resource.log.UpdateLog() error - UpdateLog")
		return newLog, err
	}

	return newLog, nil
}

func (resource *Resource) DeleteLog(ctx context.Context, ID int64) error {

	err := resource.log.DeleteLog(ctx, ID)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "resource.log.DeleteLog() error - DeleteLog")
		return err
	}

	return nil
}
