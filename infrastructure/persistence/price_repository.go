package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// PriceRepo is a struct to store db connection.
type PriceRepo struct {
	db *gorm.DB
}

// NewPriceRepository will initialize Price repository.
func NewPriceRepository(db *gorm.DB) *PriceRepo {
	return &PriceRepo{db}
}

// PriceRepo implements the repository.priceRepository interface.
var _ repository.PriceRepository = &PriceRepo{}

// SavePrice will create a new price.
func (r PriceRepo) SavePrice(Price *entity.Price) (*entity.Price, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&Price).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Price, nil, nil
}

func (r PriceRepo) UpdatePrice(uuid string, price *entity.Price) (*entity.Price, map[string]string, error) {
	errDesc := map[string]string{}
	priceData := &entity.Price{
		PassengerTypeID: price.PassengerTypeID,
		Price:           price.Price,
	}

	err := r.db.First(&price, "uuid = ?", uuid).Updates(priceData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextPriceInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextPriceNotFound
		}
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			errDesc["type"] = exception.ErrorTextUserEmailAlreadyTaken.Error()
			return nil, errDesc, exception.ErrorTextUnprocessableEntity
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return price, nil, nil
}

func (r PriceRepo) DeletePrice(uuid string) error {
	var price entity.Price
	err := r.db.Where("uuid = ?", uuid).Take(&price).Delete(&price).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextPriceNotFound
		}
		return err
	}
	return nil
}

func (r PriceRepo) GetPrice(uuid string) (*entity.Price, error) {
	var price entity.Price
	err := r.db.Where("uuid = ?", uuid).Take(&price).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextPriceNotFound
		}
	}
	return &price, nil
}

func (r PriceRepo) GetPrices(p *repository.Parameters) ([]*entity.Price, *repository.Meta, error) {
	var total int64
	var prices []*entity.Price
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&prices).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&prices).Error
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
	return prices, meta, nil
}
