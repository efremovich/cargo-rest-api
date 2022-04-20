package repository

import (
	"cargo-rest-api/domain/entity"
)

// PassengerRepository is an interface.
type PassengerRepository interface {
	SavePassenger(tour *entity.Passenger) (*entity.Passenger, map[string]string, error)
	UpdatePassenger(UUID string, tour *entity.Passenger) (*entity.Passenger, map[string]string, error)
	DeletePassenger(UUID string) error
	GetPassenger(UUID string) (*entity.Passenger, error)
	GetPassengers(parameters *Parameters) ([]*entity.Passenger, *Meta, error)
}
