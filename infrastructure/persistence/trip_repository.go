package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"

	"gorm.io/gorm"
)

// TripRepo is a struct to store db connection.
type TripRepo struct {
	db *gorm.DB
}

// NewTripRepository will initialize Trip repository.
func NewTripRepository(db *gorm.DB) *TripRepo {
	return &TripRepo{db}
}

// TripRepo implements the repository.tripRepository interface.
var _ repository.TripRepository = &TripRepo{}

// SaveTrip will create a new trip.
func (r TripRepo) SaveTrip(Trip *entity.Trip) (*entity.Trip, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&Trip).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Trip, nil, nil
}

func (r TripRepo) UpdateTrip(uuid string, trip *entity.Trip) (*entity.Trip, map[string]string, error) {
	errDesc := map[string]string{}
	dirverData := &entity.Trip{
		RouteUUID:          trip.RouteUUID,
		VehicleUUID:        trip.VehicleUUID,
		DepartureTime:      trip.DepartureTime,
		ArravialTive:       trip.ArravialTive,
		RegularityTypeUUID: trip.RegularityTypeUUID,
		DriverUUID:         trip.DriverUUID,
	}

	err := r.db.First(&trip, "uuid = ?", uuid).Updates(dirverData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextTripInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextTripNotFound
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return trip, nil, nil
}

func (r TripRepo) DeleteTrip(uuid string) error {
	var trip entity.Trip
	err := r.db.Where("uuid = ?", uuid).Take(&trip).Delete(&trip).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextTripNotFound
		}
		return err
	}
	return nil
}

func (r TripRepo) GetTrip(uuid string) (*entity.Trip, error) {
	var trip entity.Trip
	err := r.db.Preload("Route").
		Preload("Vehicle").
		Preload("RegularityType").
		Preload("Driver").
		Where("uuid = ?", uuid).
		Take(&trip).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextTripNotFound
		}
	}
	return &trip, nil
}

func (r TripRepo) GetTrips(p *repository.Parameters) ([]*entity.Trip, *repository.Meta, error) {
	var total int64
	var trips []*entity.Trip
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&trips).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&trips).Error
	if errTotal != nil {
		return nil, nil, errTotal
	}
	if errList != nil {
		return nil, nil, errList
	}
	if errors.Is(errList, gorm.ErrRecordNotFound) {
		return nil, nil, errList
	}
	meta := repository.NewMeta(p, total)
	return trips, meta, nil
}
