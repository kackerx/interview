package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/kackerx/interview/api/v1/request"
	"github.com/kackerx/interview/common/resp"
	"github.com/kackerx/interview/internal/appservice"
)

type TaskHandler struct {
	*Handler
	taskAppSvc *appservice.TaskAppService
}

func NewtaskHandler(handler *Handler, taskAppSvc *appservice.TaskAppService) *TaskHandler {
	return &TaskHandler{Handler: handler, taskAppSvc: taskAppSvc}
}

func (u *TaskHandler) CreateTask(c *gin.Context) {
	ret, err := u.taskAppSvc.CreateTask(c, &request.CreateTaskReq{}, c.GetString("user_name"))
	if err != nil {
		resp.HandleErr(c, err)
		return
	}

	resp.HandleSuccess(c, ret)
}
