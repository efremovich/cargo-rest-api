package documentTypev1point00

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

// DocumentTypes is a struct defines the dependencies that will be used.
type DocumentTypes struct {
	us application.DocumentTypeAppInterface
}

// NewCountreis is constructor will initialize documentType handler.
func NewDocumentTypes(us application.DocumentTypeAppInterface) *DocumentTypes {
	return &DocumentTypes{
		us: us,
	}
}

// @Summary Create a new documentType
// @Description Create a new documentType.
// @Tags document types
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param documentType body entity.DetailDocumentType true "DocumentType documentType"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/documentType [post]
// SaveDocumentType is a function documentType to handle create a new documentType.
func (s *DocumentTypes) SaveDocumentType(c *gin.Context) {
	var documentTypeEntity entity.DocumentType
	if err := c.ShouldBindJSON(&documentTypeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := documentTypeEntity.ValidateSaveDocumentType()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newDocumentType, errDesc, errException := s.us.SaveDocumentType(&documentTypeEntity)
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
	response.NewSuccess(c, newDocumentType.DetailDocumentType(), success.DocumentTypeSuccessfullyCreateDocumentType).JSON()
}

// @Summary Update documentType
// @Description Update an existing documentType.
// @Tags document types
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "DocumentType UUID"
// @Param name formData string true "DocumentType name"
// @Param region formData string true "DocumentType region"
// @Param latitude formData string true "DocumentType latitude"
// @Param longitude formData string true "DocumentType longitude"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/documentType/uuid [put]
// UpdateDocumentType is a function uses to handle update documentType by UUID.
func (s *DocumentTypes) UpdateDocumentType(c *gin.Context) {
	var documentTypeEntity entity.DocumentType
	if err := c.ShouldBindUri(&documentTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&documentTypeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := c.Param("uuid")
	_, err := s.us.GetDocumentType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextDocumentTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextDocumentTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedDocumentType, errDesc, errException := s.us.UpdateDocumentType(UUID, &documentTypeEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextDocumentTypeNotFound) {
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
	response.NewSuccess(c, updatedDocumentType.DetailDocumentType(), success.DocumentTypeSuccessfullyUpdateDocumentType).JSON()
}

// @Summary Delete documentType
// @Description Delete an existing documentType.
// @Tags document types
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "DocumentType UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/documentType/{uuid} [delete]
// DeleteDocumentType is a function uses to handle delete documentType by UUID.
func (s *DocumentTypes) DeleteDocumentType(c *gin.Context) {
	var documentTypeEntity entity.DocumentType
	if err := c.ShouldBindUri(&documentTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeleteDocumentType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextDocumentTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextDocumentTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.DocumentTypeSuccessfullyDeleteDocumentType).JSON()
}

// @Summary Get documentTypes
// @Description Get list of existing documentTypes.
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
// @Router /api/v1/external/documentTypes [get]
// GetDocumentTypes is a function uses to handle get documentType list.
func (s *DocumentTypes) GetDocumentTypes(c *gin.Context) {
	var documentType entity.DocumentType
	var documentTypes entity.DocumentTypes
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(documentType.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	documentTypes, meta, err := s.us.GetDocumentTypes(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, documentTypes.DetailDocumentTypes(), success.DocumentTypeSuccessfullyGetDocumentTypeList).
		WithMeta(meta).
		JSON()
}

// @Summary Get documentType
// @Description Get detail of existing documentType.
// @Tags document types
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "DocumentType UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/documentType/{uuid} [get]
// GetDocumentType is a function uses to handle get documentType detail by UUID.
func (s *DocumentTypes) GetDocumentType(c *gin.Context) {
	var documentTypeEntity entity.DocumentType
	if err := c.ShouldBindUri(&documentTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	documentType, err := s.us.GetDocumentType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextDocumentTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextDocumentTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, documentType.DetailDocumentType(), success.DocumentTypeSuccessfullyGetDocumentTypeDetail).JSON()
}
