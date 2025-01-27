package entity

import (
	"time"
)

type Log struct {
	ID            int64
	JobId         int64
	ExecutionTime time.Time
	Status        int64 // 1=success 0=fail
	Duration      int64
	ErrorMessage  string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}
