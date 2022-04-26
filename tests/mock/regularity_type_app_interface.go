package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// RegularityTypeAppInterface is a mock of application.RegularityTypeAppInterface.
type RegularityTypeAppInterface struct {
	SaveRegularityTypeFn   func(*entity.RegularityType) (*entity.RegularityType, map[string]string, error)
	UpdateRegularityTypeFn func(string, *entity.RegularityType) (*entity.RegularityType, map[string]string, error)
	DeleteRegularityTypeFn func(UUID string) error
	GetRegularityTypesFn   func(params *repository.Parameters) ([]*entity.RegularityType, *repository.Meta, error)
	GetRegularityTypeFn    func(UUID string) (*entity.RegularityType, error)
}

// SaveRegularityType calls the SaveRegularityTypeFn.
func (u *RegularityTypeAppInterface) SaveRegularityType(
	regularityType *entity.RegularityType,
) (*entity.RegularityType, map[string]string, error) {
	return u.SaveRegularityTypeFn(regularityType)
}

// UpdateRegularityType calls the UpdateRegularityTypeFn.
func (u *RegularityTypeAppInterface) UpdateRegularityType(
	uuid string,
	regularityType *entity.RegularityType,
) (*entity.RegularityType, map[string]string, error) {
	return u.UpdateRegularityTypeFn(uuid, regularityType)
}

// DeleteRegularityType calls the DeleteRegularityTypeFn.
func (u *RegularityTypeAppInterface) DeleteRegularityType(uuid string) error {
	return u.DeleteRegularityTypeFn(uuid)
}

// GetRegularityTypes calls the GetRegularityTypesFn.
func (u *RegularityTypeAppInterface) GetRegularityTypes(
	params *repository.Parameters,
) ([]*entity.RegularityType, *repository.Meta, error) {
	return u.GetRegularityTypesFn(params)
}

// GetRegularityType calls the GetRegularityTypeFn.
func (u *RegularityTypeAppInterface) GetRegularityType(uuid string) (*entity.RegularityType, error) {
	return u.GetRegularityTypeFn(uuid)
}
