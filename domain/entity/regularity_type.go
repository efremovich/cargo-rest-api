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

// RegularityType represent schema of table sities.
type RegularityType struct {
	UUID      string         `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid,omitempty"`
	Type      string         `gorm:"size:100;not null;"                        json:"type,omitempty"       form:"type"`
	CreatedAt time.Time      `                                                 json:"created_at,omitempty"`
	UpdatedAt time.Time      `                                                 json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `                                                 json:"deleted_at,omitempty"`
}

// RegularityTypeFaker represent content when generate fake data of document_type.
type RegularityTypeFaker struct {
	UUID string `faker:"uuid_hyphenated"`
	Type string `faker:"type"`
}

// RegularityTypes represent multiple RegularityType.
type RegularityTypes []*RegularityType

// DetailRegularityType represent format of detail RegularityType.
type DetailRegularityType struct {
	RegularityTypeFieldsForDetail
}

// DetailRegularityTypeList represent format of DetailRegularityType for RegularityType list.
type DetailRegularityTypeList struct {
	RegularityTypeFieldsForDetail
	RegularityTypeFieldsForList
}

// RegularityTypeFieldsForDetail represent fields of detail RegularityType.
type RegularityTypeFieldsForDetail struct {
	UUID string `json:"uuid"`
	Type string `json:"Type"`
}

// RegularityTypeFieldsForList represent fields of detail RegularityType for RegularityType list.
type RegularityTypeFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *RegularityType) TableName() string {
	return "document_types"
}

// FilterableFields return fields.
func (u *RegularityType) FilterableFields() []interface{} {
	return []interface{}{"uuid", "type"}
}

// Prepare will prepare submitted data of document_type.
func (u *RegularityType) Prepare() {
	u.Type = html.EscapeString(strings.TrimSpace(u.Type))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *RegularityType) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailRegularityTypes will return formatted document_type detail of multiple document_type.
func (regularityType RegularityTypes) DetailRegularityTypes() []interface{} {
	result := make([]interface{}, len(regularityType))
	for index, document_type := range regularityType {
		result[index] = document_type.DetailRegularityTypeList()
	}
	return result
}

// DetailRegularityType will return formatted document_type detail of document_type.
func (u *RegularityType) DetailRegularityType() interface{} {
	return &DetailRegularityType{
		RegularityTypeFieldsForDetail: RegularityTypeFieldsForDetail{
			UUID: u.UUID,
			Type: u.Type,
		},
	}
}

// DetailRegularityTypeList will return formatted document_type detail of document_type for document_type list.
func (u *RegularityType) DetailRegularityTypeList() interface{} {
	return &DetailRegularityTypeList{
		RegularityTypeFieldsForDetail: RegularityTypeFieldsForDetail{
			UUID: u.UUID,
			Type: u.Type,
		},
		RegularityTypeFieldsForList: RegularityTypeFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSaveRegularityType will validate create a new document_type request.
func (u *RegularityType) ValidateSaveRegularityType() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set(
			"type",
			u.Type,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		)
	return validation.Validate()
}

// ValidateUpdateRegularityType will validate update a new document_type request.
func (u *RegularityType) ValidateUpdateRegularityType() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set(
			"type",
			u.Type,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		)
	return validation.Validate()
}
