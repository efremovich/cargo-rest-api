package routers

import (
	"cargo-rest-api/interfaces/handler"
	"cargo-rest-api/interfaces/middleware"
	"cargo-rest-api/interfaces/service"

	"github.com/gin-gonic/gin"
)

func authRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	authenticate := handler.NewAuthenticate(
		service.NewAuthService(r.dbService.User, r.dbService.UserForgotPassword),
		r.redisService.Auth,
		rg.authToken,
		r.notificationService.Notification)

	v1 := e.Group("/api/v1/external")

	v1.GET("/profile", middleware.Auth(rg.authGateway), authenticate.Profile)
	v1.POST("/login", authenticate.Login)
	v1.POST("/logout", middleware.Auth(rg.authGateway), authenticate.Logout)
	v1.POST("/refresh", authenticate.Refresh)
	v1.POST("/password/forgot", authenticate.ForgotPassword)
	v1.POST("/password/reset/:token", authenticate.ResetPassword)
}
