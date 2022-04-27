package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// OrderStatusTypeAppInterface is a mock of application.OrderStatusTypeAppInterface.
type OrderStatusTypeAppInterface struct {
	SaveOrderStatusTypeFn   func(*entity.OrderStatusType) (*entity.OrderStatusType, map[string]string, error)
	UpdateOrderStatusTypeFn func(string, *entity.OrderStatusType) (*entity.OrderStatusType, map[string]string, error)
	DeleteOrderStatusTypeFn func(UUID string) error
	GetOrderStatusTypesFn   func(params *repository.Parameters) ([]*entity.OrderStatusType, *repository.Meta, error)
	GetOrderStatusTypeFn    func(UUID string) (*entity.OrderStatusType, error)
}

// SaveOrderStatusType calls the SaveOrderStatusTypeFn.
func (u *OrderStatusTypeAppInterface) SaveOrderStatusType(
	orderStatusType *entity.OrderStatusType,
) (*entity.OrderStatusType, map[string]string, error) {
	return u.SaveOrderStatusTypeFn(orderStatusType)
}

// UpdateOrderStatusType calls the UpdateOrderStatusTypeFn.
func (u *OrderStatusTypeAppInterface) UpdateOrderStatusType(
	uuid string,
	orderStatusType *entity.OrderStatusType,
) (*entity.OrderStatusType, map[string]string, error) {
	return u.UpdateOrderStatusTypeFn(uuid, orderStatusType)
}

// DeleteOrderStatusType calls the DeleteOrderStatusTypeFn.
func (u *OrderStatusTypeAppInterface) DeleteOrderStatusType(uuid string) error {
	return u.DeleteOrderStatusTypeFn(uuid)
}

// GetOrderStatusTypes calls the GetOrderStatusTypesFn.
func (u *OrderStatusTypeAppInterface) GetOrderStatusTypes(
	params *repository.Parameters,
) ([]*entity.OrderStatusType, *repository.Meta, error) {
	return u.GetOrderStatusTypesFn(params)
}

// GetOrderStatusType calls the GetOrderStatusTypeFn.
func (u *OrderStatusTypeAppInterface) GetOrderStatusType(uuid string) (*entity.OrderStatusType, error) {
	return u.GetOrderStatusTypeFn(uuid)
}
