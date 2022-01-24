package repository

import "cargo-rest-api/domain/entity"

// SityRepository is an interface.
type SityRepository interface {
	SaveSity(tour *entity.Sity) (*entity.Sity, map[string]string, error)
	UpdateSity(UUID string, tour *entity.Sity) (*entity.Sity, map[string]string, error)
	DeleteSity(UUID string) error
	GetSity(UUID string) (*entity.Sity, error)
	GetSities(parameters *Parameters) ([]entity.Sity, interface{}, error)
}
