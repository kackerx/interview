package repo

import (
	"context"

	"github.com/kackerx/interview/internal/domain/do"
)

type UserRepo interface {
	CreateUser(ctx context.Context) error
	FindUserByName(ctx context.Context, userName string) (*do.User, bool, error)
	FindUserByID(ctx context.Context, userID string) (*do.User, bool, error)
}
