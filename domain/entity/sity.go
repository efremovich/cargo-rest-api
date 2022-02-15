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

// Sity represent schema of table sities.
type Sity struct {
	UUID      string    `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid"`
	Name      string    `gorm:"size:100;not null;" json:"name" form:"name"`
	Region    string    `gorm:"size:100;" json:"region" form:"region"`
	Latitude  string    `gorm:"size:100;" json:"latitude" form:"latitude"`
	Longitude string    `gorm:"size:100;" json:"longitude" form:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// SityFaker represent content when generate fake data of sity.
type SityFaker struct {
	UUID      string `faker:"uuid_hyphenated"`
	Name      string `faker:"name"`
	Region    string `faker:"region"`
	Latitude  string `faker:"latitude"`
	Longitude string `faker:"longitude"`
}

// Sities represent multiple Sity.
type Sities []*Sity

// DetailSity represent format of detail Sity.
type DetailSity struct {
	SityFieldsForDetail
}

// DetailSityList represent format of DetailSity for Sity list.
type DetailSityList struct {
	SityFieldsForDetail
	SityFieldsForList
}

// SityFieldsForDetail represent fields of detail Sity.
type SityFieldsForDetail struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Region    string `json:"region"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// SityFieldsForList represent fields of detail Sity for Sity list.
type SityFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *Sity) TableName() string {
	return "sities"
}

// FilterableFields return fields.
func (u *Sity) FilterableFields() []interface{} {
	return []interface{}{"uuid", "name", "region", "latitude", "longitude"}
}

// Prepare will prepare submitted data of sity.
func (u *Sity) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Region = html.EscapeString(strings.TrimSpace(u.Region))
	u.Latitude = html.EscapeString(strings.TrimSpace(u.Latitude))
	u.Longitude = html.EscapeString(strings.TrimSpace(u.Longitude))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *Sity) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailSities will return formatted sity detail of multiple sity.
func (sities Sities) DetailSities() []interface{} {
	result := make([]interface{}, len(sities))
	for index, sity := range sities {
		result[index] = sity.DetailSityList()
	}
	return result
}

// DetailSity will return formatted sity detail of sity.
func (u *Sity) DetailSity() interface{} {
	return &DetailSity{
		SityFieldsForDetail: SityFieldsForDetail{
			UUID:      u.UUID,
			Name:      u.Name,
			Region:    u.Region,
			Latitude:  u.Latitude,
			Longitude: u.Longitude,
		},
	}
}

// DetailSityList will return formatted sity detail of sity for sity list.
func (u *Sity) DetailSityList() interface{} {
	return &DetailSityList{
		SityFieldsForDetail: SityFieldsForDetail{
			UUID:      u.UUID,
			Name:      u.Name,
			Region:    u.Region,
			Latitude:  u.Latitude,
			Longitude: u.Longitude,
		},
		SityFieldsForList: SityFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSaveSity will validate create a new sity request.
func (u *Sity) ValidateSaveSity() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("name", u.Name, validation.AddRule().Required().IsAlphaSpace().Length(3, 64).Apply()).
		Set("region", u.Region, validation.AddRule().Required().IsAlphaNumericSpace().Length(3, 64).Apply()).
		Set("latitude", u.Latitude, validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply()).
		Set("longitude", u.Longitude, validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply())
	return validation.Validate()
}

// ValidateUpdateSity will validate update a new sity request.
func (u *Sity) ValidateUpdateSity() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("name", u.Name, validation.AddRule().Required().IsAlphaSpace().Length(3, 64).Apply()).
		Set("region", u.Region, validation.AddRule().Required().IsAlphaNumericSpace().Length(3, 64).Apply()).
		Set("latitude", u.Latitude, validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply()).
		Set("longitude", u.Longitude, validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply())
	return validation.Validate()
}
