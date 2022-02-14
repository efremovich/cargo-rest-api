package entity

import (
	"html"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

// Vehicle represent schema of table sities.
type Vehicle struct {
	UUID      string    `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid"`
	Name      string    `gorm:"size:100;not null;" json:"name" form:"name"`
	Region    string    `gorm:"size:100;" json:"region" form:"region"`
	Latitude  string    `gorm:"size:100;" json:"latitude" form:"latitude"`
	Longitude string    `gorm:"size:100;" json:"longitude" form:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// VehicleFaker represent content when generate fake data of vehicle.
type VehicleFaker struct {
	UUID      string `faker:"uuid_hyphenated"`
	Name      string `faker:"name"`
	Region    string `faker:"region"`
	Latitude  string `faker:"latitude"`
	Longitude string `faker:"longitude"`
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
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Region    string `json:"region"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// VehicleFieldsForList represent fields of detail Vehicle for Vehicle list.
type VehicleFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *Vehicle) TableName() string {
	return "sities"
}

// FilterableFields return fields.
func (u *Vehicle) FilterableFields() []interface{} {
	return []interface{}{"uuid", "name", "region", "latitude", "longitude"}
}

// Prepare will prepare submitted data of vehicle.
func (u *Vehicle) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Region = html.EscapeString(strings.TrimSpace(u.Region))
	u.Latitude = html.EscapeString(strings.TrimSpace(u.Latitude))
	u.Longitude = html.EscapeString(strings.TrimSpace(u.Longitude))
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
			UUID:      u.UUID,
			Name:      u.Name,
			Region:    u.Region,
			Latitude:  u.Latitude,
			Longitude: u.Longitude,
		},
	}
}

// DetailVehicleList will return formatted vehicle detail of vehicle for vehicle list.
func (u *Vehicle) DetailVehicleList() interface{} {
	return &DetailVehicleList{
		VehicleFieldsForDetail: VehicleFieldsForDetail{
			UUID:      u.UUID,
			Name:      u.Name,
			Region:    u.Region,
			Latitude:  u.Latitude,
			Longitude: u.Longitude,
		},
		VehicleFieldsForList: VehicleFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}
