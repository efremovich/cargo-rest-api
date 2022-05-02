package routers

import (
	RouteV1Point00 "cargo-rest-api/interfaces/handler/v1.0/route"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func routeRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	RouteV1 := RouteV1Point00.NewRoutes(r.dbService.Route)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/routes", guard.Authenticate(), RouteV1.GetRoutes)
	v1.POST("/route", guard.Authenticate(), RouteV1.SaveRoute)
	v1.GET("/route/:uuid", guard.Authenticate(), RouteV1.GetRoute)
	v1.PUT("/route/:uuid", guard.Authenticate(), RouteV1.UpdateRoute)
	v1.DELETE("/route/:uuid", guard.Authenticate(), RouteV1.DeleteRoute)

	v1.POST("/route/price_add", guard.Authenticate(), RouteV1.AddRoutePrice)
	v1.POST("/route/price_del", guard.Authenticate(), RouteV1.DeleteRoutePrice)
}
