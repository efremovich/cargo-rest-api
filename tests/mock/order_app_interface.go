package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// OrderAppInterface is a mock of application.OrderAppInterface.
type OrderAppInterface struct {
	SaveOrderFn   func(*entity.Order) (*entity.Order, map[string]string, error)
	UpdateOrderFn func(string, *entity.Order) (*entity.Order, map[string]string, error)
	DeleteOrderFn func(UUID string) error
	GetOrdersFn   func(params *repository.Parameters) ([]*entity.Order, *repository.Meta, error)
	GetOrderFn    func(UUID string) (*entity.Order, error)
}

// SaveOrder calls the SaveOrderFn.
func (u *OrderAppInterface) SaveOrder(order *entity.Order) (*entity.Order, map[string]string, error) {
	return u.SaveOrderFn(order)
}

// UpdateOrder calls the UpdateOrderFn.
func (u *OrderAppInterface) UpdateOrder(uuid string, order *entity.Order) (*entity.Order, map[string]string, error) {
	return u.UpdateOrderFn(uuid, order)
}

// DeleteOrder calls the DeleteOrderFn.
func (u *OrderAppInterface) DeleteOrder(uuid string) error {
	return u.DeleteOrderFn(uuid)
}

// GetOrders calls the GetOrdersFn.
func (u *OrderAppInterface) GetOrders(
	params *repository.Parameters,
) ([]*entity.Order, *repository.Meta, error) {
	return u.GetOrdersFn(params)
}

// GetOrder calls the GetOrderFn.
func (u *OrderAppInterface) GetOrder(uuid string) (*entity.Order, error) {
	return u.GetOrderFn(uuid)
}
