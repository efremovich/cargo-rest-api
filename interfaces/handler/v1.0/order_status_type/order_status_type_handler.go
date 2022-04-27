package orderStatusTypev1point00

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

// OrderStatusTypes is a struct defines the dependencies that will be used.
type OrderStatusTypes struct {
	us application.OrderStatusTypeAppInterface
}

// NewCountreis is constructor will initialize orderStatusType handler.
func NewOrderStatusTypes(us application.OrderStatusTypeAppInterface) *OrderStatusTypes {
	return &OrderStatusTypes{
		us: us,
	}
}

// @Summary Create a new orderStatusType
// @Description Create a new orderStatusType.
// @Tags ordes status types
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param orderStatusType body entity.DetailOrderStatusType true "OrderStatusType orderStatusType"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/orderStatusType [post]
// SaveOrderStatusType is a function orderStatusType to handle create a new orderStatusType.
func (s *OrderStatusTypes) SaveOrderStatusType(c *gin.Context) {
	var orderStatusTypeEntity entity.OrderStatusType
	if err := c.ShouldBindJSON(&orderStatusTypeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := orderStatusTypeEntity.ValidateSaveOrderStatusType()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newOrderStatusType, errDesc, errException := s.us.SaveOrderStatusType(&orderStatusTypeEntity)
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
	response.NewSuccess(c, newOrderStatusType.DetailOrderStatusType(), success.OrderStatusTypeSuccessfullyCreateOrderStatusType).
		JSON()
}

// @Summary Update orderStatusType
// @Description Update an existing orderStatusType.
// @Tags ordes status types
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "OrderStatusType UUID"
// @Param name formData string true "OrderStatusType name"
// @Param region formData string true "OrderStatusType region"
// @Param latitude formData string true "OrderStatusType latitude"
// @Param longitude formData string true "OrderStatusType longitude"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/orderStatusType/uuid [put]
// UpdateOrderStatusType is a function uses to handle update orderStatusType by UUID.
func (s *OrderStatusTypes) UpdateOrderStatusType(c *gin.Context) {
	var orderStatusTypeEntity entity.OrderStatusType
	if err := c.ShouldBindUri(&orderStatusTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&orderStatusTypeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := c.Param("uuid")
	_, err := s.us.GetOrderStatusType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextOrderStatusTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextOrderStatusTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedOrderStatusType, errDesc, errException := s.us.UpdateOrderStatusType(UUID, &orderStatusTypeEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextOrderStatusTypeNotFound) {
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
	response.NewSuccess(c, updatedOrderStatusType.DetailOrderStatusType(), success.OrderStatusTypeSuccessfullyUpdateOrderStatusType).
		JSON()
}

// @Summary Delete orderStatusType
// @Description Delete an existing orderStatusType.
// @Tags ordes status types
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "OrderStatusType UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/orderStatusType/{uuid} [delete]
// DeleteOrderStatusType is a function uses to handle delete orderStatusType by UUID.
func (s *OrderStatusTypes) DeleteOrderStatusType(c *gin.Context) {
	var orderStatusTypeEntity entity.OrderStatusType
	if err := c.ShouldBindUri(&orderStatusTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeleteOrderStatusType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextOrderStatusTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextOrderStatusTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.OrderStatusTypeSuccessfullyDeleteOrderStatusType).JSON()
}

// @Summary Get orderStatusTypes
// @Description Get list of existing orderStatusTypes.
// @Tags ordes status types
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
// @Router /api/v1/external/orderStatusTypes [get]
// GetOrderStatusTypes is a function uses to handle get orderStatusType list.
func (s *OrderStatusTypes) GetOrderStatusTypes(c *gin.Context) {
	var orderStatusType entity.OrderStatusType
	var orderStatusTypes entity.OrderStatusTypes
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(orderStatusType.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	orderStatusTypes, meta, err := s.us.GetOrderStatusTypes(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, orderStatusTypes.DetailOrderStatusTypes(), success.OrderStatusTypeSuccessfullyGetOrderStatusTypeList).
		WithMeta(meta).
		JSON()
}

// @Summary Get orderStatusType
// @Description Get detail of existing orderStatusType.
// @Tags ordes status types
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "OrderStatusType UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/orderStatusType/{uuid} [get]
// GetOrderStatusType is a function uses to handle get orderStatusType detail by UUID.
func (s *OrderStatusTypes) GetOrderStatusType(c *gin.Context) {
	var orderStatusTypeEntity entity.OrderStatusType
	if err := c.ShouldBindUri(&orderStatusTypeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	orderStatusType, err := s.us.GetOrderStatusType(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextOrderStatusTypeNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextOrderStatusTypeNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, orderStatusType.DetailOrderStatusType(), success.OrderStatusTypeSuccessfullyGetOrderStatusTypeDetail).
		JSON()
}
