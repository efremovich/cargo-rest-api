package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// passengerTypeFactory is a function uses to create []seed.Seed.
func passengerTypeFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.PassengerTypeFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		passengerType := &entity.PassengerType{
			UUID: a.UUID,
			Type: a.Type,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.Type),
			Run: func(db *gorm.DB) error {
				_, errDB := createPassengerType(db, passengerType)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createPassengerType will create fake passengerType and insert into DB.
func createPassengerType(db *gorm.DB, passengerType *entity.PassengerType) (*entity.PassengerType, error) {
	var passengerTypeExists entity.PassengerType
	err := db.Where("type = ?", passengerType.Type).Take(&passengerTypeExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(passengerType).Error
			if err != nil {
				return passengerType, err
			}
			return passengerType, err
		}
		return passengerType, err
	}
	return passengerType, err
}
