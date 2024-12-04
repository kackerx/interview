package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/kackerx/interview/api/v1/request"
	"github.com/kackerx/interview/common/code"
	"github.com/kackerx/interview/common/resp"
	"github.com/kackerx/interview/internal/appservice"
	"github.com/kackerx/interview/pkg/validate"
)

type UserHandler struct {
	*Handler
	userAppSvc *appservice.UserAppService
}

func NewUserHandler(handler *Handler, userAppSvc *appservice.UserAppService) *UserHandler {
	return &UserHandler{Handler: handler, userAppSvc: userAppSvc}
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	req := new(request.RegisterReq)
	trans, _ := validate.GetLocalTrans("")
	if err := c.ShouldBind(req); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			resp.HandleErr(c, code.ErrParams.WithCause(err).WithArgs(validate.RemoveTopStruct(errs.Translate(trans))))
			return
		}
		resp.HandleErr(c, code.ErrParams.WithCause(err))
		return
	}

	if ret, err := u.userAppSvc.RegisterUser(c, req); err != nil {
		resp.HandleErr(c, err)
		return
	} else {
		resp.HandleSuccess(c, ret)
	}
}

func (u *UserHandler) LoginUser(c *gin.Context) {
	req := new(request.LoginReq)
	if err := c.ShouldBind(req); err != nil {
		resp.HandleErr(c, code.ErrParams.WithCause(err).WithArgs(err))
		return
	}

	// if err := c.ShouldBindHeader(req.Header); err != nil {
	// 	resp.HandleErr(c, code.ErrParams.WithCause(err).WithArgs(err))
	// 	return
	// }
	//
	token, err := u.userAppSvc.LoginUser(c, req)
	if err != nil {
		resp.HandleErr(c, err)
		return
	}

	resp.HandleSuccess(c, token)
}
