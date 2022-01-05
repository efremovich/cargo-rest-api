package countryv1point00

import (
	"cargo-rest-api/application"
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"cargo-rest-api/infrastructure/message/success"
	"cargo-rest-api/pkg/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Countries is a struct defines the dependencies that will be used.
type Countries struct {
	us application.CountryAppInterface
}

// NewCountreis is constructor will initialize country handler.
func NewCountries(us application.CountryAppInterface) *Countries {
	return &Countries{
		us: us,
	}
}

// @Summary Create a new country
// @Description Create a new country.
// @Tags countries
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param name formData string true "Country name"
// @Param latitude formData string true "Country latitude"
// @Param longitude formData string true "Country longitude"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/countries [post]
// SaveCountry is a function country to handle create a new country.
func (s *Countries) SaveCountries(c *gin.Context) {
	var countryEntity entity.Country
	if err := c.ShouldBindJSON(&countryEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	newCountry, errDesc, errException := s.us.SaveCountry(&countryEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextUnprocessableEntity) {
			_ = c.AbortWithError(http.StatusUnprocessableEntity, errException)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
	response.NewSuccess(c, newCountry.DetailCountry(), success.CountrySuccessfullyCreateCountry).JSON()
}

// @Summary Update country
// @Description Update an existing country.
// @Tags countries
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Country UUID"
// @Param name formData string true "Country name"
// @Param latitude formData string true "Country latitude"
// @Param longitude formData string true "Country longitude"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/countries/uuid [put]
// UpdateCountry is a function uses to handle update country by UUID.
func (s *Countries) UpdateCountries(c *gin.Context) {
	var countryEntity entity.Country
	if err := c.ShouldBindUri(&countryEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&countryEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := c.Param("uuid")
	_, err := s.us.GetCountry(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextCountryNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextCountryNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedCountry, errDesc, errException := s.us.UpdateCountry(UUID, &countryEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextCountryNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, errException)
			return
		}
		if errors.Is(errException, exception.ErrorTextUnprocessableEntity) {
			_ = c.AbortWithError(http.StatusUnprocessableEntity, errException)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextInternalServerError)
		return
	}
	c.Status(http.StatusOK)
	response.NewSuccess(c, updatedCountry.DetailCountry(), success.CountrySuccessfullyUpdateCountry).JSON()
}

// @Summary Delete country
// @Description Delete an existing country.
// @Tags countries
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Country UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/countries/{uuid} [delete]
// DeleteCountry is a function uses to handle delete country by UUID.
func (s *Countries) DeleteCountry(c *gin.Context) {
	var countryEntity entity.Country
	if err := c.ShouldBindUri(&countryEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeleteCountry(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextCountryNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextCountryNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.CountrySuccessfullyDeleteCountry).JSON()
}

// @Summary Get countries
// @Description Get list of existing countries.
// @Tags countries
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/countries [get]
// GetCountries is a function uses to handle get country list.
func (s *Countries) GetCountries(c *gin.Context) {
	var country entity.Country
	var countries entity.Countries
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(country.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	countries, meta, err := s.us.GetCountries(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, countries.DetailCountries(), success.CountrySuccessfullyGetCountryList).WithMeta(meta).JSON()
}

// @Summary Get country
// @Description Get detail of existing country.
// @Tags countries
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, id) default(id)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Country UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/countries/{uuid} [get]
// GetCountry is a function uses to handle get country detail by UUID.
func (s *Countries) GetCountry(c *gin.Context) {
	var countryEntity entity.Country
	if err := c.ShouldBindUri(&countryEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	country, err := s.us.GetCountry(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextCountryNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextCountryNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, country.DetailCountry(), success.CountrySuccessfullyGetCountryDetail).JSON()
}
