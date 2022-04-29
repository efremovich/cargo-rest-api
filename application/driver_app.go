package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type driverApp struct {
	tr repository.DriverRepository
}

// AddDiriverVehicle implements DriverAppInterface
// driverApp implement the DriverAppInterface.
var _ DriverAppInterface = &driverApp{}

// DriverAppInterface is an interface.
type DriverAppInterface interface {
	SaveDriver(*entity.Driver) (*entity.Driver, map[string]string, error)
	UpdateDriver(UUID string, driver *entity.Driver) (*entity.Driver, map[string]string, error)
	DeleteDriver(UUID string) error
	GetDrivers(p *repository.Parameters) ([]*entity.Driver, *repository.Meta, error)
	GetDriver(UUID string) (*entity.Driver, error)

	AddDriverVehicle(driver *entity.Driver) (*entity.Driver, map[string]string, error)
	DeleteDriverVehicle(dirive *entity.Driver) (*entity.Driver, map[string]string, error)
}

func (t driverApp) SaveDriver(driver *entity.Driver) (*entity.Driver, map[string]string, error) {
	return t.tr.SaveDriver(driver)
}

func (t driverApp) UpdateDriver(UUID string, driver *entity.Driver) (*entity.Driver, map[string]string, error) {
	return t.tr.UpdateDriver(UUID, driver)
}

func (t driverApp) DeleteDriver(UUID string) error {
	return t.tr.DeleteDriver(UUID)
}

func (t driverApp) GetDrivers(p *repository.Parameters) ([]*entity.Driver, *repository.Meta, error) {
	return t.tr.GetDrivers(p)
}

func (t driverApp) GetDriver(UUID string) (*entity.Driver, error) {
	return t.tr.GetDriver(UUID)
}

func (t driverApp) AddDriverVehicle(driver *entity.Driver) (*entity.Driver, map[string]string, error) {
	return t.tr.AddDriverVehicle(driver)
}

func (t driverApp) DeleteDriverVehicle(driver *entity.Driver) (*entity.Driver, map[string]string, error) {
	return t.tr.DeleteDriverVehicle(driver)
}
