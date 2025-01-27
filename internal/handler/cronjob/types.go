package cronjob

type CronjobResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
	Task     string `json:"code"`
	Status   bool   `json:"status"`
}
