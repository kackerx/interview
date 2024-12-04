package appservice

import (
	"context"

	"github.com/kackerx/interview/api/v1/reply"
	"github.com/kackerx/interview/api/v1/request"
	"github.com/kackerx/interview/internal/data"
	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/enum"
	"github.com/kackerx/interview/internal/domain/service"
)

type TaskAppService struct {
	*AppService
	taskDomainSvc *service.TaskDomainService
	tm            data.Transaction
}

func NewtaskAppService(
	appService *AppService,
	taskDomainService *service.TaskDomainService,
	tm data.Transaction,
) *TaskAppService {
	return &TaskAppService{AppService: appService, taskDomainSvc: taskDomainService, tm: tm}
}

func (u *TaskAppService) CreateTask(ctx context.Context, req *request.CreateTaskReq, userName string) (resp *reply.CreateTaskResp, err error) {
	var taskID uint
	err = u.tm.Transaction(ctx, func(ctx context.Context) error {
		taskID, err = u.taskDomainSvc.Create(ctx, &do.Task{
			Status:    enum.TaskStatusCreated,
			CreatedBy: userName,
		})
		return err
	})

	return &reply.CreateTaskResp{TaskID: taskID}, err
}
