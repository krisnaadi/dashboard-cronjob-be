package app

import (
	"github.com/krisnaadi/dashboard-cronjob-be/internal/repository/postgre/cronjob"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/repository/postgre/log"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/repository/postgre/user"
	"gorm.io/gorm"
)

type Repositories struct {
	user    user.RepositoryProvider
	cronjob cronjob.RepositoryProvider
	log     log.RepositoryProvider
	db      *gorm.DB
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		user:    user.New(db),
		cronjob: cronjob.New(db),
		log:     log.New(db),
		db:      db,
	}
}
