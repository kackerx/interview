package repo

import (
	"context"

	"github.com/kackerx/interview/internal/domain/do"
)

type TaskRepo interface {
	CreateTask(ctx context.Context, task *do.Task) (uint, error)
	FindTaskByTaskID(ctx context.Context, taskID uint) (*do.Task, error)
}
