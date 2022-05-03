package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// tripFactory is a function uses to create []seed.Seed.
func tripFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.TripFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		trip := &entity.Trip{
			UUID:               a.UUID,
			RouteUUID:          a.RouteUUID,
			VehicleUUID:        a.VehicleUUID,
			DepartureTime:      a.DepartureTime,
			ArravialTive:       a.ArravialTive,
			RegularityTypeUUID: a.RegularityTypeUUID,
			DriverUUID:         a.DriverUUID,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create trip UUID:%s", a.UUID),
			Run: func(db *gorm.DB) error {
				_, errDB := createTrip(db, trip)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createTrip will create fake trip and insert into DB.
func createTrip(db *gorm.DB, trip *entity.Trip) (*entity.Trip, error) {
	var tripExists entity.Trip
	err := db.Where("route_uuid = ? and vehicle_uuid = ? and departure_time = ?", trip.RouteUUID, trip.VehicleUUID, trip.DepartureTime).
		Take(&tripExists).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(trip).Error
			if err != nil {
				return trip, err
			}
			return trip, err
		}
		return trip, err
	}
	return trip, err
}
