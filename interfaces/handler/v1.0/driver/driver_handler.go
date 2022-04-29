package driverv1point00

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

// Drivers is a struct defines the dependencies that will be used.
type Drivers struct {
	us application.DriverAppInterface
}

// NewCountreis is constructor will initialize driver handler.
func NewDrivers(us application.DriverAppInterface) *Drivers {
	return &Drivers{
		us: us,
	}
}

// @Summary Create a new driver
// @Description Create a new driver.
// @Tags drivers
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param driver body entity.DetailDriver true "Driver driver"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/driver [post]
// SaveDriver is a function driver to handle create a new driver.
func (s *Drivers) SaveDriver(c *gin.Context) {
	var driverEntity entity.Driver
	if err := c.ShouldBindJSON(&driverEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := driverEntity.ValidateSaveDriver()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newDriver, errDesc, errException := s.us.SaveDriver(&driverEntity)
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
	response.NewSuccess(c, newDriver.DetailDriver(), success.DriverSuccessfullyCreateDriver).
		JSON()
}

// @Summary Update driver
// @Description Update an existing driver.
// @Tags drivers
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param driver body entity.DetailDriver true "Driver driver"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/driver/uuid [put]
// UpdateDriver is a function uses to handle update driver by UUID.
func (s *Drivers) UpdateDriver(c *gin.Context) {
	var driverEntity entity.Driver
	if err := c.ShouldBindUri(&driverEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&driverEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := driverEntity.UUID
	_, err := s.us.GetDriver(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextDriverNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextDriverNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedDriver, errDesc, errException := s.us.UpdateDriver(UUID, &driverEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextDriverNotFound) {
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
	response.NewSuccess(c, updatedDriver.DetailDriver(), success.DriverSuccessfullyUpdateDriver).
		JSON()
}

// @Summary Delete driver
// @Description Delete an existing driver.
// @Tags drivers
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Driver UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/driver/{uuid} [delete]
// DeleteDriver is a function uses to handle delete driver by UUID.
func (s *Drivers) DeleteDriver(c *gin.Context) {
	var driverEntity entity.Driver
	if err := c.ShouldBindUri(&driverEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeleteDriver(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextDriverNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextDriverNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.DriverSuccessfullyDeleteDriver).JSON()
}

// @Summary Get drivers
// @Description Get list of existing drivers.
// @Tags drivers
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
// @Router /api/v1/external/drivers [get]
// GetDrivers is a function uses to handle get driver list.
func (s *Drivers) GetDrivers(c *gin.Context) {
	var driver entity.Driver
	var drivers entity.Drivers
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(driver.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	drivers, meta, err := s.us.GetDrivers(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, drivers.DetailDrivers(), success.DriverSuccessfullyGetDriverList).
		WithMeta(meta).
		JSON()
}

// @Summary Get driver
// @Description Get detail of existing driver.
// @Tags drivers
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Driver UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/driver/{uuid} [get]
// GetDriver is a function uses to handle get driver detail by UUID.
func (s *Drivers) GetDriver(c *gin.Context) {
	var driverEntity entity.Driver
	if err := c.ShouldBindUri(&driverEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	driver, err := s.us.GetDriver(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextDriverNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextDriverNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, driver.DetailDriver(), success.DriverSuccessfullyGetDriverDetail).
		JSON()
}

// @Summary Add vehicle to a driver
// @Description Add vehicle to a driver.
// @Tags drivers
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param driver body entity.DetailDriver true "Driver driver"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/driver/vehicle_add [post]
// SaveDriver is a function driver to handle create a new driver.
func (s *Drivers) AddDriverVehicle(c *gin.Context) {
	var driverEntity entity.Driver
	if err := c.ShouldBindJSON(&driverEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	_, err := s.us.GetDriver(driverEntity.UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextDriverNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextDriverNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, errDesc, errException := s.us.AddDriverVehicle(&driverEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextDriverNotFound) {
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
	driver, err := s.us.GetDriver(driverEntity.UUID)
	c.Status(http.StatusOK)
	response.NewSuccess(c, driver.DetailDriver(), success.DriverSuccessfullyAddDriverVehicle).JSON()
}

// @Summary Delete vehicle to a driver
// @Description Delete vehicle to a driver.
// @Tags drivers
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param driver body entity.DetailDriver true "Driver driver"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/driver/vehicle_del [post]
// SaveDriver is a function driver to handle create a new driver.
func (s *Drivers) DeleteDriverVehicle(c *gin.Context) {
	var driverEntity entity.Driver
	if err := c.ShouldBindJSON(&driverEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	_, err := s.us.GetDriver(driverEntity.UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextDriverNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextDriverNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, errDesc, errException := s.us.DeleteDriverVehicle(&driverEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextDriverNotFound) {
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
	driver, err := s.us.GetDriver(driverEntity.UUID)
	c.Status(http.StatusOK)
	response.NewSuccess(c, driver.DetailDriver(), success.DriverSuccessfullyDeleteDriverVehicle).JSON()
}
