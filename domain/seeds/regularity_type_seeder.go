package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// regularityTypeFactory is a function uses to create []seed.Seed.
func regularityTypeFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.RegularityTypeFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		regularityType := &entity.RegularityType{
			UUID: a.UUID,
			Type: a.Type,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.Type),
			Run: func(db *gorm.DB) error {
				_, errDB := createRegularityType(db, regularityType)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createRegularityType will create fake regularityType and insert into DB.
func createRegularityType(db *gorm.DB, regularityType *entity.RegularityType) (*entity.RegularityType, error) {
	var regularityTypeExists entity.RegularityType
	err := db.Where("type = ?", regularityType.Type).Take(&regularityTypeExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(regularityType).Error
			if err != nil {
				return regularityType, err
			}
			return regularityType, err
		}
		return regularityType, err
	}
	return regularityType, err
}
