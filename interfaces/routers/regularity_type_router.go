package routers

import (
	RegularityTypeV1Point00 "cargo-rest-api/interfaces/handler/v1.0/regularity_type"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func regularityTypeRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	RegularityTypeV1 := RegularityTypeV1Point00.NewRegularityTypes(r.dbService.RegularityType)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/regularityTypes", guard.Authenticate(), RegularityTypeV1.GetRegularityTypes)
	v1.POST("/regularityTypes", guard.Authenticate(), RegularityTypeV1.SaveRegularityType)
	v1.GET("/regularityTypes/:uuid", guard.Authenticate(), RegularityTypeV1.GetRegularityType)
	v1.PUT("/regularityTypes/:uuid", guard.Authenticate(), RegularityTypeV1.UpdateRegularityType)
	v1.DELETE("/regularityTypes/:uuid", guard.Authenticate(), RegularityTypeV1.DeleteRegularityType)
}
