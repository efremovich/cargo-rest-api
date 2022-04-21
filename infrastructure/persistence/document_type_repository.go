package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// DocumentTypeRepo is a struct to store db connection.
type DocumentTypeRepo struct {
	db *gorm.DB
}

// NewDocumentTypeRepository will initialize DocumentType repository.
func NewDocumentTypeRepository(db *gorm.DB) *DocumentTypeRepo {
	return &DocumentTypeRepo{db}
}

// DocumentTypeRepo implements the repository.documentTypeRepository interface.
var _ repository.DocumentTypeRepository = &DocumentTypeRepo{}

// SaveDocumentType will create a new documentType.
func (r DocumentTypeRepo) SaveDocumentType(
	DocumentType *entity.DocumentType,
) (*entity.DocumentType, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&DocumentType).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return DocumentType, nil, nil
}

func (r DocumentTypeRepo) UpdateDocumentType(
	uuid string,
	documentType *entity.DocumentType,
) (*entity.DocumentType, map[string]string, error) {
	errDesc := map[string]string{}
	documentTypeData := &entity.DocumentType{
		Type: documentType.Type,
	}

	err := r.db.First(&documentType, "uuid = ?", uuid).Updates(documentTypeData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextDocumentTypeInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextDocumentTypeNotFound
		}
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			errDesc["type"] = exception.ErrorTextUserEmailAlreadyTaken.Error()
			return nil, errDesc, exception.ErrorTextUnprocessableEntity
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return documentType, nil, nil
}

func (r DocumentTypeRepo) DeleteDocumentType(uuid string) error {
	var documentType entity.DocumentType
	err := r.db.Where("uuid = ?", uuid).Take(&documentType).Delete(&documentType).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextDocumentTypeNotFound
		}
		return err
	}
	return nil
}

func (r DocumentTypeRepo) GetDocumentType(uuid string) (*entity.DocumentType, error) {
	var documentType entity.DocumentType
	err := r.db.Where("uuid = ?", uuid).Take(&documentType).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextDocumentTypeNotFound
		}
	}
	return &documentType, nil
}

func (r DocumentTypeRepo) GetDocumentTypes(p *repository.Parameters) ([]*entity.DocumentType, *repository.Meta, error) {
	var total int64
	var documentTypes []*entity.DocumentType
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&documentTypes).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&documentTypes).Error
	if errTotal != nil {
		return nil, nil, errTotal
	}
	if errList != nil {
		return nil, nil, errList
	}
	if errors.Is(errList, gorm.ErrRecordNotFound) {
		return nil, nil, errList
	}
	meta := repository.NewMeta(p, total)
	return documentTypes, meta, nil
}
