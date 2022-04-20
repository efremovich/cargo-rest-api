package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type passengerApp struct {
	tr repository.PassengerRepository
}

// passengerApp implement the PassengerAppInterface.
var _ PassengerAppInterface = &passengerApp{}

// PassengerAppInterface is an interface.
type PassengerAppInterface interface {
	SavePassenger(*entity.Passenger) (*entity.Passenger, map[string]string, error)
	UpdatePassenger(UUID string, passenger *entity.Passenger) (*entity.Passenger, map[string]string, error)
	DeletePassenger(UUID string) error
	GetPassengers(p *repository.Parameters) ([]*entity.Passenger, *repository.Meta, error)
	GetPassenger(UUID string) (*entity.Passenger, error)
}

func (t passengerApp) SavePassenger(passenger *entity.Passenger) (*entity.Passenger, map[string]string, error) {
	return t.tr.SavePassenger(passenger)
}

func (t passengerApp) UpdatePassenger(
	UUID string,
	passenger *entity.Passenger,
) (*entity.Passenger, map[string]string, error) {
	return t.tr.UpdatePassenger(UUID, passenger)
}

func (t passengerApp) DeletePassenger(UUID string) error {
	return t.tr.DeletePassenger(UUID)
}

func (t passengerApp) GetPassengers(p *repository.Parameters) ([]*entity.Passenger, *repository.Meta, error) {
	return t.tr.GetPassengers(p)
}

func (t passengerApp) GetPassenger(UUID string) (*entity.Passenger, error) {
	return t.tr.GetPassenger(UUID)
}
