package data

import (
	"context"

	"github.com/kackerx/interview/common/code"
	"github.com/kackerx/interview/internal/data/convertor"
	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/repo"
)

type TaskDao struct {
	*Data
}

func (t *TaskDao) CreateTask(ctx context.Context, task *do.Task) (uint, error) {
	taskMoel := convertor.TaskDo2Model(task)
	if err := t.master.Create(taskMoel).Error; err != nil {
		return 0, code.ErrDBUnknow.WithCause(err)
	}

	return taskMoel.ID, nil
}

func NewTaskRepo(data *Data) repo.TaskRepo {
	return &TaskDao{Data: data}
}
