package data

import (
	"context"

	"github.com/kackerx/interview/common/code"
	"github.com/kackerx/interview/internal/data/convertor"
	"github.com/kackerx/interview/internal/data/model"
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

func (t *DocumentDao) FindDocumentByTaskID(ctx context.Context, taskID uint) (*do.Document, error) {
	docModel := new(model.Document)
	if err := t.DB(ctx).Where("task_id = ?", taskID).First(docModel).Error; err != nil {
		return nil, code.ErrDBUnknow.WithCause(err)
	}

	return convertor.DocumentModel2Do(docModel), nil
}

func NewDocumentRepo(data *Data) repo.DocumentRepo {
	return &DocumentDao{Data: data}
}

func (t *DocumentDao) UpdateContent(ctx context.Context, i18nContent string, taskID uint) error {
	return t.DB(ctx).Model(&model.Document{}).Where("task_id = ?", taskID).Update("i18n_content", i18nContent).Error
}
