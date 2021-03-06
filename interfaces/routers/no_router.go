package routers

import (
	"cargo-rest-api/infrastructure/message/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func noRoutes(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		err := exception.ErrorTextNotFound
		_ = c.AbortWithError(http.StatusNotFound, err)
	})
}
