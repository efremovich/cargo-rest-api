package routers

import (
	OrderV1Point00 "cargo-rest-api/interfaces/handler/v1.0/order"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func orderRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	OrderV1 := OrderV1Point00.NewOrders(r.dbService.Order)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/orders", guard.Authenticate(), OrderV1.GetOrders)
	v1.POST("/order", guard.Authenticate(), OrderV1.SaveOrder)
	v1.GET("/order/:uuid", guard.Authenticate(), OrderV1.GetOrder)
	v1.PUT("/order/:uuid", guard.Authenticate(), OrderV1.UpdateOrder)
	v1.DELETE("/order/:uuid", guard.Authenticate(), OrderV1.DeleteOrder)
}
