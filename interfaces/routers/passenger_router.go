package routers

import (
	PassengerV1Point00 "cargo-rest-api/interfaces/handler/v1.0/passenger"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func passengerRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	PassengerV1 := PassengerV1Point00.NewPassengers(r.dbService.Passenger)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/passengers", guard.Authenticate(), PassengerV1.GetPassengers)
	v1.POST("/passenger", guard.Authenticate(), PassengerV1.SavePassenger)
	v1.GET("/passenger/:uuid", guard.Authenticate(), PassengerV1.GetPassenger)
	v1.PUT("/passenger/:uuid", guard.Authenticate(), PassengerV1.UpdatePassenger)
	v1.DELETE("/passenger/:uuid", guard.Authenticate(), PassengerV1.DeletePassenger)
}
