package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// passengerFactory is a function uses to create []seed.Seed.
func passengerFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.PassengerFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		passenger := &entity.Passenger{
			UUID:              a.UUID,
			FirstName:         a.FirstName,
			LastName:          a.LastName,
			Patronomic:        a.Patronomic,
			BirthDay:          a.BirthDay,
			DocumentSeries:    a.DocumentSeries,
			DocumentNumber:    a.DocumentNumber,
			UserUUID:          a.UserUUID,
			PassengerTypeUUID: a.PassengerTypeUUID,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.FirstName),
			Run: func(db *gorm.DB) error {
				_, errDB := createPassenger(db, passenger)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createPassenger will create fake passenger and insert into DB.
func createPassenger(db *gorm.DB, passenger *entity.Passenger) (*entity.Passenger, error) {
	var passengerExists entity.Passenger
	err := db.Where("document_number = ? and document_series = ?", passenger.DocumentNumber, passenger.DocumentNumber).
		Take(&passengerExists).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(passenger).Error
			if err != nil {
				return passenger, err
			}
			return passenger, err
		}
		return passenger, err
	}
	return passenger, err
}
