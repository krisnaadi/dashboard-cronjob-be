package log

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/repository/postgre/log"
)

type ResourceProvider interface {
	GetLogs(ctx context.Context, JobId int64) ([]entity.Log, error)
	AddLog(ctx context.Context, log entity.Log) (entity.Log, error)
}

type Resource struct {
	log log.RepositoryProvider
}

func New(log log.RepositoryProvider) ResourceProvider {
	resource := Resource{}

	resource.log = log

	return &resource
}
