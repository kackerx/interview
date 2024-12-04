package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/kackerx/interview/api/v1/request"
	"github.com/kackerx/interview/internal/appservice"
)

type UserHandler struct {
	*Handler
	userAppSvc *appservice.UserAppService
}

func NewUserHandler(handler *Handler, userAppSvc *appservice.UserAppService) *UserHandler {
	return &UserHandler{Handler: handler, userAppSvc: userAppSvc}
}

func (u *UserHandler) Register(c *gin.Context) {
	u.userAppSvc.RegisterUser(c, &request.RegisterReq{})
}
