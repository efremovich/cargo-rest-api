package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// sityFactory is a function uses to create []seed.Seed.
func sityFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.SityFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		sity := &entity.Sity{
			UUID:      a.UUID,
			Name:      a.Name,
			Region:    a.Region,
			Latitude:  a.Latitude,
			Longitude: a.Longitude,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.Name),
			Run: func(db *gorm.DB) error {
				_, errDB := createSity(db, sity)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createSity will create fake sity and insert into DB.
func createSity(db *gorm.DB, sity *entity.Sity) (*entity.Sity, error) {
	var sityExists entity.Sity
	err := db.Where("name = ?", sity.Name).Take(&sityExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(sity).Error
			if err != nil {
				return sity, err
			}
			return sity, err
		}
		return sity, err
	}
	return sity, err
}
