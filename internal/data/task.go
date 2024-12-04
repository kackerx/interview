package data

import (
	"context"

	"github.com/kackerx/interview/common/code"
	"github.com/kackerx/interview/internal/data/convertor"
	"github.com/kackerx/interview/internal/data/model"
	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/enum"
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

func (t *TaskDao) FindTaskByTaskID(ctx context.Context, taskID uint) (*do.Task, error) {
	taskModel := &model.Task{}
	if err := t.DB(ctx).Where("id = ?", taskID).First(taskModel).Error; err != nil {
		return nil, code.ErrDBUnknow.WithCause(err)
	}

	return &do.Task{
		ID:        taskID,
		Status:    enum.TaskStatus(taskModel.Status),
		CreatedBy: taskModel.CreatedBy,
		CreatedAt: taskModel.CreatedAt,
		UpdatedAt: taskModel.UpdatedAt,
	}, nil
}
