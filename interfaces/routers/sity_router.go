package routers

import (
	SityV1Point00 "cargo-rest-api/interfaces/handler/v1.0/sity"

	"github.com/gin-gonic/gin"
)

func sityRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	SityV1 := SityV1Point00.NewSities(r.dbService.Sity)

	v1 := e.Group("/api/v1/external")

	v1.GET("/sities", SityV1.GetSities)
	v1.POST("/sities", SityV1.SaveSities)
	v1.GET("/coutries/:uuid", SityV1.GetSity)
	v1.PUT("/sities/:uuid", SityV1.UpdateSities)
	v1.DELETE("/sities/:uuid", SityV1.DeleteSity)
}
