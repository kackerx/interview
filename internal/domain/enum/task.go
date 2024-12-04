package enum

type TaskStatus string

const (
	TaskStatusCreated   TaskStatus = "created"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusFinished  TaskStatus = "finished"
	TaskStatusCompleted TaskStatus = "completed"
)
