package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// RouteAppInterface is a mock of application.RouteAppInterface.
type RouteAppInterface struct {
	SaveRouteFn   func(*entity.Route) (*entity.Route, map[string]string, error)
	UpdateRouteFn func(string, *entity.Route) (*entity.Route, map[string]string, error)
	DeleteRouteFn func(UUID string) error
	GetRoutesFn   func(params *repository.Parameters) ([]*entity.Route, *repository.Meta, error)
	GetRouteFn    func(UUID string) (*entity.Route, error)

	AddRoutePriceFn    func(*entity.Route) (*entity.Route, map[string]string, error)
	DeleteRoutePriceFn func(*entity.Route) (*entity.Route, map[string]string, error)
}

// SaveRoute calls the SaveRouteFn.
func (u *RouteAppInterface) SaveRoute(route *entity.Route) (*entity.Route, map[string]string, error) {
	return u.SaveRouteFn(route)
}

// UpdateRoute calls the UpdateRouteFn.
func (u *RouteAppInterface) UpdateRoute(uuid string, route *entity.Route) (*entity.Route, map[string]string, error) {
	return u.UpdateRouteFn(uuid, route)
}

// DeleteRoute calls the DeleteRouteFn.
func (u *RouteAppInterface) DeleteRoute(uuid string) error {
	return u.DeleteRouteFn(uuid)
}

// GetRoutes calls the GetRoutesFn.
func (u *RouteAppInterface) GetRoutes(
	params *repository.Parameters,
) ([]*entity.Route, *repository.Meta, error) {
	return u.GetRoutesFn(params)
}

// GetRoute calls the GetRouteFn.
func (u *RouteAppInterface) GetRoute(uuid string) (*entity.Route, error) {
	return u.GetRouteFn(uuid)
}

// AddRoutePrice calls the AddRoutePriceFn.
func (u *RouteAppInterface) AddRoutePrice(route *entity.Route) (*entity.Route, map[string]string, error) {
	return u.AddRoutePriceFn(route)
}

// DeleteRoutePrice calls the DeleteRoutePriceFn.
func (u *RouteAppInterface) DeleteRoutePrice(route *entity.Route) (*entity.Route, map[string]string, error) {
	return u.DeleteRoutePriceFn(route)
}
