package cronjob

import "time"

type CronjobResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
	Task     string `json:"task"`
	Status   bool   `json:"status"`
}

type LogResponse struct {
	ID            int64     `json:"id"`
	ExecutionTime time.Time `json:"execution_time"`
	Status        bool      `json:"status"`
	Duration      int64     `json:"duration"`
	ErrorMessage  string    `json:"error_message"`
}
