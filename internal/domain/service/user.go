package service

import (
	"context"
	"time"

	"github.com/kackerx/interview/common/code"
	"github.com/kackerx/interview/common/middleware"
	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/repo"
	"github.com/kackerx/interview/pkg/util"
)

type UserDomainService struct {
	*Service
	jwt      *middleware.JWT
	userRepo repo.UserRepo
}

func NewUserDomainService(service *Service, userRepo repo.UserRepo, jwt *middleware.JWT) *UserDomainService {
	return &UserDomainService{Service: service, userRepo: userRepo, jwt: jwt}
}

func (u *UserDomainService) Register(ctx context.Context, user *do.User) (uid uint, err error) {
	_, exist, err := u.userRepo.FindUserByName(ctx, user.UserName)
	if err != nil {
		return 0, err
	}

	if exist {
		return 0, code.ErrUserExist.WithArgs(user.UserName)
	}

	user.Password, err = util.EncryptPass(user.Password)
	if err != nil {
		return 0, code.Wrap("密码加密失败", err)
	}

	return u.userRepo.CreateUser(ctx, user)
}

func (u *UserDomainService) Login(ctx context.Context, userName string, password string) (string, error) {
	user, exist, err := u.userRepo.FindUserByName(ctx, userName)
	if err != nil {
		return "", err
	}

	if !exist {
		return "", code.ErrUserNotExist.WithArgs(userName)
	}

	if err = util.ComparePass(user.Password, password); err != nil {
		return "", code.ErrUserPassInvalid.WithArgs(userName)
	}

	token, err := u.jwt.GenToken(user.UserID, user.UserName, time.Now().Add(time.Hour*24*7))
	if err != nil {
		return "", code.ErrUserTokenGenFaild.WithCause(err).WithArgs(user.UserID)
	}

	return token, nil
}
