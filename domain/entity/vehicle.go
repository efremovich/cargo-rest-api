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

// Vehicle represent schema of table sities.
type Vehicle struct {
	UUID          string    `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid"`
	RegCode       string    `gorm:"size:100;not null;" json:"reg_code" form:"req_code"`
	Model         string    `gorm:"size:100;not null;" json:"model" form:"model"`
	NumberOfSeats string    `gorm:"size:100;" json:"number_of_seats" form:"number_of_seats"`
	Class         string    `gorm:"size:100;" json:"class" form:"class"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     gorm.DeletedAt
}

// VehicleFaker represent content when generate fake data of vehicle.
type VehicleFaker struct {
	UUID          string `faker:"uuid_hyphenated"`
	Model         string `faker:"model"`
	RegCode       string `faker:"reg_code"`
	NumberOfSeats string `faker:"number_of_seats"`
	Class         string `faker:"class"`
}

// Vehicles represent multiple Vehicle.
type Vehicles []*Vehicle

// DetailVehicle represent format of detail Vehicle.
type DetailVehicle struct {
	VehicleFieldsForDetail
}

// DetailVehicleList represent format of DetailVehicle for Vehicle list.
type DetailVehicleList struct {
	VehicleFieldsForDetail
	VehicleFieldsForList
}

// VehicleFieldsForDetail represent fields of detail Vehicle.
type VehicleFieldsForDetail struct {
	UUID          string `json:"uuid"`
	Model         string `json:"model"`
	RegCode       string `json:"reg_code"`
	NumberOfSeats string `json:"number_of_seats"`
	Class         string `json:"class"`
}

// VehicleFieldsForList represent fields of detail Vehicle for Vehicle list.
type VehicleFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *Vehicle) TableName() string {
	return "vehicles"
}

// FilterableFields return fields.
func (u *Vehicle) FilterableFields() []interface{} {
	return []interface{}{"uuid", "model", "reg_code", "number_of_seats", "class"}
}

// Prepare will prepare submitted data of vehicle.
func (u *Vehicle) Prepare() {
	u.Model = html.EscapeString(strings.TrimSpace(u.Model))
	u.RegCode = html.EscapeString(strings.TrimSpace(u.RegCode))
	u.NumberOfSeats = html.EscapeString(strings.TrimSpace(u.NumberOfSeats))
	u.Class = html.EscapeString(strings.TrimSpace(u.Class))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *Vehicle) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailVehicles will return formatted vehicle detail of multiple vehicle.
func (sities Vehicles) DetailVehicles() []interface{} {
	result := make([]interface{}, len(sities))
	for index, vehicle := range sities {
		result[index] = vehicle.DetailVehicleList()
	}
	return result
}

// DetailVehicle will return formatted vehicle detail of vehicle.
func (u *Vehicle) DetailVehicle() interface{} {
	return &DetailVehicle{
		VehicleFieldsForDetail: VehicleFieldsForDetail{
			UUID:          u.UUID,
			Model:         u.Model,
			RegCode:       u.RegCode,
			NumberOfSeats: u.NumberOfSeats,
			Class:         u.Class,
		},
	}
}

// DetailVehicleList will return formatted vehicle detail of vehicle for vehicle list.
func (u *Vehicle) DetailVehicleList() interface{} {
	return &DetailVehicleList{
		VehicleFieldsForDetail: VehicleFieldsForDetail{
			UUID:          u.UUID,
			Model:         u.Model,
			RegCode:       u.RegCode,
			NumberOfSeats: u.NumberOfSeats,
			Class:         u.Class,
		},
		VehicleFieldsForList: VehicleFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSaveVehicle will validate create a new vehicle request.
func (u *Vehicle) ValidateSaveVehicle() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("model", u.Model, validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply()).
		Set("reg_code", u.RegCode, validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply()).
		Set("number_of_seats", u.NumberOfSeats, validation.AddRule().Required().IsAlphaNumericSpace().Length(1, 64).Apply()).
		Set("class", u.Class, validation.AddRule().Required().IsAlphaSpace().Length(3, 64).Apply())
	return validation.Validate()
}

// ValidateUpdateVehicle will validate update a new vehicle request.
func (u *Vehicle) ValidateUpdateVehicle() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("model", u.Model, validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply()).
		Set("reg_code", u.RegCode, validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply()).
		Set("number_of_seats", u.NumberOfSeats, validation.AddRule().Required().IsAlphaNumericSpace().Length(1, 64).Apply()).
		Set("class", u.Class, validation.AddRule().Required().IsAlphaSpace().Length(3, 64).Apply())
	return validation.Validate()
}
