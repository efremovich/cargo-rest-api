package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// driverFactory is a function uses to create []seed.Seed.
func driverFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.DriverFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		driver := &entity.Driver{
			UUID:     a.UUID,
			Name:     a.Name,
			UserUUID: a.UserUUID,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.Name),
			Run: func(db *gorm.DB) error {
				_, errDB := createDriver(db, driver)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createDriver will create fake driver and insert into DB.
func createDriver(db *gorm.DB, driver *entity.Driver) (*entity.Driver, error) {
	var driverExists entity.Driver
	err := db.Where("name = ?", driver.Name).Take(&driverExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(driver).Error
			if err != nil {
				return driver, err
			}
			return driver, err
		}
		return driver, err
	}
	return driver, err
}
