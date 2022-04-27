package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type orderStatusTypeApp struct {
	tr repository.OrderStatusTypeRepository
}

// orderStatusTypeApp implement the OrderStatusTypeAppInterface.
var _ OrderStatusTypeAppInterface = &orderStatusTypeApp{}

// OrderStatusTypeAppInterface is an interface.
type OrderStatusTypeAppInterface interface {
	SaveOrderStatusType(*entity.OrderStatusType) (*entity.OrderStatusType, map[string]string, error)
	UpdateOrderStatusType(
		UUID string,
		orderStatusType *entity.OrderStatusType,
	) (*entity.OrderStatusType, map[string]string, error)
	DeleteOrderStatusType(UUID string) error
	GetOrderStatusTypes(p *repository.Parameters) ([]*entity.OrderStatusType, *repository.Meta, error)
	GetOrderStatusType(UUID string) (*entity.OrderStatusType, error)
}

func (t orderStatusTypeApp) SaveOrderStatusType(
	orderStatusType *entity.OrderStatusType,
) (*entity.OrderStatusType, map[string]string, error) {
	return t.tr.SaveOrderStatusType(orderStatusType)
}

func (t orderStatusTypeApp) UpdateOrderStatusType(
	UUID string,
	orderStatusType *entity.OrderStatusType,
) (*entity.OrderStatusType, map[string]string, error) {
	return t.tr.UpdateOrderStatusType(UUID, orderStatusType)
}

func (t orderStatusTypeApp) DeleteOrderStatusType(UUID string) error {
	return t.tr.DeleteOrderStatusType(UUID)
}

func (t orderStatusTypeApp) GetOrderStatusTypes(
	p *repository.Parameters,
) ([]*entity.OrderStatusType, *repository.Meta, error) {
	return t.tr.GetOrderStatusTypes(p)
}

func (t orderStatusTypeApp) GetOrderStatusType(UUID string) (*entity.OrderStatusType, error) {
	return t.tr.GetOrderStatusType(UUID)
}
