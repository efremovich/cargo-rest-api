package repository

import (
	"cargo-rest-api/domain/entity"
)

// PaymentRepository is an interface.
type PaymentRepository interface {
	SavePayment(paynemnt *entity.Payment) (*entity.Payment, map[string]string, error)
	UpdatePayment(UUID string, tour *entity.Payment) (*entity.Payment, map[string]string, error)
	DeletePayment(UUID string) error
	GetPayment(UUID string) (*entity.Payment, error)
	GetPayments(parameters *Parameters) ([]*entity.Payment, *Meta, error)

	AddOrderPayment(paynemnt *entity.Payment) (*entity.Payment, map[string]string, error)
	DeleteOrderPayment(paynemnt *entity.Payment) (*entity.Payment, map[string]string, error)
}
