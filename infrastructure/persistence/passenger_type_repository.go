package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// PassengerTypeRepo is a struct to store db connection.
type PassengerTypeRepo struct {
	db *gorm.DB
}

// NewPassengerTypeRepository will initialize PassengerType repository.
func NewPassengerTypeRepository(db *gorm.DB) *PassengerTypeRepo {
	return &PassengerTypeRepo{db}
}

// PassengerTypeRepo implements the repository.passengerTypeRepository interface.
var _ repository.PassengerTypeRepository = &PassengerTypeRepo{}

// SavePassengerType will create a new passengerType.
func (r PassengerTypeRepo) SavePassengerType(PassengerType *entity.PassengerType) (*entity.PassengerType, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&PassengerType).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return PassengerType, nil, nil
}

func (r PassengerTypeRepo) UpdatePassengerType(uuid string, passengerType *entity.PassengerType) (*entity.PassengerType, map[string]string, error) {
	errDesc := map[string]string{}
	passengerTypeData := &entity.PassengerType{
		Type: passengerType.Type,
	}

	err := r.db.First(&passengerType, "uuid = ?", uuid).Updates(passengerTypeData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextPassengerTypeInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextPassengerTypeNotFound
		}
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			errDesc["type"] = exception.ErrorTextUserEmailAlreadyTaken.Error()
			return nil, errDesc, exception.ErrorTextUnprocessableEntity
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return passengerType, nil, nil
}

func (r PassengerTypeRepo) DeletePassengerType(uuid string) error {
	var passengerType entity.PassengerType
	err := r.db.Where("uuid = ?", uuid).Take(&passengerType).Delete(&passengerType).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextPassengerTypeNotFound
		}
		return err
	}
	return nil
}

func (r PassengerTypeRepo) GetPassengerType(uuid string) (*entity.PassengerType, error) {
	var passengerType entity.PassengerType
	err := r.db.Where("uuid = ?", uuid).Take(&passengerType).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextPassengerTypeNotFound
		}
	}
	return &passengerType, nil
}

func (r PassengerTypeRepo) GetPassengerTypes(p *repository.Parameters) ([]*entity.PassengerType, *repository.Meta, error) {
	var total int64
	var passengerTypes []*entity.PassengerType
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&passengerTypes).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&passengerTypes).Error
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
	return passengerTypes, meta, nil
}
