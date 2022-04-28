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

// OrderStatusType represent schema of table order_status_type.
type OrderStatusType struct {
	UUID      string         `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid,omitempty"`
	Type      string         `gorm:"size:100;not null;"                        json:"type,omitempty"       form:"type"`
	CreatedAt time.Time      `                                                 json:"created_at,omitempty"`
	UpdatedAt time.Time      `                                                 json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `                                                 json:"deleted_at,omitempty"`
}

// OrderStatusTypeFaker represent content when generate fake data of document_type.
type OrderStatusTypeFaker struct {
	UUID string `faker:"uuid_hyphenated"`
	Type string `faker:"type"`
}

// OrderStatusTypes represent multiple OrderStatusType.
type OrderStatusTypes []*OrderStatusType

// DetailOrderStatusType represent format of detail OrderStatusType.
type DetailOrderStatusType struct {
	OrderStatusTypeFieldsForDetail
}

// DetailOrderStatusTypeList represent format of DetailOrderStatusType for OrderStatusType list.
type DetailOrderStatusTypeList struct {
	OrderStatusTypeFieldsForDetail
	OrderStatusTypeFieldsForList
}

// OrderStatusTypeFieldsForDetail represent fields of detail OrderStatusType.
type OrderStatusTypeFieldsForDetail struct {
	UUID string `json:"uuid"`
	Type string `json:"Type"`
}

// OrderStatusTypeFieldsForList represent fields of detail OrderStatusType for OrderStatusType list.
type OrderStatusTypeFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *OrderStatusType) TableName() string {
	return "order_status_types"
}

// FilterableFields return fields.
func (u *OrderStatusType) FilterableFields() []interface{} {
	return []interface{}{"uuid", "type"}
}

// Prepare will prepare submitted data of document_type.
func (u *OrderStatusType) Prepare() {
	u.Type = html.EscapeString(strings.TrimSpace(u.Type))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *OrderStatusType) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailOrderStatusTypes will return formatted document_type detail of multiple document_type.
func (orderStatusType OrderStatusTypes) DetailOrderStatusTypes() []interface{} {
	result := make([]interface{}, len(orderStatusType))
	for index, document_type := range orderStatusType {
		result[index] = document_type.DetailOrderStatusTypeList()
	}
	return result
}

// DetailOrderStatusType will return formatted document_type detail of document_type.
func (u *OrderStatusType) DetailOrderStatusType() interface{} {
	return &DetailOrderStatusType{
		OrderStatusTypeFieldsForDetail: OrderStatusTypeFieldsForDetail{
			UUID: u.UUID,
			Type: u.Type,
		},
	}
}

// DetailOrderStatusTypeList will return formatted document_type detail of document_type for document_type list.
func (u *OrderStatusType) DetailOrderStatusTypeList() interface{} {
	return &DetailOrderStatusTypeList{
		OrderStatusTypeFieldsForDetail: OrderStatusTypeFieldsForDetail{
			UUID: u.UUID,
			Type: u.Type,
		},
		OrderStatusTypeFieldsForList: OrderStatusTypeFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSaveOrderStatusType will validate create a new document_type request.
func (u *OrderStatusType) ValidateSaveOrderStatusType() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set(
			"type",
			u.Type,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		)
	return validation.Validate()
}

// ValidateUpdateOrderStatusType will validate update a new document_type request.
func (u *OrderStatusType) ValidateUpdateOrderStatusType() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set(
			"type",
			u.Type,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		)
	return validation.Validate()
}
