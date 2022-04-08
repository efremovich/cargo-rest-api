package passengerTypev1point00

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

// PassengerTypes is a struct defines the dependencies that will be used.
type PassengerTypes struct {
	us application.PassengerTypeAppInterface
}

// NewCountreis is constructor will initialize passengerType handler.
func NewPassengerTypes(us application.PassengerTypeAppInterface) *PassengerTypes {
	return &PassengerTypes{
		us: us,
	}
}

// @Summary Create a new passengerType
// @Description Create a new passengerType.
// @Tags passengerTypes
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param passengerType body entity.DetailPassengerType true "PassengerType passengerType"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/passengerTypes [post]
// SavePassengerType is a function passengerType to handle create a new passengerType.
func (s *PassengerTypes) SavePassengerType(c *gin.Context) {
	var passengerTypeEntity entity.PassengerType
	if err := c.ShouldBindJSON(&passengerTypeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := passengerTypeEntity.ValidateSavePassengerType()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newPassengerType, errDesc, errException := s.us.SavePassengerType(&passengerTypeEntity)
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
	response.NewSuccess(c, newPassengerType.DetailPassengerType(), success.PassengerTypeSuccessfullyCreatePassengerType).JSON()
}

// @Summary Update passengerType
// @Description Update an existing passengerType.
// @Tags passengerTypes
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "PassengerType UUID"
// @Param name formData string true "PassengerType name"
// @Param region formData string true "PassengerType region"
// @Param latitude formData string true "PassengerType latitude"
// @Param longitude formData string true "PassengerType longitude"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/passengerTypes/uuid [put]
// UpdatePassengerType is a function uses to handle update passengerType by UUID.
func (s *PassengerTypes) UpdatePassengerType(c *gin.Context) {
	var passengerTypeEntity entity.PassengerType
	if err := c.ShouldBindUri(&passengerTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&passengerTypeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := c.Param("uuid")
	_, err := s.us.GetPassengerType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPassengerTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPassengerTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedPassengerType, errDesc, errException := s.us.UpdatePassengerType(UUID, &passengerTypeEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextPassengerTypeNotFound) {
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
	response.NewSuccess(c, updatedPassengerType.DetailPassengerType(), success.PassengerTypeSuccessfullyUpdatePassengerType).JSON()
}

// @Summary Delete passengerType
// @Description Delete an existing passengerType.
// @Tags passengerTypes
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "PassengerType UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/passengerTypes/{uuid} [delete]
// DeletePassengerType is a function uses to handle delete passengerType by UUID.
func (s *PassengerTypes) DeletePassengerType(c *gin.Context) {
	var passengerTypeEntity entity.PassengerType
	if err := c.ShouldBindUri(&passengerTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeletePassengerType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPassengerTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPassengerTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.PassengerTypeSuccessfullyDeletePassengerType).JSON()
}

// @Summary Get passengerTypes
// @Description Get list of existing passengerTypes.
// @Tags passengerTypes
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
// @Router /api/v1/external/passengerTypes [get]
// GetPassengerTypes is a function uses to handle get passengerType list.
func (s *PassengerTypes) GetPassengerTypes(c *gin.Context) {
	var passengerType entity.PassengerType
	var passengerTypes entity.PassengerTypes
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(passengerType.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	passengerTypes, meta, err := s.us.GetPassengerTypes(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, passengerTypes.DetailPassengerTypes(), success.PassengerTypeSuccessfullyGetPassengerTypeList).WithMeta(meta).JSON()
}

// @Summary Get passengerType
// @Description Get detail of existing passengerType.
// @Tags passengerTypes
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "PassengerType UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/passengerTypes/{uuid} [get]
// GetPassengerType is a function uses to handle get passengerType detail by UUID.
func (s *PassengerTypes) GetPassengerType(c *gin.Context) {
	var passengerTypeEntity entity.PassengerType
	if err := c.ShouldBindUri(&passengerTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	passengerType, err := s.us.GetPassengerType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPassengerTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPassengerTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, passengerType.DetailPassengerType(), success.PassengerTypeSuccessfullyGetPassengerTypeDetail).JSON()
}
