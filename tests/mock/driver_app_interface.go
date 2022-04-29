package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// DriverAppInterface is a mock of application.DriverAppInterface.
type DriverAppInterface struct {
	SaveDriverFn   func(*entity.Driver) (*entity.Driver, map[string]string, error)
	UpdateDriverFn func(string, *entity.Driver) (*entity.Driver, map[string]string, error)
	DeleteDriverFn func(UUID string) error
	GetDriversFn   func(params *repository.Parameters) ([]*entity.Driver, *repository.Meta, error)
	GetDriverFn    func(UUID string) (*entity.Driver, error)

	AddDriverVehicleFn    func(*entity.Driver) (*entity.Driver, map[string]string, error)
	DeleteDriverVehicleFn func(*entity.Driver) (*entity.Driver, map[string]string, error)
}

// SaveDriver calls the SaveDriverFn.
func (u *DriverAppInterface) SaveDriver(driver *entity.Driver) (*entity.Driver, map[string]string, error) {
	return u.SaveDriverFn(driver)
}

// UpdateDriver calls the UpdateDriverFn.
func (u *DriverAppInterface) UpdateDriver(uuid string, driver *entity.Driver) (*entity.Driver, map[string]string, error) {
	return u.UpdateDriverFn(uuid, driver)
}

// DeleteDriver calls the DeleteDriverFn.
func (u *DriverAppInterface) DeleteDriver(uuid string) error {
	return u.DeleteDriverFn(uuid)
}

// GetDrivers calls the GetDriversFn.
func (u *DriverAppInterface) GetDrivers(
	params *repository.Parameters,
) ([]*entity.Driver, *repository.Meta, error) {
	return u.GetDriversFn(params)
}

// GetDriver calls the GetDriverFn.
func (u *DriverAppInterface) GetDriver(uuid string) (*entity.Driver, error) {
	return u.GetDriverFn(uuid)
}

// AddDriverVehicle calls the AddDriverVehicleFn.
func (u *DriverAppInterface) AddDriverVehicle(driver *entity.Driver) (*entity.Driver, map[string]string, error) {
	return u.AddDriverVehicleFn(driver)
}

// DeleteDriverVehicle calls the DeleteDriverVehicleFn.
func (u *DriverAppInterface) DeleteDriverVehicle(driver *entity.Driver) (*entity.Driver, map[string]string, error) {
	return u.DeleteDriverVehicleFn(driver)
}
