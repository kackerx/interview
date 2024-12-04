package data

import (
	"context"

	"github.com/kackerx/interview/common/code"
	"github.com/kackerx/interview/internal/data/convertor"
	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/repo"
)

type DocumentDao struct {
	*Data
}

func (t *DocumentDao) CreateDocument(ctx context.Context, document *do.Document) (uint, error) {
	documentMoel := convertor.DocumentDo2Model(document)
	if err := t.master.Create(documentMoel).Error; err != nil {
		return 0, code.ErrDBUnknow.WithCause(err)
	}

	return documentMoel.ID, nil
}

func NewDocumentRepo(data *Data) repo.DocumentRepo {
	return &DocumentDao{Data: data}
}
