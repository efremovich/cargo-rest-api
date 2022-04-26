package regularityTypev1point00

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

// RegularityTypes is a struct defines the dependencies that will be used.
type RegularityTypes struct {
	us application.RegularityTypeAppInterface
}

// NewCountreis is constructor will initialize regularityType handler.
func NewRegularityTypes(us application.RegularityTypeAppInterface) *RegularityTypes {
	return &RegularityTypes{
		us: us,
	}
}

// @Summary Create a new regularityType
// @Description Create a new regularityType.
// @Tags document types
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param regularityType body entity.DetailRegularityType true "RegularityType regularityType"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/regularityType [post]
// SaveRegularityType is a function regularityType to handle create a new regularityType.
func (s *RegularityTypes) SaveRegularityType(c *gin.Context) {
	var regularityTypeEntity entity.RegularityType
	if err := c.ShouldBindJSON(&regularityTypeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := regularityTypeEntity.ValidateSaveRegularityType()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newRegularityType, errDesc, errException := s.us.SaveRegularityType(&regularityTypeEntity)
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
	response.NewSuccess(c, newRegularityType.DetailRegularityType(), success.RegularityTypeSuccessfullyCreateRegularityType).
		JSON()
}

// @Summary Update regularityType
// @Description Update an existing regularityType.
// @Tags document types
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "RegularityType UUID"
// @Param name formData string true "RegularityType name"
// @Param region formData string true "RegularityType region"
// @Param latitude formData string true "RegularityType latitude"
// @Param longitude formData string true "RegularityType longitude"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/regularityType/uuid [put]
// UpdateRegularityType is a function uses to handle update regularityType by UUID.
func (s *RegularityTypes) UpdateRegularityType(c *gin.Context) {
	var regularityTypeEntity entity.RegularityType
	if err := c.ShouldBindUri(&regularityTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&regularityTypeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := c.Param("uuid")
	_, err := s.us.GetRegularityType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextRegularityTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextRegularityTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedRegularityType, errDesc, errException := s.us.UpdateRegularityType(UUID, &regularityTypeEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextRegularityTypeNotFound) {
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
	response.NewSuccess(c, updatedRegularityType.DetailRegularityType(), success.RegularityTypeSuccessfullyUpdateRegularityType).
		JSON()
}

// @Summary Delete regularityType
// @Description Delete an existing regularityType.
// @Tags document types
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "RegularityType UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/regularityType/{uuid} [delete]
// DeleteRegularityType is a function uses to handle delete regularityType by UUID.
func (s *RegularityTypes) DeleteRegularityType(c *gin.Context) {
	var regularityTypeEntity entity.RegularityType
	if err := c.ShouldBindUri(&regularityTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeleteRegularityType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextRegularityTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextRegularityTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.RegularityTypeSuccessfullyDeleteRegularityType).JSON()
}

// @Summary Get regularityTypes
// @Description Get list of existing regularityTypes.
// @Tags document types
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
// @Router /api/v1/external/regularityTypes [get]
// GetRegularityTypes is a function uses to handle get regularityType list.
func (s *RegularityTypes) GetRegularityTypes(c *gin.Context) {
	var regularityType entity.RegularityType
	var regularityTypes entity.RegularityTypes
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(regularityType.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	regularityTypes, meta, err := s.us.GetRegularityTypes(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, regularityTypes.DetailRegularityTypes(), success.RegularityTypeSuccessfullyGetRegularityTypeList).
		WithMeta(meta).
		JSON()
}

// @Summary Get regularityType
// @Description Get detail of existing regularityType.
// @Tags document types
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "RegularityType UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/regularityType/{uuid} [get]
// GetRegularityType is a function uses to handle get regularityType detail by UUID.
func (s *RegularityTypes) GetRegularityType(c *gin.Context) {
	var regularityTypeEntity entity.RegularityType
	if err := c.ShouldBindUri(&regularityTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	regularityType, err := s.us.GetRegularityType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextRegularityTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextRegularityTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, regularityType.DetailRegularityType(), success.RegularityTypeSuccessfullyGetRegularityTypeDetail).
		JSON()
}
