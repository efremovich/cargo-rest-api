package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"

	"gorm.io/gorm"
)

// PaymentRepo is a struct to store db connection.
type PaymentRepo struct {
	db *gorm.DB
}

// NewPaymentRepository will initialize Payment repository.
func NewPaymentRepository(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{db}
}

// PaymentRepo implements the repository.paymentRepository interface.
var _ repository.PaymentRepository = &PaymentRepo{}

// SavePayment will create a new payment.
func (r PaymentRepo) SavePayment(Payment *entity.Payment) (*entity.Payment, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Model(&Payment).Association("Orders").Error
	err = r.db.Create(&Payment).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Payment, nil, nil
}

func (r PaymentRepo) UpdatePayment(uuid string, payment *entity.Payment) (*entity.Payment, map[string]string, error) {
	errDesc := map[string]string{}
	dirverData := &entity.Payment{
		PaymentDate:  payment.PaymentDate,
		Amount:       payment.Amount,
		UserUUID:     payment.UserUUID,
		Orders:       payment.Orders,
		ExternalUUID: payment.ExternalUUID,
	}
	r.db.Model(payment).Association("Orders")

	err := r.db.First(&payment, "uuid = ?", uuid).Updates(dirverData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextPaymentInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextPaymentNotFound
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return payment, nil, nil
}

func (r PaymentRepo) DeletePayment(uuid string) error {
	var payment entity.Payment
	err := r.db.Where("uuid = ?", uuid).Take(&payment).Delete(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextPaymentNotFound
		}
		return err
	}
	return nil
}

func (r PaymentRepo) GetPayment(uuid string) (*entity.Payment, error) {
	var payment entity.Payment
	err := r.db.Preload("Orders").Where("uuid = ?", uuid).Take(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextPaymentNotFound
		}
	}
	return &payment, nil
}

func (r PaymentRepo) GetPayments(p *repository.Parameters) ([]*entity.Payment, *repository.Meta, error) {
	var total int64
	var payments []*entity.Payment
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&payments).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&payments).Error
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
	return payments, meta, nil
}

// AddPaymentVehicle implements repository.PaymentRepository
func (r PaymentRepo) AddOrderPayment(payment *entity.Payment) (*entity.Payment, map[string]string, error) {
	errDesc := map[string]string{}

	err := r.db.Model(payment).Association("Orders").Append(payment.Orders)
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return payment, nil, nil
}

// AddPaymentVehicle implements repository.PaymentRepository
func (r PaymentRepo) DeleteOrderPayment(payment *entity.Payment) (*entity.Payment, map[string]string, error) {
	errDesc := map[string]string{}

	errDelete := r.db.Model(payment).Association("Orders").Delete(payment.Orders)
	if errDelete != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return payment, nil, nil
}
