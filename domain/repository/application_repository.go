package repository

import "cargo-rest-api/domain/entity"

// ApplicationRepository is an interface.
type ApplicationRepository interface {
	SaveApplication(*entity.Application) (*entity.Application, map[string]string, error)
	UpdateApplication(string, *entity.Application) (*entity.Application, map[string]string, error)
	ActivateApplication(string) error
	DeactivateApplication(string) error
	DeleteApplication(string) error
	GetApplication(string) (*entity.Application, error)
	GetApplications(parameters *Parameters) ([]entity.Application, interface{}, error)
}
