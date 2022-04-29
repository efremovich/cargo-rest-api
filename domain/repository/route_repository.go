package repository

import (
	"cargo-rest-api/domain/entity"
)

// RouteRepository is an interface.
type RouteRepository interface {
	SaveRoute(route *entity.Route) (*entity.Route, map[string]string, error)
	UpdateRoute(UUID string, tour *entity.Route) (*entity.Route, map[string]string, error)
	DeleteRoute(UUID string) error
	GetRoute(UUID string) (*entity.Route, error)
	GetRoutes(parameters *Parameters) ([]*entity.Route, *Meta, error)

	AddRoutePrice(route *entity.Route) (*entity.Route, map[string]string, error)
	DeleteRoutePrice(route *entity.Route) (*entity.Route, map[string]string, error)
}
