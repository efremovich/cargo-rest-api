package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type vehicleApp struct {
	tr repository.VehicleRepository
}

// vehicleApp implement the VehicleAppInterface.
var _ VehicleAppInterface = &vehicleApp{}

// VehicleAppInterface is an interface.
type VehicleAppInterface interface {
	SaveVehicle(*entity.Vehicle) (*entity.Vehicle, map[string]string, error)
	UpdateVehicle(UUID string, vehicle *entity.Vehicle) (*entity.Vehicle, map[string]string, error)
	DeleteVehicle(UUID string) error
	GetVehicles(p *repository.Parameters) ([]*entity.Vehicle, *repository.Meta, error)
	GetVehicle(UUID string) (*entity.Vehicle, error)
}

func (t vehicleApp) SaveVehicle(vehicle *entity.Vehicle) (*entity.Vehicle, map[string]string, error) {
	return t.tr.SaveVehicle(vehicle)
}

func (t vehicleApp) UpdateVehicle(UUID string, vehicle *entity.Vehicle) (*entity.Vehicle, map[string]string, error) {
	return t.tr.UpdateVehicle(UUID, vehicle)
}

func (t vehicleApp) DeleteVehicle(UUID string) error {
	return t.tr.DeleteVehicle(UUID)
}

func (t vehicleApp) GetVehicles(p *repository.Parameters) ([]*entity.Vehicle, *repository.Meta, error) {
	return t.tr.GetVehicles(p)
}

func (t vehicleApp) GetVehicle(UUID string) (*entity.Vehicle, error) {
	return t.tr.GetVehicle(UUID)
}
