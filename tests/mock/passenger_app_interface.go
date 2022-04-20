package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// PassengerAppInterface is a mock of application.PassengerAppInterface.
type PassengerAppInterface struct {
	SavePassengerFn   func(*entity.Passenger) (*entity.Passenger, map[string]string, error)
	UpdatePassengerFn func(string, *entity.Passenger) (*entity.Passenger, map[string]string, error)
	DeletePassengerFn func(UUID string) error
	GetPassengersFn   func(params *repository.Parameters) ([]*entity.Passenger, *repository.Meta, error)
	GetPassengerFn    func(UUID string) (*entity.Passenger, error)
}

// SavePassenger calls the SavePassengerFn.
func (u *PassengerAppInterface) SavePassenger(
	passenger *entity.Passenger,
) (*entity.Passenger, map[string]string, error) {
	return u.SavePassengerFn(passenger)
}

// UpdatePassenger calls the UpdatePassengerFn.
func (u *PassengerAppInterface) UpdatePassenger(
	uuid string,
	passenger *entity.Passenger,
) (*entity.Passenger, map[string]string, error) {
	return u.UpdatePassengerFn(uuid, passenger)
}

// DeletePassenger calls the DeletePassengerFn.
func (u *PassengerAppInterface) DeletePassenger(uuid string) error {
	return u.DeletePassengerFn(uuid)
}

// GetPassengers calls the GetPassengersFn.
func (u *PassengerAppInterface) GetPassengers(
	params *repository.Parameters,
) ([]*entity.Passenger, *repository.Meta, error) {
	return u.GetPassengersFn(params)
}

// GetPassenger calls the GetPassengerFn.
func (u *PassengerAppInterface) GetPassenger(uuid string) (*entity.Passenger, error) {
	return u.GetPassengerFn(uuid)
}
