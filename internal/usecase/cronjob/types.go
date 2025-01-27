package cronjob

type CronjobRequest struct {
	Name     string `json:"name" validate:"required"`
	Schedule string `json:"schedule" validate:"required"`
	Task     string `json:"task" validate:"required"`
	Status   bool   `json:"status" validate:"required"`
}
