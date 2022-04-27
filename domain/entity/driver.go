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

// Driver represent schema of table sities.
type Driver struct {
	UUID      string         `json:"uuid,omitempty"       gorm:"size:36;not null;uniqueIndex;primary_key;"`
	Name      string         `json:"name"                 gorm:"size:200;"`
	Vehicles  []Vehicle      `json:"vehicles"             gorm:"many2many:driver_vehicles;"`
	UserUUID  string         `json:"user_uuid,omitempty"  gorm:"size:36"`
	User      User           `json:"user"                 gorm:"foreignKey:UserUUID;association_foreignKey:UserUUID"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// DriverFaker represent content when generate fake data of driver_type.
type DriverFaker struct {
	UUID     string    `faker:"uuid_hyphenated"`
	Name     string    `faker:"name"`
	Vehicles []Vehicle `faker:"vehicles"`
	UserUUID string    `faker:"user_uuid"`
	User     User      `faker:"user"`
}

// Drivers represent multiple Driver.
type Drivers []*Driver

// DetailDriver represent format of detail Driver.
type DetailDriver struct {
	DriverFieldsForDetail
}

// DetailDriverList represent format of DetailDriver for Driver list.
type DetailDriverList struct {
	DriverFieldsForDetail
	DriverFieldsForList
}

// DriverFieldsForDetail represent fields of detail Driver.
type DriverFieldsForDetail struct {
	UUID     string
	Name     string        `json:"name"`
	Vehicles []interface{} `json:"vehicles"`
	UserUUID string        `json:"user_uuid"`
	User     User          `json:"user"`
}

// DriverFieldsForList represent fields of detail Driver for Driver list.
type DriverFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *Driver) TableName() string {
	return "drivers"
}

// FilterableFields return fields.
func (u *Driver) FilterableFields() []interface{} {
	return []interface{}{
		"uuid",
		"user_uuid",
	}
}

// Prepare will prepare submitted data of driver_type.
func (u *Driver) Prepare() {
	u.UUID = html.EscapeString(strings.TrimSpace(u.UUID))
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.UserUUID = html.EscapeString(strings.TrimSpace(u.UserUUID))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *Driver) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailDrivers will return formatted driver_type detail of multiple driver_type.
func (sities Drivers) DetailDrivers() []interface{} {
	result := make([]interface{}, len(sities))
	for index, driver_type := range sities {
		result[index] = driver_type.DetailDriverList()
	}
	return result
}

// DetailDriver will return formatted driver_type detail of driver_type.
func (u *Driver) DetailDriver() interface{} {
	return &DetailDriver{
		DriverFieldsForDetail: DriverFieldsForDetail{
			UUID:     u.UUID,
			Name:     u.Name,
			UserUUID: u.UserUUID,
		},
	}
}

// DetailDriverList will return formatted driver_type detail of driver_type for driver_type list.
func (u *Driver) DetailDriverList() interface{} {
	return &DetailDriverList{
		DriverFieldsForDetail: DriverFieldsForDetail{
			UUID:     u.UUID,
			Name:     u.Name,
			UserUUID: u.UserUUID,
		},
		DriverFieldsForList: DriverFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSaveDriver will validate create a new driver_type request.
func (u *Driver) ValidateSaveDriver() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("name", u.Name, validation.AddRule().IsAlphaUnicode().Required().Apply())
	return validation.Validate()
}

// ValidateUpdateDriver will validate update a new driver_type request.
func (u *Driver) ValidateUpdateDriver() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("name", u.Name, validation.AddRule().IsAlphaUnicode().Required().Apply())
	return validation.Validate()
}
