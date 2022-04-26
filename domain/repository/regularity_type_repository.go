package repository

import (
	"cargo-rest-api/domain/entity"
)

// RegularityTypeRepository is an interface.
type RegularityTypeRepository interface {
	SaveRegularityType(tour *entity.RegularityType) (*entity.RegularityType, map[string]string, error)
	UpdateRegularityType(UUID string, tour *entity.RegularityType) (*entity.RegularityType, map[string]string, error)
	DeleteRegularityType(UUID string) error
	GetRegularityType(UUID string) (*entity.RegularityType, error)
	GetRegularityTypes(parameters *Parameters) ([]*entity.RegularityType, *Meta, error)
}
