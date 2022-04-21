package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// DocumentTypeAppInterface is a mock of application.DocumentTypeAppInterface.
type DocumentTypeAppInterface struct {
	SaveDocumentTypeFn   func(*entity.DocumentType) (*entity.DocumentType, map[string]string, error)
	UpdateDocumentTypeFn func(string, *entity.DocumentType) (*entity.DocumentType, map[string]string, error)
	DeleteDocumentTypeFn func(UUID string) error
	GetDocumentTypesFn   func(params *repository.Parameters) ([]*entity.DocumentType, *repository.Meta, error)
	GetDocumentTypeFn    func(UUID string) (*entity.DocumentType, error)
}

// SaveDocumentType calls the SaveDocumentTypeFn.
func (u *DocumentTypeAppInterface) SaveDocumentType(
	documentType *entity.DocumentType,
) (*entity.DocumentType, map[string]string, error) {
	return u.SaveDocumentTypeFn(documentType)
}

// UpdateDocumentType calls the UpdateDocumentTypeFn.
func (u *DocumentTypeAppInterface) UpdateDocumentType(
	uuid string,
	documentType *entity.DocumentType,
) (*entity.DocumentType, map[string]string, error) {
	return u.UpdateDocumentTypeFn(uuid, documentType)
}

// DeleteDocumentType calls the DeleteDocumentTypeFn.
func (u *DocumentTypeAppInterface) DeleteDocumentType(uuid string) error {
	return u.DeleteDocumentTypeFn(uuid)
}

// GetDocumentTypes calls the GetDocumentTypesFn.
func (u *DocumentTypeAppInterface) GetDocumentTypes(
	params *repository.Parameters,
) ([]*entity.DocumentType, *repository.Meta, error) {
	return u.GetDocumentTypesFn(params)
}

// GetDocumentType calls the GetDocumentTypeFn.
func (u *DocumentTypeAppInterface) GetDocumentType(uuid string) (*entity.DocumentType, error) {
	return u.GetDocumentTypeFn(uuid)
}
