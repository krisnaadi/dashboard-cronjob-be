package log

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/repository/postgre/log"
)

type ResourceProvider interface {
	GetLogs(ctx context.Context) ([]entity.Log, error)
	GetLogByID(ctx context.Context, ID int64) (entity.Log, error)
	AddLog(ctx context.Context, log entity.Log) (entity.Log, error)
	UpdateLog(ctx context.Context, ID int64, log entity.Log) (entity.Log, error)
	DeleteLog(ctx context.Context, ID int64) error
}

type Resource struct {
	log log.RepositoryProvider
}

func New(log log.RepositoryProvider) ResourceProvider {
	resource := Resource{}

	resource.log = log

	return &resource
}
