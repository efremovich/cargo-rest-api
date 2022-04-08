package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// PassengerTypeAppInterface is a mock of application.PassengerTypeAppInterface.
type PassengerTypeAppInterface struct {
	SavePassengerTypeFn   func(*entity.PassengerType) (*entity.PassengerType, map[string]string, error)
	UpdatePassengerTypeFn func(string, *entity.PassengerType) (*entity.PassengerType, map[string]string, error)
	DeletePassengerTypeFn func(UUID string) error
	GetPassengerTypesFn   func(params *repository.Parameters) ([]*entity.PassengerType, *repository.Meta, error)
	GetPassengerTypeFn    func(UUID string) (*entity.PassengerType, error)
}

// SavePassengerType calls the SavePassengerTypeFn.
func (u *PassengerTypeAppInterface) SavePassengerType(passengerType *entity.PassengerType) (*entity.PassengerType, map[string]string, error) {
	return u.SavePassengerTypeFn(passengerType)
}

// UpdatePassengerType calls the UpdatePassengerTypeFn.
func (u *PassengerTypeAppInterface) UpdatePassengerType(uuid string, passengerType *entity.PassengerType) (*entity.PassengerType, map[string]string, error) {
	return u.UpdatePassengerTypeFn(uuid, passengerType)
}

// DeletePassengerType calls the DeletePassengerTypeFn.
func (u *PassengerTypeAppInterface) DeletePassengerType(uuid string) error {
	return u.DeletePassengerTypeFn(uuid)
}

// GetPassengerTypes calls the GetPassengerTypesFn.
func (u *PassengerTypeAppInterface) GetPassengerTypes(params *repository.Parameters) ([]*entity.PassengerType, *repository.Meta, error) {
	return u.GetPassengerTypesFn(params)
}

// GetPassengerType calls the GetPassengerTypeFn.
func (u *PassengerTypeAppInterface) GetPassengerType(uuid string) (*entity.PassengerType, error) {
	return u.GetPassengerTypeFn(uuid)
}
