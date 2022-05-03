package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// orderFactory is a function uses to create []seed.Seed.
func orderFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.OrderFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		order := &entity.Order{
			UUID:         a.UUID,
			OrdrDate:     a.OrderDate,
			PaymentDate:  a.PaymentDate,
			TripUUID:     a.TripUUID,
			Seat:         a.Seat,
			StatusUUID:   a.StatusUUID,
			ExternalUUID: a.ExternalUUID,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create order UUID:%s", a.UUID),
			Run: func(db *gorm.DB) error {
				_, errDB := createOrder(db, order)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createOrder will create fake order and insert into DB.
func createOrder(db *gorm.DB, order *entity.Order) (*entity.Order, error) {
	var orderExists entity.Order
	err := db.Where("external_uuid like %?%", order.ExternalUUID).Take(&orderExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(order).Error
			if err != nil {
				return order, err
			}
			return order, err
		}
		return order, err
	}
	return order, err
}
