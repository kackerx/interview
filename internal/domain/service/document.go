package service

import (
	"context"

	"github.com/kackerx/interview/internal/domain/do"
	"github.com/kackerx/interview/internal/domain/repo"
)

type DocumentDomainService struct {
	*Service
	documentRepo repo.DocumentRepo
}

func NewDocumentDomainService(service *Service, userRepo repo.DocumentRepo) *DocumentDomainService {
	return &DocumentDomainService{Service: service, documentRepo: userRepo}
}

func (u *DocumentDomainService) Create(ctx context.Context, document *do.Document) (id uint, err error) {
	return u.documentRepo.CreateDocument(ctx, document)
}
