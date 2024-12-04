package reply

type CreateTaskResp struct {
	TaskID uint `json:"task_id"`
}

type DetailTaskResp struct {
	TaskID    uint   `json:"task_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
