package app

import (
	"github.com/krisnaadi/dashboard-cronjob-be/internal/resource/cronjob"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/resource/log"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/resource/user"
)

type Resources struct {
	user    user.ResourceProvider
	cronjob cronjob.ResourceProvider
	log     log.ResourceProvider
}

// NewResource initializes resource layer.
func NewResource(repositories *Repositories) *Resources {
	return &Resources{
		user:    user.New(repositories.user),
		cronjob: cronjob.New(repositories.cronjob),
		log:     log.New(repositories.log),
	}
}
