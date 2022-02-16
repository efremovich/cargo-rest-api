package routers

import (
	VehicleV1Point00 "cargo-rest-api/interfaces/handler/v1.0/vehicle"

	"github.com/gin-gonic/gin"
)

func vehicleRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	VehicleV1 := VehicleV1Point00.NewVehicles(r.dbService.Vehicle)

	v1 := e.Group("/api/v1/external")

	v1.GET("/vehicles", VehicleV1.GetVehicles)
	v1.POST("/vehicles", VehicleV1.SaveVehicle)
	v1.GET("/vehicles/:uuid", VehicleV1.GetVehicle)
	v1.PUT("/vehicles/:uuid", VehicleV1.UpdateVehicle)
	v1.DELETE("/vehicles/:uuid", VehicleV1.DeleteVehicle)
}
