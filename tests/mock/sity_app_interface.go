package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// SityAppInterface is a mock of application.SityAppInterface.
type SityAppInterface struct {
	SaveSityFn   func(*entity.Sity) (*entity.Sity, map[string]string, error)
	UpdateSityFn func(string, *entity.Sity) (*entity.Sity, map[string]string, error)
	DeleteSityFn func(UUID string) error
	GetSitiesFn  func(params *repository.Parameters) ([]*entity.Sity, *repository.Meta, error)
	GetSityFn    func(UUID string) (*entity.Sity, error)
}

// SaveSity calls the SaveSityFn.
func (u *SityAppInterface) SaveSity(sity *entity.Sity) (*entity.Sity, map[string]string, error) {
	return u.SaveSityFn(sity)
}

// UpdateSity calls the UpdateSityFn.
func (u *SityAppInterface) UpdateSity(uuid string, sity *entity.Sity) (*entity.Sity, map[string]string, error) {
	return u.UpdateSityFn(uuid, sity)
}

// DeleteSity calls the DeleteSityFn.
func (u *SityAppInterface) DeleteSity(uuid string) error {
	return u.DeleteSityFn(uuid)
}

// GetSitys calls the GetSitysFn.
func (u *SityAppInterface) GetSities(params *repository.Parameters) ([]*entity.Sity, *repository.Meta, error) {
	return u.GetSitiesFn(params)
}

// GetSity calls the GetSityFn.
func (u *SityAppInterface) GetSity(uuid string) (*entity.Sity, error) {
	return u.GetSityFn(uuid)
}
