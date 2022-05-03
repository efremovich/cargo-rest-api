package repository

import (
	"cargo-rest-api/domain/entity"
)

// OrderRepository is an interface.
type OrderRepository interface {
	SaveOrder(driver *entity.Order) (*entity.Order, map[string]string, error)
	UpdateOrder(UUID string, driver *entity.Order) (*entity.Order, map[string]string, error)
	DeleteOrder(UUID string) error
	GetOrder(UUID string) (*entity.Order, error)
	GetOrders(parameters *Parameters) ([]*entity.Order, *Meta, error)
}
