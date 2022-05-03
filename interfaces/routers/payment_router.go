package routers

import (
	PaymentV1Point00 "cargo-rest-api/interfaces/handler/v1.0/payment"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func paymentRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	PaymentV1 := PaymentV1Point00.NewPayments(r.dbService.Payment)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/payments", guard.Authenticate(), PaymentV1.GetPayments)
	v1.POST("/payment", guard.Authenticate(), PaymentV1.SavePayment)
	v1.GET("/payment/:uuid", guard.Authenticate(), PaymentV1.GetPayment)
	v1.PUT("/payment/:uuid", guard.Authenticate(), PaymentV1.UpdatePayment)
	v1.DELETE("/payment/:uuid", guard.Authenticate(), PaymentV1.DeletePayment)

	v1.POST("/payment/price_add", guard.Authenticate(), PaymentV1.AddOrderPayment)
	v1.POST("/payment/price_del", guard.Authenticate(), PaymentV1.DeleteOrderPayment)
}
