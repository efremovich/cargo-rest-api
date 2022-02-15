package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// VehicleAppInterface is a mock of application.VehicleAppInterface.
type VehicleAppInterface struct {
	SaveVehicleFn   func(*entity.Vehicle) (*entity.Vehicle, map[string]string, error)
	UpdateVehicleFn func(string, *entity.Vehicle) (*entity.Vehicle, map[string]string, error)
	DeleteVehicleFn func(UUID string) error
	GetVehiclesFn   func(params *repository.Parameters) ([]*entity.Vehicle, *repository.Meta, error)
	GetVehicleFn    func(UUID string) (*entity.Vehicle, error)
}

// SaveVehicle calls the SaveVehicleFn.
func (u *VehicleAppInterface) SaveVehicle(vehicle *entity.Vehicle) (*entity.Vehicle, map[string]string, error) {
	return u.SaveVehicleFn(vehicle)
}

// UpdateVehicle calls the UpdateVehicleFn.
func (u *VehicleAppInterface) UpdateVehicle(uuid string, vehicle *entity.Vehicle) (*entity.Vehicle, map[string]string, error) {
	return u.UpdateVehicleFn(uuid, vehicle)
}

// DeleteVehicle calls the DeleteVehicleFn.
func (u *VehicleAppInterface) DeleteVehicle(uuid string) error {
	return u.DeleteVehicleFn(uuid)
}

// GetVehicles calls the GetVehiclesFn.
func (u *VehicleAppInterface) GetVehicles(params *repository.Parameters) ([]*entity.Vehicle, *repository.Meta, error) {
	return u.GetVehiclesFn(params)
}

// GetVehicle calls the GetVehicleFn.
func (u *VehicleAppInterface) GetVehicle(uuid string) (*entity.Vehicle, error) {
	return u.GetVehicleFn(uuid)
}
