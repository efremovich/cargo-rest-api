package entity

import (
	"cargo-rest-api/pkg/response"
	"cargo-rest-api/pkg/validator"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

// DocumentType represent schema of table docement_type.
type DocumentType struct {
	UUID      string         `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid,omitempty"`
	Type      string         `gorm:"size:100;not null;"                        json:"type,omitempty"       form:"type"`
	CreatedAt time.Time      `                                                 json:"created_at,omitempty"`
	UpdatedAt time.Time      `                                                 json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `                                                 json:"deleted_at,omitempty"`
}

// DocumentTypeFaker represent content when generate fake data of document_type.
type DocumentTypeFaker struct {
	UUID string `faker:"uuid_hyphenated"`
	Type string `faker:"type"`
}

// DocumentTypes represent multiple DocumentType.
type DocumentTypes []*DocumentType

// DetailDocumentType represent format of detail DocumentType.
type DetailDocumentType struct {
	DocumentTypeFieldsForDetail
}

// DetailDocumentTypeList represent format of DetailDocumentType for DocumentType list.
type DetailDocumentTypeList struct {
	DocumentTypeFieldsForDetail
	DocumentTypeFieldsForList
}

// DocumentTypeFieldsForDetail represent fields of detail DocumentType.
type DocumentTypeFieldsForDetail struct {
	UUID string `json:"uuid"`
	Type string `json:"Type"`
}

// DocumentTypeFieldsForList represent fields of detail DocumentType for DocumentType list.
type DocumentTypeFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *DocumentType) TableName() string {
	return "document_types"
}

// FilterableFields return fields.
func (u *DocumentType) FilterableFields() []interface{} {
	return []interface{}{"uuid", "type"}
}

// Prepare will prepare submitted data of document_type.
func (u *DocumentType) Prepare() {
	u.Type = html.EscapeString(strings.TrimSpace(u.Type))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *DocumentType) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailDocumentTypes will return formatted document_type detail of multiple document_type.
func (documentType DocumentTypes) DetailDocumentTypes() []interface{} {
	result := make([]interface{}, len(documentType))
	for index, document_type := range documentType {
		result[index] = document_type.DetailDocumentTypeList()
	}
	return result
}

// DetailDocumentType will return formatted document_type detail of document_type.
func (u *DocumentType) DetailDocumentType() interface{} {
	return &DetailDocumentType{
		DocumentTypeFieldsForDetail: DocumentTypeFieldsForDetail{
			UUID: u.UUID,
			Type: u.Type,
		},
	}
}

// DetailDocumentTypeList will return formatted document_type detail of document_type for document_type list.
func (u *DocumentType) DetailDocumentTypeList() interface{} {
	return &DetailDocumentTypeList{
		DocumentTypeFieldsForDetail: DocumentTypeFieldsForDetail{
			UUID: u.UUID,
			Type: u.Type,
		},
		DocumentTypeFieldsForList: DocumentTypeFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSaveDocumentType will validate create a new document_type request.
func (u *DocumentType) ValidateSaveDocumentType() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set(
			"type",
			u.Type,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		)
	return validation.Validate()
}

// ValidateUpdateDocumentType will validate update a new document_type request.
func (u *DocumentType) ValidateUpdateDocumentType() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set(
			"type",
			u.Type,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		)
	return validation.Validate()
}
