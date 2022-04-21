package repository

import (
	"cargo-rest-api/domain/entity"
)

// DocumentTypeRepository is an interface.
type DocumentTypeRepository interface {
	SaveDocumentType(tour *entity.DocumentType) (*entity.DocumentType, map[string]string, error)
	UpdateDocumentType(UUID string, tour *entity.DocumentType) (*entity.DocumentType, map[string]string, error)
	DeleteDocumentType(UUID string) error
	GetDocumentType(UUID string) (*entity.DocumentType, error)
	GetDocumentTypes(parameters *Parameters) ([]*entity.DocumentType, *Meta, error)
}
