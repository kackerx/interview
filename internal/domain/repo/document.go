package repo

import (
	"context"

	"github.com/kackerx/interview/internal/domain/do"
)

type DocumentRepo interface {
	CreateDocument(ctx context.Context, task *do.Document) (uint, error)
}
