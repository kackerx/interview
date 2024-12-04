package service

import (
	"context"

	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/repo"
)

type UserDomainService struct {
	*Service
	userRepo repo.UserRepo
}

func NewUserDomainService(service *Service, userRepo repo.UserRepo) *UserDomainService {
	return &UserDomainService{Service: service, userRepo: userRepo}
}

func (u *UserDomainService) Register(ctx context.Context, user *do.User) (uid uint, err error) {
	u.userRepo.CreateUser(ctx)
	return
}
