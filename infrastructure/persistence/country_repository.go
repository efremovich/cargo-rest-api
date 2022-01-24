package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"

	"gorm.io/gorm"
)

// SityRepo is a struct to store db connection.
type SityRepo struct {
	db *gorm.DB
}

// NewSityRepository will initialize Tour repository.
func NewSityRepository(db *gorm.DB) *SityRepo {
	return &SityRepo{db}
}

// SityRepo implements the repository.sityRepository interface.
var _ repository.SityRepository = &SityRepo{}

// SaveSity will create a new sity.
func (r SityRepo) SaveSity(Sity *entity.Sity) (*entity.Sity, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&Sity).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Sity, nil, nil
}

func (r SityRepo) UpdateSity(uuid string, tour *entity.Sity) (*entity.Sity, map[string]string, error) {
	panic("implement me")
}

func (r SityRepo) DeleteSity(uuid string) error {
	panic("implement me")
}

func (r SityRepo) GetSity(uuid string) (*entity.Sity, error) {
	var sity entity.Sity
	err := r.db.Where("uuid = ?", uuid).Take(&sity).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.ErrorTextRoleNotFound
	}
	return &sity, nil
}

func (r SityRepo) GetSities(p *repository.Parameters) ([]entity.Sity, interface{}, error) {
	var total int64
	var sities []entity.Sity
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&sities).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&sities).Error
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
	return sities, meta, nil
}
