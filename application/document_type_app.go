package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type documentTypeApp struct {
	tr repository.DocumentTypeRepository
}

// documentTypeApp implement the DocumentTypeAppInterface.
var _ DocumentTypeAppInterface = &documentTypeApp{}

// DocumentTypeAppInterface is an interface.
type DocumentTypeAppInterface interface {
	SaveDocumentType(*entity.DocumentType) (*entity.DocumentType, map[string]string, error)
	UpdateDocumentType(UUID string, documentType *entity.DocumentType) (*entity.DocumentType, map[string]string, error)
	DeleteDocumentType(UUID string) error
	GetDocumentTypes(p *repository.Parameters) ([]*entity.DocumentType, *repository.Meta, error)
	GetDocumentType(UUID string) (*entity.DocumentType, error)
}

func (t documentTypeApp) SaveDocumentType(
	documentType *entity.DocumentType,
) (*entity.DocumentType, map[string]string, error) {
	return t.tr.SaveDocumentType(documentType)
}

func (t documentTypeApp) UpdateDocumentType(
	UUID string,
	documentType *entity.DocumentType,
) (*entity.DocumentType, map[string]string, error) {
	return t.tr.UpdateDocumentType(UUID, documentType)
}

func (t documentTypeApp) DeleteDocumentType(UUID string) error {
	return t.tr.DeleteDocumentType(UUID)
}

func (t documentTypeApp) GetDocumentTypes(p *repository.Parameters) ([]*entity.DocumentType, *repository.Meta, error) {
	return t.tr.GetDocumentTypes(p)
}

func (t documentTypeApp) GetDocumentType(UUID string) (*entity.DocumentType, error) {
	return t.tr.GetDocumentType(UUID)
}
