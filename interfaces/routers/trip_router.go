package routers

import (
	TripV1Point00 "cargo-rest-api/interfaces/handler/v1.0/trip"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func tripRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	TripV1 := TripV1Point00.NewTrips(r.dbService.Trip)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/trips", guard.Authenticate(), TripV1.GetTrips)
	v1.POST("/trip", guard.Authenticate(), TripV1.SaveTrip)
	v1.GET("/trip/:uuid", guard.Authenticate(), TripV1.GetTrip)
	v1.PUT("/trip/:uuid", guard.Authenticate(), TripV1.UpdateTrip)
	v1.DELETE("/trip/:uuid", guard.Authenticate(), TripV1.DeleteTrip)
}
