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

// PassengerType represent schema of table passenger_type.
type PassengerType struct {
	UUID      string         `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid,omitempty"`
	Type      string         `gorm:"size:100;not null;"                        json:"type,omitempty"       form:"type"`
	CreatedAt time.Time      `                                                 json:"created_at,omitempty"`
	UpdatedAt time.Time      `                                                 json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `                                                 json:"deleted_at,omitempty"`
}

// PassengerTypeFaker represent content when generate fake data of passenger_type.
type PassengerTypeFaker struct {
	UUID string `faker:"uuid_hyphenated"`
	Type string `faker:"type"`
}

// PassengerTypes represent multiple PassengerType.
type PassengerTypes []*PassengerType

// DetailPassengerType represent format of detail PassengerType.
type DetailPassengerType struct {
	PassengerTypeFieldsForDetail
}

// DetailPassengerTypeList represent format of DetailPassengerType for PassengerType list.
type DetailPassengerTypeList struct {
	PassengerTypeFieldsForDetail
	PassengerTypeFieldsForList
}

// PassengerTypeFieldsForDetail represent fields of detail PassengerType.
type PassengerTypeFieldsForDetail struct {
	UUID string `json:"uuid"`
	Type string `json:"Type"`
}

// PassengerTypeFieldsForList represent fields of detail PassengerType for PassengerType list.
type PassengerTypeFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *PassengerType) TableName() string {
	return "passenger_types"
}

// FilterableFields return fields.
func (u *PassengerType) FilterableFields() []interface{} {
	return []interface{}{"uuid", "type"}
}

// Prepare will prepare submitted data of passenger_type.
func (u *PassengerType) Prepare() {
	u.Type = html.EscapeString(strings.TrimSpace(u.Type))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *PassengerType) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailPassengerTypes will return formatted passenger_type detail of multiple passenger_type.
func (passengerType PassengerTypes) DetailPassengerTypes() []interface{} {
	result := make([]interface{}, len(passengerType))
	for index, passenger_type := range passengerType {
		result[index] = passenger_type.DetailPassengerTypeList()
	}
	return result
}

// DetailPassengerType will return formatted passenger_type detail of passenger_type.
func (u *PassengerType) DetailPassengerType() interface{} {
	return &DetailPassengerType{
		PassengerTypeFieldsForDetail: PassengerTypeFieldsForDetail{
			UUID: u.UUID,
			Type: u.Type,
		},
	}
}

// DetailPassengerTypeList will return formatted passenger_type detail of passenger_type for passenger_type list.
func (u *PassengerType) DetailPassengerTypeList() interface{} {
	return &DetailPassengerTypeList{
		PassengerTypeFieldsForDetail: PassengerTypeFieldsForDetail{
			UUID: u.UUID,
			Type: u.Type,
		},
		PassengerTypeFieldsForList: PassengerTypeFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSavePassengerType will validate create a new passenger_type request.
func (u *PassengerType) ValidateSavePassengerType() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set(
			"type",
			u.Type,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		)
	return validation.Validate()
}

// ValidateUpdatePassengerType will validate update a new passenger_type request.
func (u *PassengerType) ValidateUpdatePassengerType() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set(
			"type",
			u.Type,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		)
	return validation.Validate()
}
