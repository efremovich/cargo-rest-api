package repository

import (
	"cargo-rest-api/domain/entity"
)

// VehicleRepository is an interface.
type VehicleRepository interface {
	SaveVehicle(tour *entity.Vehicle) (*entity.Vehicle, map[string]string, error)
	UpdateVehicle(UUID string, tour *entity.Vehicle) (*entity.Vehicle, map[string]string, error)
	DeleteVehicle(UUID string) error
	GetVehicle(UUID string) (*entity.Vehicle, error)
	GetVehicles(parameters *Parameters) ([]*entity.Vehicle, *Meta, error)
}
