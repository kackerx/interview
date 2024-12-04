package appservice

import (
	"context"
	"time"

	"github.com/kackerx/interview/api/v1/reply"
	"github.com/kackerx/interview/api/v1/request"
	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/enum"
	"github.com/kackerx/interview/internal/domain/service"
)

type UserAppService struct {
	*AppService
	userDomainSvc *service.UserDomainService
}

func NewUserAppService(appService *AppService, userDomainService *service.UserDomainService) *UserAppService {
	return &UserAppService{AppService: appService, userDomainSvc: userDomainService}
}

func (u *UserAppService) RegisterUser(ctx context.Context, req *request.RegisterReq) (*reply.RegisterResp, error) {
	uid, err := u.userDomainSvc.Register(ctx, &do.User{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
		Status:   enum.UserStatusNormal,
	})
	if err != nil {
		return nil, err
	}

	return &reply.RegisterResp{UserID: uid}, nil
}

func (u *UserAppService) LoginUser(ctx context.Context, req *request.LoginReq) (*reply.LoginResp, error) {
	token, err := u.userDomainSvc.Login(ctx, req.UserName, req.Password)
	if err != nil {
		return nil, err
	}

	return &reply.LoginResp{
		Token: token,
		// Duration:  ,
		CreatedAt: time.Now().Format(time.DateTime),
	}, nil
}
