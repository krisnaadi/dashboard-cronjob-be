package entity

import (
	"time"
)

type Cronjob struct {
	ID        int64
	Name      string
	Schedule  string
	Task      string
	Status    string
	UserId    int64
	CreatedAt time.Time
	UpdatedAt *time.Time
}
