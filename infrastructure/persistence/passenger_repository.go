package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// PassengerRepo is a struct to store db connection.
type PassengerRepo struct {
	db *gorm.DB
}

// NewPassengerRepository will initialize Passenger repository.
func NewPassengerRepository(db *gorm.DB) *PassengerRepo {
	return &PassengerRepo{db}
}

// PassengerRepo implements the repository.PassengerRepository interface.
var _ repository.PassengerRepository = &PassengerRepo{}

// SavePassenger will create a new Passenger.
func (r PassengerRepo) SavePassenger(Passenger *entity.Passenger) (*entity.Passenger, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&Passenger).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Passenger, nil, nil
}

func (r PassengerRepo) UpdatePassenger(
	uuid string,
	Passenger *entity.Passenger,
) (*entity.Passenger, map[string]string, error) {
	errDesc := map[string]string{}
	PassengerData := &entity.Passenger{
		FirstName:         Passenger.FirstName,
		LastName:          Passenger.LastName,
		Patronomic:        Passenger.Patronomic,
		BirthDay:          Passenger.BirthDay,
		DocumentSeries:    Passenger.DocumentSeries,
		DocumentNumber:    Passenger.DocumentNumber,
		UserUUID:          Passenger.UserUUID,
		PassengerTypeUUID: Passenger.PassengerTypeUUID,
	}

	err := r.db.First(&Passenger, "uuid = ?", uuid).Updates(PassengerData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextPassengerInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextPassengerNotFound
		}
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			errDesc["type"] = exception.ErrorTextUserEmailAlreadyTaken.Error()
			return nil, errDesc, exception.ErrorTextUnprocessableEntity
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Passenger, nil, nil
}

func (r PassengerRepo) DeletePassenger(uuid string) error {
	var Passenger entity.Passenger
	err := r.db.Where("uuid = ?", uuid).Take(&Passenger).Delete(&Passenger).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextPassengerNotFound
		}
		return err
	}
	return nil
}

func (r PassengerRepo) GetPassenger(uuid string) (*entity.Passenger, error) {
	var Passenger entity.Passenger
	err := r.db.Preload("DocumentType").Where("uuid = ?", uuid).Take(&Passenger).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextPassengerNotFound
		}
	}
	return &Passenger, nil
}

func (r PassengerRepo) GetPassengers(p *repository.Parameters) ([]*entity.Passenger, *repository.Meta, error) {
	var total int64
	var Passengers []*entity.Passenger
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&Passengers).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&Passengers).Error
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
	return Passengers, meta, nil
}
