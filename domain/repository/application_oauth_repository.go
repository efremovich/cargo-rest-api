package repository

import "cargo-rest-api/domain/entity"

// ApplicationOauthRepository is an interface.
type ApplicationOauthRepository interface {
	SaveApplicationOauth(*entity.ApplicationOauth) (*entity.ApplicationOauth, map[string]string, error)
	UpdateApplicationOauth(string, *entity.ApplicationOauth) (*entity.ApplicationOauth, map[string]string, error)
	DeleteApplicationOauth(string) error
	GetApplicationOauth(string) (*entity.ApplicationOauth, error)
	GetApplicationOauthList(parameters *Parameters) ([]entity.ApplicationOauth, interface{}, error)
}
