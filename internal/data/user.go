package data

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/kackerx/interview/common/code"
	"github.com/kackerx/interview/internal/data/convertor"
	"github.com/kackerx/interview/internal/data/model"
	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/repo"
)

type UserDao struct {
	*Data
}

func NewUserRepo(data *Data) repo.UserRepo {
	return &UserDao{Data: data}
}

func (u *UserDao) CreateUser(ctx context.Context, user *do.User) (uint, error) {
	userModel := convertor.UserDo2Model(user)
	if err := u.Data.DB(ctx).Create(userModel).Error; err != nil {
		return 0, code.ErrDBUnknow.WithCause(err)
	}

	return userModel.ID, nil
}

func (u *UserDao) FindUserByName(ctx context.Context, userName string) (*do.User, bool, error) {
	return u.getUserByCond(ctx, map[string]any{"user_name": userName})
}

func (u *UserDao) FindUserByID(ctx context.Context, userID uint) (*do.User, bool, error) {
	return u.getUserByCond(ctx, map[string]any{"user_id": userID})
}

func (u *UserDao) getUserByCond(ctx context.Context, conds map[string]any) (*do.User, bool, error) {
	user := new(model.User)
	if err := u.DB(ctx).Where(conds).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, code.ErrDBUnknow.WithCause(err)
	}

	return convertor.UserModel2Do(user), true, nil
}
