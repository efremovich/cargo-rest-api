package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// vehicleFactory is a function uses to create []seed.Seed.
func vehicleFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.VehicleFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		vehicle := &entity.Vehicle{
			UUID:          a.UUID,
			Model:         a.Model,
			RegCode:       a.RegCode,
			NumberOfSeats: a.NumberOfSeats,
			Class:         a.Class,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.RegCode),
			Run: func(db *gorm.DB) error {
				_, errDB := createVehicle(db, vehicle)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createVehicle will create fake vehicle and insert into DB.
func createVehicle(db *gorm.DB, vehicle *entity.Vehicle) (*entity.Vehicle, error) {
	var vehicleExists entity.Vehicle
	err := db.Where("reg_code = ?", vehicle.RegCode).Take(&vehicleExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(vehicle).Error
			if err != nil {
				return vehicle, err
			}
			return vehicle, err
		}
		return vehicle, err
	}
	return vehicle, err
}
