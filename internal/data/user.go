package data

import (
	"context"

	"github.com/kackerx/interview/common/log"
	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/repo"
)

type UserDao struct {
	*Data
}

func NewUserRepo(data *Data) repo.UserRepo {
	return &UserDao{Data: data}
}

func (u *UserDao) CreateUser(ctx context.Context) error {
	log.New(ctx).Info("create user", 1, 2)
	return nil
}

func (u *UserDao) FindUserByName(ctx context.Context, userName string) (*do.User, bool, error) {
	// TODO implement me
	panic("implement me")
}

func (u *UserDao) FindUserByID(ctx context.Context, userID string) (*do.User, bool, error) {
	// TODO implement me
	panic("implement me")
}
