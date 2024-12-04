package appservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/kackerx/interview/api/v1/reply"
	"github.com/kackerx/interview/api/v1/request"
	"github.com/kackerx/interview/internal/data"
	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/enum"
	"github.com/kackerx/interview/internal/domain/service"
	"github.com/kackerx/interview/library/gpt"
)

type TaskAppService struct {
	*AppService
	taskDomainSvc *service.TaskDomainService
	docDoaminSvc  *service.DocumentDomainService
	tm            data.Transaction
}

func NewtaskAppService(
	appService *AppService,
	taskDomainService *service.TaskDomainService,
	domainService *service.DocumentDomainService,
	tm data.Transaction,
) *TaskAppService {
	return &TaskAppService{
		AppService:    appService,
		taskDomainSvc: taskDomainService,
		tm:            tm,
		docDoaminSvc:  domainService,
	}
}

func (u *TaskAppService) CreateTask(ctx context.Context, req *request.CreateTaskReq, userName string) (resp *reply.CreateTaskResp, err error) {
	var taskID uint
	err = u.tm.Transaction(ctx, func(ctx context.Context) error {
		taskID, err = u.taskDomainSvc.Create(ctx, &do.Task{
			Status:    enum.TaskStatusCreated,
			CreatedBy: userName,
		})

		_, err = u.docDoaminSvc.Create(ctx, &do.Document{
			TaskID:    taskID,
			Content:   req.Content,
			CreatedBy: userName,
		})
		return err
	})

	return &reply.CreateTaskResp{TaskID: taskID}, err
}

func (u *TaskAppService) TransTask(ctx context.Context, taskID uint) (err error) {
	doc, err := u.docDoaminSvc.GetDocumentByTaskID(ctx, taskID)
	if err != nil {
		return err
	}

	i18Content, err := gpt.Trans(ctx, doc.Content)
	if err != nil {
		return err
	}

	return u.docDoaminSvc.Trans(ctx, i18Content, taskID)
}

func (u *TaskAppService) DetailTask(ctx context.Context, taskID uint) (*reply.DetailTaskResp, error) {
	task, err := u.taskDomainSvc.GetTaskByTaskID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	return &reply.DetailTaskResp{
		TaskID:    task.ID,
		Status:    string(task.Status),
		CreatedAt: task.CreatedAt.Format(time.DateTime),
		UpdatedAt: task.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (u *TaskAppService) DownTaskFile(ctx context.Context, taskID uint) (string, error) {
	task, err := u.taskDomainSvc.GetTaskByTaskID(ctx, taskID)
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(task)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("/Users/apple/GolandProjects/interview/output/%d.json", taskID)
	// 将数据写入文件
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return "", err
	}

	return fileName, err
}
