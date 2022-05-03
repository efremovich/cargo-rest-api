package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// paymentFactory is a function uses to create []seed.Seed.
func paymentFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.PaymentFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		payment := &entity.Payment{
			UUID:         a.UUID,
			PaymentDate:  a.PaymentDate,
			Amount:       a.Amount,
			UserUUID:     a.UserUUID,
			ExternalUUID: a.ExternalUUID,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create payment UUID:%s", a.UUID),
			Run: func(db *gorm.DB) error {
				_, errDB := createPayment(db, payment)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createPayment will create fake payment and insert into DB.
func createPayment(db *gorm.DB, payment *entity.Payment) (*entity.Payment, error) {
	var paymentExists entity.Payment
	err := db.Where("external_uuid = ? and uuid = ?", payment.ExternalUUID, payment.UUID).
		Take(&paymentExists).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(payment).Error
			if err != nil {
				return payment, err
			}
			return payment, err
		}
		return payment, err
	}
	return payment, err
}
