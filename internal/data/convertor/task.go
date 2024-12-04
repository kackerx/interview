package convertor

import (
	"github.com/kackerx/interview/internal/data/model"
	"github.com/kackerx/interview/internal/domain/do"
)

func TaskDo2Model(task *do.Task) *model.Task {
	return &model.Task{
		Status:    string(task.Status),
		CreatedBy: task.CreatedBy,
	}
}
