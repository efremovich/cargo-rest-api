package sityv1point00

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

// Sities is a struct defines the dependencies that will be used.
type Sities struct {
	us application.SityAppInterface
}

// NewCountreis is constructor will initialize sity handler.
func NewSities(us application.SityAppInterface) *Sities {
	return &Sities{
		us: us,
	}
}

// @Summary Create a new sity
// @Description Create a new sity.
// @Tags sities
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param name formData string true "Sity name"
// @Param latitude formData string true "Sity latitude"
// @Param longitude formData string true "Sity longitude"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/sities [post]
// SaveSity is a function sity to handle create a new sity.
func (s *Sities) SaveSities(c *gin.Context) {
	var sityEntity entity.Sity
	if err := c.ShouldBindJSON(&sityEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	newSity, errDesc, errException := s.us.SaveSity(&sityEntity)
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
	response.NewSuccess(c, newSity.DetailSity(), success.SitySuccessfullyCreateSity).JSON()
}

// @Summary Update sity
// @Description Update an existing sity.
// @Tags sities
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Sity UUID"
// @Param name formData string true "Sity name"
// @Param latitude formData string true "Sity latitude"
// @Param longitude formData string true "Sity longitude"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/sities/uuid [put]
// UpdateSity is a function uses to handle update sity by UUID.
func (s *Sities) UpdateSities(c *gin.Context) {
	var sityEntity entity.Sity
	if err := c.ShouldBindUri(&sityEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&sityEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := c.Param("uuid")
	_, err := s.us.GetSity(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextSityNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextSityNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedSity, errDesc, errException := s.us.UpdateSity(UUID, &sityEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextSityNotFound) {
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
	response.NewSuccess(c, updatedSity.DetailSity(), success.SitySuccessfullyUpdateSity).JSON()
}

// @Summary Delete sity
// @Description Delete an existing sity.
// @Tags sities
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Sity UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/sities/{uuid} [delete]
// DeleteSity is a function uses to handle delete sity by UUID.
func (s *Sities) DeleteSity(c *gin.Context) {
	var sityEntity entity.Sity
	if err := c.ShouldBindUri(&sityEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeleteSity(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextSityNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextSityNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.SitySuccessfullyDeleteSity).JSON()
}

// @Summary Get sities
// @Description Get list of existing sities.
// @Tags sities
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
// @Router /api/v1/external/sities [get]
// GetSities is a function uses to handle get sity list.
func (s *Sities) GetSities(c *gin.Context) {
	var sity entity.Sity
	var sities entity.Sities
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(sity.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	sities, meta, err := s.us.GetSities(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, sities.DetailSities(), success.SitySuccessfullyGetSityList).WithMeta(meta).JSON()
}

// @Summary Get sity
// @Description Get detail of existing sity.
// @Tags sities
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Sity UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/sities/{uuid} [get]
// GetSity is a function uses to handle get sity detail by UUID.
func (s *Sities) GetSity(c *gin.Context) {
	var sityEntity entity.Sity
	if err := c.ShouldBindUri(&sityEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	sity, err := s.us.GetSity(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextSityNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextSityNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, sity.DetailSity(), success.SitySuccessfullyGetSityDetail).JSON()
}
