package repository

import (
	"cargo-rest-api/domain/entity"
)

// OrderStatusTypeRepository is an interface.
type OrderStatusTypeRepository interface {
	SaveOrderStatusType(tour *entity.OrderStatusType) (*entity.OrderStatusType, map[string]string, error)
	UpdateOrderStatusType(UUID string, tour *entity.OrderStatusType) (*entity.OrderStatusType, map[string]string, error)
	DeleteOrderStatusType(UUID string) error
	GetOrderStatusType(UUID string) (*entity.OrderStatusType, error)
	GetOrderStatusTypes(parameters *Parameters) ([]*entity.OrderStatusType, *Meta, error)
}
