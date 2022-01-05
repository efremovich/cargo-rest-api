package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"

	"gorm.io/gorm"
)

// CountryRepo is a struct to store db connection.
type CountryRepo struct {
	db *gorm.DB
}

// NewCountryRepository will initialize Tour repository.
func NewCountryRepository(db *gorm.DB) *CountryRepo {
	return &CountryRepo{db}
}

// CountryRepo implements the repository.countryRepository interface.
var _ repository.CountryRepository = &CountryRepo{}

// SaveCountry will create a new country.
func (r CountryRepo) SaveCountry(Country *entity.Country) (*entity.Country, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&Country).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Country, nil, nil
}

func (r CountryRepo) UpdateCountry(uuid string, tour *entity.Country) (*entity.Country, map[string]string, error) {
	panic("implement me")
}

func (r CountryRepo) DeleteCountry(uuid string) error {
	panic("implement me")
}

func (r CountryRepo) GetCountry(uuid string) (*entity.Country, error) {
	var country entity.Country
	err := r.db.Where("uuid = ?", uuid).Take(&country).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.ErrorTextRoleNotFound
	}
	return &country, nil
}

func (r CountryRepo) GetCountries(p *repository.Parameters) ([]entity.Country, interface{}, error) {
	var total int64
	var countries []entity.Country
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&countries).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&countries).Error
	if errTotal != nil {
		return nil, nil, errTotal
	}
	if errList != nil {
		return nil, nil, errList
	}
	if errors.Is(errList, gorm.ErrRecordNotFound) {
		return nil, nil, errList
	}
	meta := repository.NewMeta(p, total)
	return countries, meta, nil
}
