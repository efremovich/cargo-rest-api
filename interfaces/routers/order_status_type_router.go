package routers

import (
	OrderStatusTypeV1Point00 "cargo-rest-api/interfaces/handler/v1.0/order_status_type"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func orderStatusTypeRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	OrderStatusTypeV1 := OrderStatusTypeV1Point00.NewOrderStatusTypes(r.dbService.OrderStatusType)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/orderStatusTypes", guard.Authenticate(), OrderStatusTypeV1.GetOrderStatusTypes)
	v1.POST("/orderStatusTypes", guard.Authenticate(), OrderStatusTypeV1.SaveOrderStatusType)
	v1.GET("/orderStatusTypes/:uuid", guard.Authenticate(), OrderStatusTypeV1.GetOrderStatusType)
	v1.PUT("/orderStatusTypes/:uuid", guard.Authenticate(), OrderStatusTypeV1.UpdateOrderStatusType)
	v1.DELETE("/orderStatusTypes/:uuid", guard.Authenticate(), OrderStatusTypeV1.DeleteOrderStatusType)
}
