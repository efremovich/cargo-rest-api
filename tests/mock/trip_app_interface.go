package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// TripAppInterface is a mock of application.TripAppInterface.
type TripAppInterface struct {
	SaveTripFn   func(*entity.Trip) (*entity.Trip, map[string]string, error)
	UpdateTripFn func(string, *entity.Trip) (*entity.Trip, map[string]string, error)
	DeleteTripFn func(UUID string) error
	GetTripsFn   func(params *repository.Parameters) ([]*entity.Trip, *repository.Meta, error)
	GetTripFn    func(UUID string) (*entity.Trip, error)
}

// SaveTrip calls the SaveTripFn.
func (u *TripAppInterface) SaveTrip(trip *entity.Trip) (*entity.Trip, map[string]string, error) {
	return u.SaveTripFn(trip)
}

// UpdateTrip calls the UpdateTripFn.
func (u *TripAppInterface) UpdateTrip(uuid string, trip *entity.Trip) (*entity.Trip, map[string]string, error) {
	return u.UpdateTripFn(uuid, trip)
}

// DeleteTrip calls the DeleteTripFn.
func (u *TripAppInterface) DeleteTrip(uuid string) error {
	return u.DeleteTripFn(uuid)
}

// GetTrips calls the GetTripsFn.
func (u *TripAppInterface) GetTrips(
	params *repository.Parameters,
) ([]*entity.Trip, *repository.Meta, error) {
	return u.GetTripsFn(params)
}

// GetTrip calls the GetTripFn.
func (u *TripAppInterface) GetTrip(uuid string) (*entity.Trip, error) {
	return u.GetTripFn(uuid)
}
