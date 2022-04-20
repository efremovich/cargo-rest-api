package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// priceFactory is a function uses to create []seed.Seed.
func priceFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.PriceFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		price := &entity.Price{
			UUID:              a.UUID,
			PassengerTypeUUID: a.PassengerTypeUUID,
			Price:             a.Price,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.PassengerTypeUUID),
			Run: func(db *gorm.DB) error {
				_, errDB := createPrice(db, price)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createPrice will create fake price and insert into DB.
func createPrice(db *gorm.DB, price *entity.Price) (*entity.Price, error) {
	var priceExists entity.Price
	err := db.Where("passenger_type_uuid = ?", price.PassengerTypeUUID).Take(&priceExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(price).Error
			if err != nil {
				return price, err
			}
			return price, err
		}
		return price, err
	}
	return price, err
}
