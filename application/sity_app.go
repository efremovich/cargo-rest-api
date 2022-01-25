package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type sityApp struct {
	tr repository.SityRepository
}

// sityApp implement the SityAppInterface.
var _ SityAppInterface = &sityApp{}

// SityAppInterface is an interface.
type SityAppInterface interface {
	SaveSity(*entity.Sity) (*entity.Sity, map[string]string, error)
	UpdateSity(UUID string, sity *entity.Sity) (*entity.Sity, map[string]string, error)
	DeleteSity(UUID string) error
	GetSities(p *repository.Parameters) ([]entity.Sity, *repository.Meta, error)
	GetSity(UUID string) (*entity.Sity, error)
}

func (t sityApp) SaveSity(sity *entity.Sity) (*entity.Sity, map[string]string, error) {
	return t.tr.SaveSity(sity)
}

func (t sityApp) UpdateSity(UUID string, sity *entity.Sity) (*entity.Sity, map[string]string, error) {
	return t.tr.UpdateSity(UUID, sity)
}

func (t sityApp) DeleteSity(UUID string) error {
	return t.tr.DeleteSity(UUID)
}

func (t sityApp) GetSities(p *repository.Parameters) ([]entity.Sity, *repository.Meta, error) {
	return t.tr.GetSities(p)
}

func (t sityApp) GetSity(UUID string) (*entity.Sity, error) {
	return t.tr.GetSity(UUID)
}
