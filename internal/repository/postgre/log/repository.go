package log

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"

	"gorm.io/gorm"
)

type RepositoryProvider interface {
	GetLogs(ctx context.Context, JobId int64) ([]entity.Log, error)
	InsertLog(ctx context.Context, log entity.Log) (entity.Log, error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryProvider {
	return &Repository{
		db: db,
	}
}
