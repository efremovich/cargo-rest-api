package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type paymentApp struct {
	tr repository.PaymentRepository
}

// paymentApp implement the PaymentAppInterface.
var _ PaymentAppInterface = &paymentApp{}

// PaymentAppInterface is an interface.
type PaymentAppInterface interface {
	SavePayment(*entity.Payment) (*entity.Payment, map[string]string, error)
	UpdatePayment(
		UUID string,
		payment *entity.Payment,
	) (*entity.Payment, map[string]string, error)
	DeletePayment(UUID string) error
	GetPayments(p *repository.Parameters) ([]*entity.Payment, *repository.Meta, error)
	GetPayment(UUID string) (*entity.Payment, error)

	AddOrderPayment(payment *entity.Payment) (*entity.Payment, map[string]string, error)
	DeleteOrderPayment(payment *entity.Payment) (*entity.Payment, map[string]string, error)
}

func (t paymentApp) SavePayment(
	payment *entity.Payment,
) (*entity.Payment, map[string]string, error) {
	return t.tr.SavePayment(payment)
}

func (t paymentApp) UpdatePayment(
	UUID string,
	payment *entity.Payment,
) (*entity.Payment, map[string]string, error) {
	return t.tr.UpdatePayment(UUID, payment)
}

func (t paymentApp) DeletePayment(UUID string) error {
	return t.tr.DeletePayment(UUID)
}

func (t paymentApp) GetPayments(
	p *repository.Parameters,
) ([]*entity.Payment, *repository.Meta, error) {
	return t.tr.GetPayments(p)
}

func (t paymentApp) GetPayment(UUID string) (*entity.Payment, error) {
	return t.tr.GetPayment(UUID)
}

func (t paymentApp) AddOrderPayment(payment *entity.Payment) (*entity.Payment, map[string]string, error) {
	return t.tr.AddOrderPayment(payment)
}

func (t paymentApp) DeleteOrderPayment(payment *entity.Payment) (*entity.Payment, map[string]string, error) {
	return t.tr.DeleteOrderPayment(payment)
}
