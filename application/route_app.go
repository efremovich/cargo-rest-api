package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type routeApp struct {
	tr repository.RouteRepository
}

// routeApp implement the RouteAppInterface.
var _ RouteAppInterface = &routeApp{}

// RouteAppInterface is an interface.
type RouteAppInterface interface {
	SaveRoute(*entity.Route) (*entity.Route, map[string]string, error)
	UpdateRoute(
		UUID string,
		route *entity.Route,
	) (*entity.Route, map[string]string, error)
	DeleteRoute(UUID string) error
	GetRoutes(p *repository.Parameters) ([]*entity.Route, *repository.Meta, error)
	GetRoute(UUID string) (*entity.Route, error)

	AddRoutePrice(route *entity.Route) (*entity.Route, map[string]string, error)
	DeleteRoutePrice(route *entity.Route) (*entity.Route, map[string]string, error)
}

func (t routeApp) SaveRoute(
	route *entity.Route,
) (*entity.Route, map[string]string, error) {
	return t.tr.SaveRoute(route)
}

func (t routeApp) UpdateRoute(
	UUID string,
	route *entity.Route,
) (*entity.Route, map[string]string, error) {
	return t.tr.UpdateRoute(UUID, route)
}

func (t routeApp) DeleteRoute(UUID string) error {
	return t.tr.DeleteRoute(UUID)
}

func (t routeApp) GetRoutes(
	p *repository.Parameters,
) ([]*entity.Route, *repository.Meta, error) {
	return t.tr.GetRoutes(p)
}

func (t routeApp) GetRoute(UUID string) (*entity.Route, error) {
	return t.tr.GetRoute(UUID)
}

func (t routeApp) AddRoutePrice(route *entity.Route) (*entity.Route, map[string]string, error) {
	return t.tr.AddRoutePrice(route)
}

func (t routeApp) DeleteRoutePrice(route *entity.Route) (*entity.Route, map[string]string, error) {
	return t.tr.DeleteRoutePrice(route)
}
