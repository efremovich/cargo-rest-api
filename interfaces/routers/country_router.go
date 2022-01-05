package routers

import (
	CountryV1Point00 "cargo-rest-api/interfaces/handler/v1.0/country"

	"github.com/gin-gonic/gin"
)

func countryRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	CountryV1 := CountryV1Point00.NewCountries(r.dbService.Country)

	v1 := e.Group("/api/v1/external")

	v1.GET("/countries", CountryV1.GetCountries)
	v1.POST("/countries", CountryV1.SaveCountries)
	v1.GET("/coutries/:uuid", CountryV1.GetCountry)
	v1.PUT("/countries/:uuid", CountryV1.UpdateCountries)
	v1.DELETE("/countries/:uuid", CountryV1.DeleteCountry)
}
