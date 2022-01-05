package repository

import "cargo-rest-api/domain/entity"

// CountryRepository is an interface.
type CountryRepository interface {
	SaveCountry(tour *entity.Country) (*entity.Country, map[string]string, error)
	UpdateCountry(UUID string, tour *entity.Country) (*entity.Country, map[string]string, error)
	DeleteCountry(UUID string) error
	GetCountry(UUID string) (*entity.Country, error)
	GetCountries(parameters *Parameters) ([]entity.Country, interface{}, error)
}
