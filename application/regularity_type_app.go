package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type regularityTypeApp struct {
	tr repository.RegularityTypeRepository
}

// regularityTypeApp implement the RegularityTypeAppInterface.
var _ RegularityTypeAppInterface = &regularityTypeApp{}

// RegularityTypeAppInterface is an interface.
type RegularityTypeAppInterface interface {
	SaveRegularityType(*entity.RegularityType) (*entity.RegularityType, map[string]string, error)
	UpdateRegularityType(
		UUID string,
		regularityType *entity.RegularityType,
	) (*entity.RegularityType, map[string]string, error)
	DeleteRegularityType(UUID string) error
	GetRegularityTypes(p *repository.Parameters) ([]*entity.RegularityType, *repository.Meta, error)
	GetRegularityType(UUID string) (*entity.RegularityType, error)
}

func (t regularityTypeApp) SaveRegularityType(
	regularityType *entity.RegularityType,
) (*entity.RegularityType, map[string]string, error) {
	return t.tr.SaveRegularityType(regularityType)
}

func (t regularityTypeApp) UpdateRegularityType(
	UUID string,
	regularityType *entity.RegularityType,
) (*entity.RegularityType, map[string]string, error) {
	return t.tr.UpdateRegularityType(UUID, regularityType)
}

func (t regularityTypeApp) DeleteRegularityType(UUID string) error {
	return t.tr.DeleteRegularityType(UUID)
}

func (t regularityTypeApp) GetRegularityTypes(
	p *repository.Parameters,
) ([]*entity.RegularityType, *repository.Meta, error) {
	return t.tr.GetRegularityTypes(p)
}

func (t regularityTypeApp) GetRegularityType(UUID string) (*entity.RegularityType, error) {
	return t.tr.GetRegularityType(UUID)
}
