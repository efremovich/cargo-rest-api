package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type countryApp struct {
	tr repository.CountryRepository
}

// countryApp implement the CountryAppInterface.
var _ CountryAppInterface = &countryApp{}

// CountryAppInterface is an interface.
type CountryAppInterface interface {
	SaveCountry(*entity.Country) (*entity.Country, map[string]string, error)
	UpdateCountry(UUID string, country *entity.Country) (*entity.Country, map[string]string, error)
	DeleteCountry(UUID string) error
	GetCountries(p *repository.Parameters) ([]entity.Country, interface{}, error)
	GetCountry(UUID string) (*entity.Country, error)
}

func (t countryApp) SaveCountry(country *entity.Country) (*entity.Country, map[string]string, error) {
	return t.tr.SaveCountry(country)
}

func (t countryApp) UpdateCountry(UUID string, country *entity.Country) (*entity.Country, map[string]string, error) {
	return t.tr.UpdateCountry(UUID, country)
}

func (t countryApp) DeleteCountry(UUID string) error {
	return t.tr.DeleteCountry(UUID)
}

func (t countryApp) GetCountries(p *repository.Parameters) ([]entity.Country, interface{}, error) {
	return t.tr.GetCountries(p)
}

func (t countryApp) GetCountry(UUID string) (*entity.Country, error) {
	return t.tr.GetCountry(UUID)
}
