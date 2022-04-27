package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// OrderStatusTypeRepo is a struct to store db connection.
type OrderStatusTypeRepo struct {
	db *gorm.DB
}

// NewOrderStatusTypeRepository will initialize OrderStatusType repository.
func NewOrderStatusTypeRepository(db *gorm.DB) *OrderStatusTypeRepo {
	return &OrderStatusTypeRepo{db}
}

// OrderStatusTypeRepo implements the repository.regularityRepository interface.
var _ repository.OrderStatusTypeRepository = &OrderStatusTypeRepo{}

// SaveOrderStatusType will create a new regularity.
func (r OrderStatusTypeRepo) SaveOrderStatusType(
	OrderStatusType *entity.OrderStatusType,
) (*entity.OrderStatusType, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&OrderStatusType).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return OrderStatusType, nil, nil
}

func (r OrderStatusTypeRepo) UpdateOrderStatusType(
	uuid string,
	regularity *entity.OrderStatusType,
) (*entity.OrderStatusType, map[string]string, error) {
	errDesc := map[string]string{}
	regularityData := &entity.OrderStatusType{
		Type: regularity.Type,
	}

	err := r.db.First(&regularity, "uuid = ?", uuid).Updates(regularityData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextOrderStatusTypeInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextOrderStatusTypeNotFound
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

func (r OrderStatusTypeRepo) DeleteOrderStatusType(uuid string) error {
	var regularity entity.OrderStatusType
	err := r.db.Where("uuid = ?", uuid).Take(&regularity).Delete(&regularity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextOrderStatusTypeNotFound
		}
		return err
	}
	return nil
}

func (r OrderStatusTypeRepo) GetOrderStatusType(uuid string) (*entity.OrderStatusType, error) {
	var regularity entity.OrderStatusType
	err := r.db.Where("uuid = ?", uuid).Take(&regularity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextOrderStatusTypeNotFound
		}
	}
	return &regularity, nil
}

func (r OrderStatusTypeRepo) GetOrderStatusTypes(p *repository.Parameters) ([]*entity.OrderStatusType, *repository.Meta, error) {
	var total int64
	var regularitys []*entity.OrderStatusType
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
