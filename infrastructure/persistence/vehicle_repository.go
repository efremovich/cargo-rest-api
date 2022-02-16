package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// VehicleRepo is a struct to store db connection.
type VehicleRepo struct {
	db *gorm.DB
}

// NewVehicleRepository will initialize Vehicle repository.
func NewVehicleRepository(db *gorm.DB) *VehicleRepo {
	return &VehicleRepo{db}
}

// VehicleRepo implements the repository.vehicleRepository interface.
var _ repository.VehicleRepository = &VehicleRepo{}

// SaveVehicle will create a new vehicle.
func (r VehicleRepo) SaveVehicle(Vehicle *entity.Vehicle) (*entity.Vehicle, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&Vehicle).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Vehicle, nil, nil
}

func (r VehicleRepo) UpdateVehicle(uuid string, vehicle *entity.Vehicle) (*entity.Vehicle, map[string]string, error) {
	errDesc := map[string]string{}
	vehicleData := &entity.Vehicle{
		Model:         vehicle.Model,
		RegCode:       vehicle.RegCode,
		NumberOfSeats: vehicle.NumberOfSeats,
		Class:         vehicle.Class,
	}

	err := r.db.First(&vehicle, "uuid = ?", uuid).Updates(vehicleData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextVehicleInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextVehicleNotFound
		}
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			errDesc["reg_code"] = exception.ErrorTextUserEmailAlreadyTaken.Error()
			return nil, errDesc, exception.ErrorTextUnprocessableEntity
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return vehicle, nil, nil
}

func (r VehicleRepo) DeleteVehicle(uuid string) error {
	var vehicle entity.Vehicle
	err := r.db.Where("uuid = ?", uuid).Take(&vehicle).Delete(&vehicle).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextVehicleNotFound
		}
		return err
	}
	return nil
}

func (r VehicleRepo) GetVehicle(uuid string) (*entity.Vehicle, error) {
	var vehicle entity.Vehicle
	err := r.db.Where("uuid = ?", uuid).Take(&vehicle).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextVehicleNotFound
		}
	}
	return &vehicle, nil
}

func (r VehicleRepo) GetVehicles(p *repository.Parameters) ([]*entity.Vehicle, *repository.Meta, error) {
	var total int64
	var vehicles []*entity.Vehicle
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&vehicles).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&vehicles).Error
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
	return vehicles, meta, nil
}
