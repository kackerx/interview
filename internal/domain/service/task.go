package service

import (
	"context"

	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/repo"
)

type TaskDomainService struct {
	*Service
	taskRepo repo.TaskRepo
}

func NewTaskDomainService(service *Service, userRepo repo.TaskRepo) *TaskDomainService {
	return &TaskDomainService{Service: service, taskRepo: userRepo}
}

func (u *TaskDomainService) Create(ctx context.Context, task *do.Task) (uid uint, err error) {
	return u.taskRepo.CreateTask(ctx, task)
}
func (u *TaskDomainService) GetTaskByTaskID(ctx context.Context, taskID uint) (*do.Task, error) {
	return u.taskRepo.FindTaskByTaskID(ctx, taskID)
}
