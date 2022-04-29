package repository

import (
	"cargo-rest-api/domain/entity"
)

// DriverRepository is an interface.
type DriverRepository interface {
	SaveDriver(driver *entity.Driver) (*entity.Driver, map[string]string, error)
	UpdateDriver(UUID string, driver *entity.Driver) (*entity.Driver, map[string]string, error)
	DeleteDriver(UUID string) error
	GetDriver(UUID string) (*entity.Driver, error)
	GetDrivers(parameters *Parameters) ([]*entity.Driver, *Meta, error)

	AddDriverVehicle(driver *entity.Driver) (*entity.Driver, map[string]string, error)
	DeleteDriverVehicle(driver *entity.Driver) (*entity.Driver, map[string]string, error)
}
