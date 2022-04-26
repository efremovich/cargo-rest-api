package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// RegularityTypeRepo is a struct to store db connection.
type RegularityTypeRepo struct {
	db *gorm.DB
}

// NewRegularityTypeRepository will initialize RegularityType repository.
func NewRegularityTypeRepository(db *gorm.DB) *RegularityTypeRepo {
	return &RegularityTypeRepo{db}
}

// RegularityTypeRepo implements the repository.regularityRepository interface.
var _ repository.RegularityTypeRepository = &RegularityTypeRepo{}

// SaveRegularityType will create a new regularity.
func (r RegularityTypeRepo) SaveRegularityType(
	RegularityType *entity.RegularityType,
) (*entity.RegularityType, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&RegularityType).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return RegularityType, nil, nil
}

func (r RegularityTypeRepo) UpdateRegularityType(
	uuid string,
	regularity *entity.RegularityType,
) (*entity.RegularityType, map[string]string, error) {
	errDesc := map[string]string{}
	regularityData := &entity.RegularityType{
		Type: regularity.Type,
	}

	err := r.db.First(&regularity, "uuid = ?", uuid).Updates(regularityData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextRegularityTypeInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextRegularityTypeNotFound
		}
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			errDesc["type"] = exception.ErrorTextUserEmailAlreadyTaken.Error()
			return nil, errDesc, exception.ErrorTextUnprocessableEntity
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return regularity, nil, nil
}

func (r RegularityTypeRepo) DeleteRegularityType(uuid string) error {
	var regularity entity.RegularityType
	err := r.db.Where("uuid = ?", uuid).Take(&regularity).Delete(&regularity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextRegularityTypeNotFound
		}
		return err
	}
	return nil
}

func (r RegularityTypeRepo) GetRegularityType(uuid string) (*entity.RegularityType, error) {
	var regularity entity.RegularityType
	err := r.db.Where("uuid = ?", uuid).Take(&regularity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextRegularityTypeNotFound
		}
	}
	return &regularity, nil
}

func (r RegularityTypeRepo) GetRegularityTypes(p *repository.Parameters) ([]*entity.RegularityType, *repository.Meta, error) {
	var total int64
	var regularitys []*entity.RegularityType
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&regularitys).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&regularitys).Error
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
	return regularitys, meta, nil
}
