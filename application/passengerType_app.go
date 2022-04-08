package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type passengerTypeApp struct {
	tr repository.PassengerTypeRepository
}

// passengerTypeApp implement the PassengerTypeAppInterface.
var _ PassengerTypeAppInterface = &passengerTypeApp{}

// PassengerTypeAppInterface is an interface.
type PassengerTypeAppInterface interface {
	SavePassengerType(*entity.PassengerType) (*entity.PassengerType, map[string]string, error)
	UpdatePassengerType(UUID string, passengerType *entity.PassengerType) (*entity.PassengerType, map[string]string, error)
	DeletePassengerType(UUID string) error
	GetPassengerTypes(p *repository.Parameters) ([]*entity.PassengerType, *repository.Meta, error)
	GetPassengerType(UUID string) (*entity.PassengerType, error)
}

func (t passengerTypeApp) SavePassengerType(passengerType *entity.PassengerType) (*entity.PassengerType, map[string]string, error) {
	return t.tr.SavePassengerType(passengerType)
}

func (t passengerTypeApp) UpdatePassengerType(UUID string, passengerType *entity.PassengerType) (*entity.PassengerType, map[string]string, error) {
	return t.tr.UpdatePassengerType(UUID, passengerType)
}

func (t passengerTypeApp) DeletePassengerType(UUID string) error {
	return t.tr.DeletePassengerType(UUID)
}

func (t passengerTypeApp) GetPassengerTypes(p *repository.Parameters) ([]*entity.PassengerType, *repository.Meta, error) {
	return t.tr.GetPassengerTypes(p)
}

func (t passengerTypeApp) GetPassengerType(UUID string) (*entity.PassengerType, error) {
	return t.tr.GetPassengerType(UUID)
}
