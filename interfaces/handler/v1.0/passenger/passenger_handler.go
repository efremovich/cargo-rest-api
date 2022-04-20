package passengerv1point00

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

// Passengers is a struct defines the dependencies that will be used.
type Passengers struct {
	us application.PassengerAppInterface
}

// NewCountreis is constructor will initialize passenger handler.
func NewPassengers(us application.PassengerAppInterface) *Passengers {
	return &Passengers{
		us: us,
	}
}

// @Summary Create a new passenger
// @Description Create a new passenger.
// @Tags passengers
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param passenger body entity.DetailPassenger true "Passenger passenger"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/passenger [post]
// SavePassenger is a function passenger to handle create a new passenger.
func (s *Passengers) SavePassenger(c *gin.Context) {
	var passengerEntity entity.Passenger
	if err := c.ShouldBindJSON(&passengerEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := passengerEntity.ValidateSavePassenger()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newPassenger, errDesc, errException := s.us.SavePassenger(&passengerEntity)
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
	response.NewSuccess(c, newPassenger.DetailPassenger(), success.PassengerSuccessfullyCreatePassenger).JSON()
}

// @Summary Update passenger
// @Description Update an existing passenger.
// @Tags passengers
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Passenger UUID"
// @Param name formData string true "Passenger name"
// @Param region formData string true "Passenger region"
// @Param latitude formData string true "Passenger latitude"
// @Param longitude formData string true "Passenger longitude"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/passenger/uuid [put]
// UpdatePassenger is a function uses to handle update passenger by UUID.
func (s *Passengers) UpdatePassenger(c *gin.Context) {
	var passengerEntity entity.Passenger
	if err := c.ShouldBindUri(&passengerEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&passengerEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := c.Param("uuid")
	_, err := s.us.GetPassenger(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPassengerNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPassengerNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedPassenger, errDesc, errException := s.us.UpdatePassenger(UUID, &passengerEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextPassengerNotFound) {
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
	response.NewSuccess(c, updatedPassenger.DetailPassenger(), success.PassengerSuccessfullyUpdatePassenger).JSON()
}

// @Summary Delete passenger
// @Description Delete an existing passenger.
// @Tags passengers
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Passenger UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/passenger/{uuid} [delete]
// DeletePassenger is a function uses to handle delete passenger by UUID.
func (s *Passengers) DeletePassenger(c *gin.Context) {
	var passengerEntity entity.Passenger
	if err := c.ShouldBindUri(&passengerEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeletePassenger(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPassengerNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPassengerNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.PassengerSuccessfullyDeletePassenger).JSON()
}

// @Summary Get passengers
// @Description Get list of existing passengers.
// @Tags passengers
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/passengers [get]
// GetPassengers is a function uses to handle get passenger list.
func (s *Passengers) GetPassengers(c *gin.Context) {
	var passenger entity.Passenger
	var passengers entity.Passengers
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(passenger.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	passengers, meta, err := s.us.GetPassengers(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, passengers.DetailPassengers(), success.PassengerSuccessfullyGetPassengerList).
		WithMeta(meta).
		JSON()
}

// @Summary Get passenger
// @Description Get detail of existing passenger.
// @Tags passengers
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Passenger UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/passenger/{uuid} [get]
// GetPassenger is a function uses to handle get passenger detail by UUID.
func (s *Passengers) GetPassenger(c *gin.Context) {
	var passengerEntity entity.Passenger
	if err := c.ShouldBindUri(&passengerEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	passenger, err := s.us.GetPassenger(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPassengerNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPassengerNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, passenger.DetailPassenger(), success.PassengerSuccessfullyGetPassengerDetail).JSON()
}
