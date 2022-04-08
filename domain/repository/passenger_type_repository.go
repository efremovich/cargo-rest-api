package repository

import (
	"cargo-rest-api/domain/entity"
)

// PassengerTypeRepository is an interface.
type PassengerTypeRepository interface {
	SavePassengerType(tour *entity.PassengerType) (*entity.PassengerType, map[string]string, error)
	UpdatePassengerType(UUID string, tour *entity.PassengerType) (*entity.PassengerType, map[string]string, error)
	DeletePassengerType(UUID string) error
	GetPassengerType(UUID string) (*entity.PassengerType, error)
	GetPassengerTypes(parameters *Parameters) ([]*entity.PassengerType, *Meta, error)
}
