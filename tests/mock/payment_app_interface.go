package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// PaymentAppInterface is a mock of application.PaymentAppInterface.
type PaymentAppInterface struct {
	SavePaymentFn   func(*entity.Payment) (*entity.Payment, map[string]string, error)
	UpdatePaymentFn func(string, *entity.Payment) (*entity.Payment, map[string]string, error)
	DeletePaymentFn func(UUID string) error
	GetPaymentsFn   func(params *repository.Parameters) ([]*entity.Payment, *repository.Meta, error)
	GetPaymentFn    func(UUID string) (*entity.Payment, error)

	AddOrderPaymentFn    func(*entity.Payment) (*entity.Payment, map[string]string, error)
	DeleteOrderPaymentFn func(*entity.Payment) (*entity.Payment, map[string]string, error)
}

// SavePayment calls the SavePaymentFn.
func (u *PaymentAppInterface) SavePayment(payment *entity.Payment) (*entity.Payment, map[string]string, error) {
	return u.SavePaymentFn(payment)
}

// UpdatePayment calls the UpdatePaymentFn.
func (u *PaymentAppInterface) UpdatePayment(uuid string, payment *entity.Payment) (*entity.Payment, map[string]string, error) {
	return u.UpdatePaymentFn(uuid, payment)
}

// DeletePayment calls the DeletePaymentFn.
func (u *PaymentAppInterface) DeletePayment(uuid string) error {
	return u.DeletePaymentFn(uuid)
}

// GetPayments calls the GetPaymentsFn.
func (u *PaymentAppInterface) GetPayments(
	params *repository.Parameters,
) ([]*entity.Payment, *repository.Meta, error) {
	return u.GetPaymentsFn(params)
}

// GetPayment calls the GetPaymentFn.
func (u *PaymentAppInterface) GetPayment(uuid string) (*entity.Payment, error) {
	return u.GetPaymentFn(uuid)
}

// AddPaymentPrice calls the AddPaymentPriceFn.
func (u *PaymentAppInterface) AddOrderPayment(payment *entity.Payment) (*entity.Payment, map[string]string, error) {
	return u.AddOrderPaymentFn(payment)
}

// DeletePaymentPrice calls the DeletePaymentPriceFn.
func (u *PaymentAppInterface) DeleteOrderPayment(payment *entity.Payment) (*entity.Payment, map[string]string, error) {
	return u.DeleteOrderPaymentFn(payment)
}
