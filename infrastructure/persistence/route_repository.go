package persistence

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"errors"

	"gorm.io/gorm"
)

// RouteRepo is a struct to store db connection.
type RouteRepo struct {
	db *gorm.DB
}

// NewRouteRepository will initialize Route repository.
func NewRouteRepository(db *gorm.DB) *RouteRepo {
	return &RouteRepo{db}
}

// RouteRepo implements the repository.routeRepository interface.
var _ repository.RouteRepository = &RouteRepo{}

// SaveRoute will create a new route.
func (r RouteRepo) SaveRoute(Route *entity.Route) (*entity.Route, map[string]string, error) {
	errDesc := map[string]string{}
	err := r.db.Model(&Route).Association("Vehicles").Error
	err = r.db.Create(&Route).Error
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return Route, nil, nil
}

func (r RouteRepo) UpdateRoute(uuid string, route *entity.Route) (*entity.Route, map[string]string, error) {
	errDesc := map[string]string{}
	dirverData := &entity.Route{
		FromUUID:     route.FromUUID,
		ToUUID:       route.ToUUID,
		Distance:     route.Distance,
		DistanceTime: route.DistanceTime,
	}
	r.db.Model(route).Association("Vehicles")

	err := r.db.First(&route, "uuid = ?", uuid).Updates(dirverData).Error
	if err != nil {
		//If record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errDesc["uuid"] = exception.ErrorTextRouteInvalidUUID.Error()
			return nil, errDesc, exception.ErrorTextRouteNotFound
		}
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return route, nil, nil
}

func (r RouteRepo) DeleteRoute(uuid string) error {
	var route entity.Route
	err := r.db.Where("uuid = ?", uuid).Take(&route).Delete(&route).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrorTextRouteNotFound
		}
		return err
	}
	return nil
}

func (r RouteRepo) GetRoute(uuid string) (*entity.Route, error) {
	var route entity.Route
	err := r.db.Preload("Prices").Where("uuid = ?", uuid).Take(&route).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorTextRouteNotFound
		}
	}
	return &route, nil
}

func (r RouteRepo) GetRoutes(p *repository.Parameters) ([]*entity.Route, *repository.Meta, error) {
	var total int64
	var routes []*entity.Route
	errTotal := r.db.Where(p.QueryKey, p.QueryValue...).Find(&routes).Count(&total).Error
	errList := r.db.Where(p.QueryKey, p.QueryValue...).Limit(p.Limit).Offset(p.Offset).Find(&routes).Error
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
	return routes, meta, nil
}

// AddRouteVehicle implements repository.RouteRepository
func (r RouteRepo) AddRoutePrice(route *entity.Route) (*entity.Route, map[string]string, error) {
	errDesc := map[string]string{}

	err := r.db.Model(route).Association("Prices").Append(route.Prices)
	if err != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return route, nil, nil
}

// AddRouteVehicle implements repository.RouteRepository
func (r RouteRepo) DeleteRoutePrice(route *entity.Route) (*entity.Route, map[string]string, error) {
	errDesc := map[string]string{}

	errDelete := r.db.Model(route).Association("Prices").Delete(route.Prices)
	if errDelete != nil {
		return nil, errDesc, exception.ErrorTextAnErrorOccurred
	}
	return route, nil, nil
}
