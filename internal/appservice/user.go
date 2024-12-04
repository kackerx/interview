package appservice

import (
	"context"
	"fmt"

	"github.com/kackerx/interview/api/v1/reply"
	"github.com/kackerx/interview/api/v1/request"
	"github.com/kackerx/interview/internal/domain/do"
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
	uid, err := u.userDomainSvc.Register(ctx, &do.User{})
	if err != nil {
		return nil, err
	}

	fmt.Println(uid)

	return nil, nil
}
