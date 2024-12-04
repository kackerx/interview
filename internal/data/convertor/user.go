package convertor

import (
	"github.com/kackerx/interview/internal/data/model"
	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/enum"
)

func UserDo2Model(user *do.User) *model.User {
	return &model.User{
		UserName: user.UserName,
		Password: user.Password,
		NickName: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Status:   enum.GetUserStatusNum(user.Status),
		Gender:   enum.GetUserGenderNum(user.Gender),
	}
}

func UserModel2Do(user *model.User) *do.User {
	return &do.User{
		UserID:    user.ID,
		UserName:  user.UserName,
		Password:  user.Password,
		Nickname:  user.NickName,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Status:    enum.GetUserStatus(user.Status),
		Gender:    enum.GetUserGender(user.Gender),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
