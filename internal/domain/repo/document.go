package repo

import (
	"context"

	"github.com/kackerx/interview/internal/domain/do"
)

type DocumentRepo interface {
	CreateDocument(ctx context.Context, task *do.Document) (uint, error)
	FindDocumentByTaskID(ctx context.Context, taskID uint) (*do.Document, error)
	UpdateContent(ctx context.Context, i18nContent string, taskID uint) error
}
