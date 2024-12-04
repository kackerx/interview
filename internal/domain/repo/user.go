package repo

import (
	"context"

	"github.com/kackerx/interview/internal/domain/do"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *do.User) (uint, error)
	FindUserByName(ctx context.Context, userName string) (*do.User, bool, error)
	FindUserByID(ctx context.Context, userID uint) (*do.User, bool, error)
}
