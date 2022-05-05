package orderv1point00

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

// Orders is a struct defines the dependencies that will be used.
type Orders struct {
	us application.OrderAppInterface
}

// NewCountreis is constructor will initialize order handler.
func NewOrders(us application.OrderAppInterface) *Orders {
	return &Orders{
		us: us,
	}
}

// @Summary Create a new order
// @Description Create a new order.
// @Tags orders
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param order body entity.DetailOrder true "Order order"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/order [post]
// SaveOrder is a function order to handle create a new order.
func (s *Orders) SaveOrder(c *gin.Context) {
	var orderEntity entity.Order
	if err := c.ShouldBindJSON(&orderEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := orderEntity.ValidateSaveOrder()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newOrder, errDesc, errException := s.us.SaveOrder(&orderEntity)
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
	response.NewSuccess(c, newOrder.DetailOrder(), success.OrderSuccessfullyCreateOrder).
		JSON()
}

// @Summary Update order
// @Description Update an existing order.
// @Tags orders
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param order body entity.DetailOrder true "Order order"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/order/uuid [put]
// UpdateOrder is a function uses to handle update order by UUID.
func (s *Orders) UpdateOrder(c *gin.Context) {
	var orderEntity entity.Order
	if err := c.ShouldBindUri(&orderEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&orderEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := orderEntity.UUID
	_, err := s.us.GetOrder(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextOrderNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextOrderNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedOrder, errDesc, errException := s.us.UpdateOrder(UUID, &orderEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextOrderNotFound) {
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
	response.NewSuccess(c, updatedOrder.DetailOrder(), success.OrderSuccessfullyUpdateOrder).
		JSON()
}

// @Summary Delete order
// @Description Delete an existing order.
// @Tags orders
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Order UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/order/{uuid} [delete]
// DeleteOrder is a function uses to handle delete order by UUID.
func (s *Orders) DeleteOrder(c *gin.Context) {
	var orderEntity entity.Order
	if err := c.ShouldBindUri(&orderEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeleteOrder(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextOrderNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextOrderNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.OrderSuccessfullyDeleteOrder).JSON()
}

// @Summary Get orders
// @Description Get list of existing orders.
// @Tags orders
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
// @Router /api/v1/external/orders [get]
// GetOrders is a function uses to handle get order list.
func (s *Orders) GetOrders(c *gin.Context) {
	var order entity.Order
	var orders entity.Orders
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(order.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	orders, meta, err := s.us.GetOrders(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, orders.DetailOrders(), success.OrderSuccessfullyGetOrderList).
		WithMeta(meta).
		JSON()
}

// @Summary Get order
// @Description Get detail of existing order.
// @Tags orders
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Order UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/order/{uuid} [get]
// GetOrder is a function uses to handle get order detail by UUID.
func (s *Orders) GetOrder(c *gin.Context) {
	var orderEntity entity.Order
	if err := c.ShouldBindUri(&orderEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	order, err := s.us.GetOrder(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextOrderNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextOrderNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, order.DetailOrder(), success.OrderSuccessfullyGetOrderDetail).
		JSON()
}
