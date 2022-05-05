package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"

	"gorm.io/gorm"
)

// OrderRepo is a struct to store db connection.
type OrderRepo struct {
	db *gorm.DB
}

// NewOrderRepository will initialize Order repository.
func NewOrderRepository(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db}
}

// OrderRepo implements the repository.orderRepository interface.
var _ repository.OrderRepository = &OrderRepo{}

// SaveOrder will create a new order.
func (r OrderRepo) SaveOrder(Order *entity.Order) (*entity.Order, map[string]string, error) {
	errDesc := map[string]string{}

	r.db.Model(&Order).Association("Passengers")

	err := r.db.Create(&Order).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Order, nil, nil
}

func (r OrderRepo) UpdateOrder(uuid string, order *entity.Order) (*entity.Order, map[string]string, error) {
	errDesc := map[string]string{}
	dirverData := &entity.Order{
		OrderDate:    order.OrderDate,
		TripUUID:     order.TripUUID,
		ExternalUUID: order.ExternalUUID,
		Seat:         order.Seat,
		StatusUUID:   order.StatusUUID,
		Passengers:   order.Passengers,
	}
	r.db.Model(order).Association("Passengers")

	err := r.db.First(&order, "uuid = ?", uuid).Updates(dirverData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextOrderInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextOrderNotFound
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return order, nil, nil
}

func (r OrderRepo) DeleteOrder(uuid string) error {
	var order entity.Order
	err := r.db.Where("uuid = ?", uuid).Take(&order).Delete(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextOrderNotFound
		}
		return err
	}
	return nil
}

func (r OrderRepo) GetOrder(uuid string) (*entity.Order, error) {
	var order entity.Order
	err := r.db.Preload("Trip").Preload("Passengers").Preload("Status").
		Where("uuid = ?", uuid).
		Take(&order).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextOrderNotFound
		}
	}
	return &order, nil
}

func (r OrderRepo) GetOrders(p *repository.Parameters) ([]*entity.Order, *repository.Meta, error) {
	var total int64
	var orders []*entity.Order
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&orders).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&orders).Error
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
	return orders, meta, nil
}
