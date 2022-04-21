package routers

import (
	VehicleV1Point00 "cargo-rest-api/interfaces/handler/v1.0/vehicle"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func vehicleRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	vehicleV1 := VehicleV1Point00.NewVehicles(r.dbService.Vehicle)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/vehicles", guard.Authenticate(), guard.Authorize("vehicle_read"), vehicleV1.GetVehicles)
	v1.POST("/vehicles", guard.Authenticate(), guard.Authorize("vehicle_create"), vehicleV1.SaveVehicle)
	v1.GET("/vehicles/:uuid", guard.Authenticate(), guard.Authorize("vehicle_detail"), vehicleV1.GetVehicle)
	v1.PUT("/vehicles/:uuid", guard.Authenticate(), guard.Authorize("vehicle_update"), vehicleV1.UpdateVehicle)
	v1.DELETE("/vehicles/:uuid", guard.Authenticate(), guard.Authorize("vehicle_delete"), vehicleV1.DeleteVehicle)
}
