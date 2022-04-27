package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// orderStatusTypeFactory is a function uses to create []seed.Seed.
func orderStatusTypeFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.OrderStatusTypeFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		orderStatusType := &entity.OrderStatusType{
			UUID: a.UUID,
			Type: a.Type,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.Type),
			Run: func(db *gorm.DB) error {
				_, errDB := createOrderStatusType(db, orderStatusType)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createOrderStatusType will create fake orderStatusType and insert into DB.
func createOrderStatusType(db *gorm.DB, orderStatusType *entity.OrderStatusType) (*entity.OrderStatusType, error) {
	var orderStatusTypeExists entity.OrderStatusType
	err := db.Where("type = ?", orderStatusType.Type).Take(&orderStatusTypeExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(orderStatusType).Error
			if err != nil {
				return orderStatusType, err
			}
			return orderStatusType, err
		}
		return orderStatusType, err
	}
	return orderStatusType, err
}
