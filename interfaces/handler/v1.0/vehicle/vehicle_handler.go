package vehiclev1point00

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

// Vehicles is a struct defines the dependencies that will be used.
type Vehicles struct {
	us application.VehicleAppInterface
}

// NewCountreis is constructor will initialize vehicle handler.
func NewVehicles(us application.VehicleAppInterface) *Vehicles {
	return &Vehicles{
		us: us,
	}
}

// @Summary Create a new vehicle
// @Description Create a new vehicle.
// @Tags vehicles
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param vehicle body entity.DetailVehicles true "Vehicle vehicle"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/vehicles [post]
// SaveVehicle is a function vehicle to handle create a new vehicle.
func (s *Vehicles) SaveVehicles(c *gin.Context) {
	var vehicleEntity entity.Vehicle
	if err := c.ShouldBindJSON(&vehicleEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := vehicleEntity.ValidateSaveVehicle()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newVehicle, errDesc, errException := s.us.SaveVehicle(&vehicleEntity)
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
	response.NewSuccess(c, newVehicle.DetailVehicle(), success.VehicleSuccessfullyCreateVehicle).JSON()
}

// @Summary Update vehicle
// @Description Update an existing vehicle.
// @Tags vehicles
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Vehicle UUID"
// @Param name formData string true "Vehicle name"
// @Param region formData string true "Vehicle region"
// @Param latitude formData string true "Vehicle latitude"
// @Param longitude formData string true "Vehicle longitude"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/vehicles/uuid [put]
// UpdateVehicle is a function uses to handle update vehicle by UUID.
func (s *Vehicles) UpdateVehicles(c *gin.Context) {
	var vehicleEntity entity.Vehicle
	if err := c.ShouldBindUri(&vehicleEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&vehicleEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := c.Param("uuid")
	_, err := s.us.GetVehicle(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextVehicleNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextVehicleNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedVehicle, errDesc, errException := s.us.UpdateVehicle(UUID, &vehicleEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextVehicleNotFound) {
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
	response.NewSuccess(c, updatedVehicle.DetailVehicle(), success.VehicleSuccessfullyUpdateVehicle).JSON()
}

// @Summary Delete vehicle
// @Description Delete an existing vehicle.
// @Tags vehicles
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Vehicle UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/vehicles/{uuid} [delete]
// DeleteVehicle is a function uses to handle delete vehicle by UUID.
func (s *Vehicles) DeleteVehicle(c *gin.Context) {
	var vehicleEntity entity.Vehicle
	if err := c.ShouldBindUri(&vehicleEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeleteVehicle(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextVehicleNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextVehicleNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.VehicleSuccessfullyDeleteVehicle).JSON()
}

// @Summary Get vehicles
// @Description Get list of existing vehicles.
// @Tags vehicles
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
// @Router /api/v1/external/vehicles [get]
// GetVehicles is a function uses to handle get vehicle list.
func (s *Vehicles) GetVehicles(c *gin.Context) {
	var vehicle entity.Vehicle
	var vehicles entity.Vehicles
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(vehicle.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	vehicles, meta, err := s.us.GetVehicles(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, vehicles.DetailVehicles(), success.VehicleSuccessfullyGetVehicleList).WithMeta(meta).JSON()
}

// @Summary Get vehicle
// @Description Get detail of existing vehicle.
// @Tags vehicles
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Vehicle UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/vehicles/{uuid} [get]
// GetVehicle is a function uses to handle get vehicle detail by UUID.
func (s *Vehicles) GetVehicle(c *gin.Context) {
	var vehicleEntity entity.Vehicle
	if err := c.ShouldBindUri(&vehicleEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	vehicle, err := s.us.GetVehicle(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextVehicleNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextVehicleNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, vehicle.DetailVehicle(), success.VehicleSuccessfullyGetVehicleDetail).JSON()
}
