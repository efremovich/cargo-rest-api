package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// DriverRepo is a struct to store db connection.
type DriverRepo struct {
	db *gorm.DB
}

// NewDriverRepository will initialize Driver repository.
func NewDriverRepository(db *gorm.DB) *DriverRepo {
	return &DriverRepo{db}
}

// DriverRepo implements the repository.driverRepository interface.
var _ repository.DriverRepository = &DriverRepo{}

// SaveDriver will create a new driver.
func (r DriverRepo) SaveDriver(
	Driver *entity.Driver,
) (*entity.Driver, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Create(&Driver).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Driver, nil, nil
}

func (r DriverRepo) UpdateDriver(
	uuid string,
	driver *entity.Driver,
) (*entity.Driver, map[string]string, error) {
	errDesc := map[string]string{}
	dirverData := &entity.Driver{
		Name:     driver.Name,
		UserUUID: driver.UserUUID,
	}

	err := r.db.First(&driver, "uuid = ?", uuid).Updates(dirverData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextDriverInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextDriverNotFound
		}
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			errDesc["type"] = exception.ErrorTextUserEmailAlreadyTaken.Error()
			return nil, errDesc, exception.ErrorTextUnprocessableEntity
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return driver, nil, nil
}

func (r DriverRepo) DeleteDriver(uuid string) error {
	var driver entity.Driver
	err := r.db.Where("uuid = ?", uuid).Take(&driver).Delete(&driver).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextDriverNotFound
		}
		return err
	}
	return nil
}

func (r DriverRepo) GetDriver(uuid string) (*entity.Driver, error) {
	var driver entity.Driver
	err := r.db.Preload("Vehicles").Where("uuid = ?", uuid).Take(&driver).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextDriverNotFound
		}
	}
	return &driver, nil
}

func (r DriverRepo) GetDrivers(p *repository.Parameters) ([]*entity.Driver, *repository.Meta, error) {
	var total int64
	var drivers []*entity.Driver
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&drivers).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&drivers).Error
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
	return drivers, meta, nil
}
