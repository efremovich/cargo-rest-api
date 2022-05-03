package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type orderApp struct {
	tr repository.OrderRepository
}

// orderApp implement the OrderAppInterface.
var _ OrderAppInterface = &orderApp{}

// OrderAppInterface is an interface.
type OrderAppInterface interface {
	SaveOrder(*entity.Order) (*entity.Order, map[string]string, error)
	UpdateOrder(
		UUID string,
		order *entity.Order,
	) (*entity.Order, map[string]string, error)
	DeleteOrder(UUID string) error
	GetOrders(p *repository.Parameters) ([]*entity.Order, *repository.Meta, error)
	GetOrder(UUID string) (*entity.Order, error)
}

func (t orderApp) SaveOrder(
	order *entity.Order,
) (*entity.Order, map[string]string, error) {
	return t.tr.SaveOrder(order)
}

func (t orderApp) UpdateOrder(
	UUID string,
	order *entity.Order,
) (*entity.Order, map[string]string, error) {
	return t.tr.UpdateOrder(UUID, order)
}

func (t orderApp) DeleteOrder(UUID string) error {
	return t.tr.DeleteOrder(UUID)
}

func (t orderApp) GetOrders(
	p *repository.Parameters,
) ([]*entity.Order, *repository.Meta, error) {
	return t.tr.GetOrders(p)
}

func (t orderApp) GetOrder(UUID string) (*entity.Order, error) {
	return t.tr.GetOrder(UUID)
}
