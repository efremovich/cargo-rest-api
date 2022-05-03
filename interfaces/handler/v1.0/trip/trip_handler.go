package tripv1point00

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

// Trips is a struct defines the dependencies that will be used.
type Trips struct {
	us application.TripAppInterface
}

// NewCountreis is constructor will initialize trip handler.
func NewTrips(us application.TripAppInterface) *Trips {
	return &Trips{
		us: us,
	}
}

// @Summary Create a new trip
// @Description Create a new trip.
// @Tags trips
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param trip body entity.DetailTrip true "Trip trip"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Tripr /api/v1/external/trip [post]
// SaveTrip is a function trip to handle create a new trip.
func (s *Trips) SaveTrip(c *gin.Context) {
	var tripEntity entity.Trip
	if err := c.ShouldBindJSON(&tripEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := tripEntity.ValidateSaveTrip()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newTrip, errDesc, errException := s.us.SaveTrip(&tripEntity)
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
	response.NewSuccess(c, newTrip.DetailTrip(), success.TripSuccessfullyCreateTrip).
		JSON()
}

// @Summary Update trip
// @Description Update an existing trip.
// @Tags trips
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param trip body entity.DetailTrip true "Trip trip"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Tripr /api/v1/external/trip/uuid [put]
// UpdateTrip is a function uses to handle update trip by UUID.
func (s *Trips) UpdateTrip(c *gin.Context) {
	var tripEntity entity.Trip
	if err := c.ShouldBindUri(&tripEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&tripEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := tripEntity.UUID
	_, err := s.us.GetTrip(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextTripNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextTripNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedTrip, errDesc, errException := s.us.UpdateTrip(UUID, &tripEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextTripNotFound) {
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
	response.NewSuccess(c, updatedTrip.DetailTrip(), success.TripSuccessfullyUpdateTrip).
		JSON()
}

// @Summary Delete trip
// @Description Delete an existing trip.
// @Tags trips
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Trip UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Tripr /api/v1/external/trip/{uuid} [delete]
// DeleteTrip is a function uses to handle delete trip by UUID.
func (s *Trips) DeleteTrip(c *gin.Context) {
	var tripEntity entity.Trip
	if err := c.ShouldBindUri(&tripEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeleteTrip(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextTripNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextTripNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.TripSuccessfullyDeleteTrip).JSON()
}

// @Summary Get trips
// @Description Get list of existing trips.
// @Tags trips
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
// @Tripr /api/v1/external/trips [get]
// GetTrips is a function uses to handle get trip list.
func (s *Trips) GetTrips(c *gin.Context) {
	var trip entity.Trip
	var trips entity.Trips
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(trip.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	trips, meta, err := s.us.GetTrips(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, trips.DetailTrips(), success.TripSuccessfullyGetTripList).
		WithMeta(meta).
		JSON()
}

// @Summary Get trip
// @Description Get detail of existing trip.
// @Tags trips
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Trip UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Tripr /api/v1/external/trip/{uuid} [get]
// GetTrip is a function uses to handle get trip detail by UUID.
func (s *Trips) GetTrip(c *gin.Context) {
	var tripEntity entity.Trip
	if err := c.ShouldBindUri(&tripEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	trip, err := s.us.GetTrip(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextTripNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextTripNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, trip.DetailTrip(), success.TripSuccessfullyGetTripDetail).
		JSON()
}
