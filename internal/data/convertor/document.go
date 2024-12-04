package convertor

import (
	"github.com/kackerx/interview/internal/data/model"
	"github.com/kackerx/interview/internal/domain/do"
)

func DocumentDo2Model(doc *do.Document) *model.Document {
	return &model.Document{
		TaskID:    doc.TaskID,
		Content:   doc.Content,
		CreatedBy: doc.CreatedBy,
	}
}