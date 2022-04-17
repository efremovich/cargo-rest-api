package routers

import (
	PriceV1Point00 "cargo-rest-api/interfaces/handler/v1.0/price"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func priceRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	PriceV1 := PriceV1Point00.NewPrices(r.dbService.Price)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/prices", guard.Authenticate(), PriceV1.GetPrices)
	v1.POST("/price", guard.Authenticate(), PriceV1.SavePrice)
	v1.GET("/price/:uuid", guard.Authenticate(), PriceV1.GetPrice)
	v1.PUT("/price/:uuid", guard.Authenticate(), PriceV1.UpdatePrice)
	v1.DELETE("/price/:uuid", guard.Authenticate(), PriceV1.DeletePrice)
}
