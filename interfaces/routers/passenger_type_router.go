package routers

import (
	PassengerTypeV1Point00 "cargo-rest-api/interfaces/handler/v1.0/passenger_type"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func passengerTypeRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	PassengerTypeV1 := PassengerTypeV1Point00.NewPassengerTypes(r.dbService.PassengerType)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/passengerTypes", guard.Authenticate(), PassengerTypeV1.GetPassengerTypes)
	v1.POST("/passengerTypes", guard.Authenticate(), PassengerTypeV1.SavePassengerType)
	v1.GET("/passengerTypes/:uuid", guard.Authenticate(), PassengerTypeV1.GetPassengerType)
	v1.PUT("/passengerTypes/:uuid", guard.Authenticate(), PassengerTypeV1.UpdatePassengerType)
	v1.DELETE("/passengerTypes/:uuid", guard.Authenticate(), PassengerTypeV1.DeletePassengerType)
}
