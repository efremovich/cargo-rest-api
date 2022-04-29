package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// routeFactory is a function uses to create []seed.Seed.
func routeFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.RouteFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		route := &entity.Route{
			UUID:         a.UUID,
			FromUUID:     a.FromUUID,
			ToUUID:       a.ToUUID,
			Distance:     a.Distance,
			DistanceTime: a.DistanceTime,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.FromUUID),
			Run: func(db *gorm.DB) error {
				_, errDB := createRoute(db, route)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createRoute will create fake route and insert into DB.
func createRoute(db *gorm.DB, route *entity.Route) (*entity.Route, error) {
	var routeExists entity.Route
	err := db.Where("from_uuid = ? and to_uuid = ? and distance = ?", route.FromUUID, route.ToUUID, route.Distance).
		Take(&routeExists).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(route).Error
			if err != nil {
				return route, err
			}
			return route, err
		}
		return route, err
	}
	return route, err
}
