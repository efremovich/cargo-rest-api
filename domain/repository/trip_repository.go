package repository

import (
	"cargo-rest-api/domain/entity"
)

// TripRepository is an interface.
type TripRepository interface {
	SaveTrip(driver *entity.Trip) (*entity.Trip, map[string]string, error)
	UpdateTrip(UUID string, driver *entity.Trip) (*entity.Trip, map[string]string, error)
	DeleteTrip(UUID string) error
	GetTrip(UUID string) (*entity.Trip, error)
	GetTrips(parameters *Parameters) ([]*entity.Trip, *Meta, error)
}
