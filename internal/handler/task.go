package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/kackerx/interview/api/v1/request"
	"github.com/kackerx/interview/common/code"
	"github.com/kackerx/interview/common/middleware"
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
	claims, _ := c.Get("claims")
	user := claims.(*middleware.MyCustomClaims)

	req := &request.CreateTaskReq{}
	if err := c.ShouldBind(req); err != nil {
		resp.HandleErr(c, code.ErrParams.WithCause(err).WithArgs(err))
		return
	}
	ret, err := u.taskAppSvc.CreateTask(c, req, user.UserName)
	if err != nil {
		resp.HandleErr(c, err)
		return
	}

	resp.HandleSuccess(c, ret)
}

func (u *TaskHandler) Translate(c *gin.Context) {
	taskID := c.Param("task_id")
	id, _ := strconv.Atoi(taskID)
	if err := u.taskAppSvc.TransTask(c, uint(id)); err != nil {
		resp.HandleErr(c, err)
	} else {
		resp.HandleSuccess(c)
	}
}

func (u *TaskHandler) DetailTask(c *gin.Context) {
	taskID := c.Param("task_id")
	id, _ := strconv.Atoi(taskID)

	task, err := u.taskAppSvc.DetailTask(c, uint(id))
	if err != nil {
		resp.HandleErr(c, err)
		return
	}

	resp.HandleSuccess(c, task)
}

func (u *TaskHandler) DownTaskFile(c *gin.Context) {
	taskID := c.Param("task_id")
	id, _ := strconv.Atoi(taskID)

	fileName, err := u.taskAppSvc.DownTaskFile(c, uint(id))
	if err != nil {
		resp.HandleErr(c, err)
		return
	}

	c.File(fileName)
}
