package routers

import (
	SityV1Point00 "cargo-rest-api/interfaces/handler/v1.0/sity"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func sityRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	sityV1 := SityV1Point00.NewSities(r.dbService.Sity)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/sities", guard.Authenticate(), guard.Authorize("sity_read"), sityV1.GetSities)
	v1.POST("/sity", guard.Authenticate(), guard.Authorize("sity_create"), sityV1.SaveSity)
	v1.GET("/sity/:uuid", guard.Authenticate(), guard.Authorize("sity_detail"), sityV1.GetSity)
	v1.PUT("/sity/:uuid", guard.Authenticate(), guard.Authorize("sity_update"), sityV1.UpdateSity)
	v1.DELETE("/sity/:uuid", guard.Authenticate(), guard.Authorize("sity_delete"), sityV1.DeleteSity)
}
