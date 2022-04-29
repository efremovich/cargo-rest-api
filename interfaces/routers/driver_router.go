package routers

import (
	DriverV1Point00 "cargo-rest-api/interfaces/handler/v1.0/driver"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func driverRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	DriverV1 := DriverV1Point00.NewDrivers(r.dbService.Driver)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/drivers", guard.Authenticate(), DriverV1.GetDrivers)
	v1.POST("/driver", guard.Authenticate(), DriverV1.SaveDriver)
	v1.GET("/driver/:uuid", guard.Authenticate(), DriverV1.GetDriver)
	v1.PUT("/driver/:uuid", guard.Authenticate(), DriverV1.UpdateDriver)
	v1.DELETE("/driver/:uuid", guard.Authenticate(), DriverV1.DeleteDriver)

	v1.POST("/driver/vehicle_add", guard.Authenticate(), DriverV1.AddDriverVehicle)
	v1.POST("/driver/vehicle_del", guard.Authenticate(), DriverV1.DeleteDriverVehicle)
}
