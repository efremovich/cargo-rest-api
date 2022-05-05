package paymentv1point00

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

// Payments is a struct defines the dependencies that will be used.
type Payments struct {
	us application.PaymentAppInterface
}

// NewCountreis is constructor will initialize payment handler.
func NewPayments(us application.PaymentAppInterface) *Payments {
	return &Payments{
		us: us,
	}
}

// @Summary Create a new payment
// @Description Create a new payment.
// @Tags payments
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param payment body entity.DetailPayment true "Payment payment"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/payment [post]
// SavePayment is a function payment to handle create a new payment.
func (s *Payments) SavePayment(c *gin.Context) {
	var paymentEntity entity.Payment
	if err := c.ShouldBindJSON(&paymentEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := paymentEntity.ValidateSavePayment()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newPayment, errDesc, errException := s.us.SavePayment(&paymentEntity)
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
	response.NewSuccess(c, newPayment.DetailPayment(), success.PaymentSuccessfullyCreatePayment).
		JSON()
}

// @Summary Update payment
// @Description Update an existing payment.
// @Tags payments
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param payment body entity.DetailPayment true "Payment payment"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/payment/uuid [put]
// UpdatePayment is a function uses to handle update payment by UUID.
func (s *Payments) UpdatePayment(c *gin.Context) {
	var paymentEntity entity.Payment
	if err := c.ShouldBindUri(&paymentEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&paymentEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := paymentEntity.UUID
	_, err := s.us.GetPayment(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPaymentNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPaymentNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedPayment, errDesc, errException := s.us.UpdatePayment(UUID, &paymentEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextPaymentNotFound) {
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
	response.NewSuccess(c, updatedPayment.DetailPayment(), success.PaymentSuccessfullyUpdatePayment).
		JSON()
}

// @Summary Delete payment
// @Description Delete an existing payment.
// @Tags payments
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Payment UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/payment/{uuid} [delete]
// DeletePayment is a function uses to handle delete payment by UUID.
func (s *Payments) DeletePayment(c *gin.Context) {
	var paymentEntity entity.Payment
	if err := c.ShouldBindUri(&paymentEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeletePayment(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPaymentNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPaymentNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.PaymentSuccessfullyDeletePayment).JSON()
}

// @Summary Get payments
// @Description Get list of existing payments.
// @Tags payments
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
// @Router /api/v1/external/payments [get]
// GetPayments is a function uses to handle get payment list.
func (s *Payments) GetPayments(c *gin.Context) {
	var payment entity.Payment
	var payments entity.Payments
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(payment.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	payments, meta, err := s.us.GetPayments(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, payments.DetailPayments(), success.PaymentSuccessfullyGetPaymentList).
		WithMeta(meta).
		JSON()
}

// @Summary Get payment
// @Description Get detail of existing payment.
// @Tags payments
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Payment UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/payment/{uuid} [get]
// GetPayment is a function uses to handle get payment detail by UUID.
func (s *Payments) GetPayment(c *gin.Context) {
	var paymentEntity entity.Payment
	if err := c.ShouldBindUri(&paymentEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	payment, err := s.us.GetPayment(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPaymentNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPaymentNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, payment.DetailPayment(), success.PaymentSuccessfullyGetPaymentDetail).
		JSON()
}

// @Summary Add order to a payment
// @Description Add order to a payment.
// @Tags payments
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param payment body entity.DetailPayment true "Payment payment"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/payment/order_add [post]
// SavePayment is a function payment to handle create a new payment.
func (s *Payments) AddOrderPayment(c *gin.Context) {
	var paymentEntity entity.Payment
	if err := c.ShouldBindJSON(&paymentEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	_, err := s.us.GetPayment(paymentEntity.UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPaymentNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPaymentNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, errDesc, errException := s.us.AddOrderPayment(&paymentEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextPaymentNotFound) {
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
	payment, _ := s.us.GetPayment(paymentEntity.UUID)
	c.Status(http.StatusOK)
	response.NewSuccess(c, payment.DetailPayment(), success.PaymentSuccessfullyAddOrderPayment).JSON()
}

// @Summary Delete order to a payment
// @Description Delete order to a payment.
// @Tags payments
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param payment body entity.DetailPayment true "Payment payment"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/payment/order_del [post]
// SavePayment is a function payment to handle create a new payment.
func (s *Payments) DeleteOrderPayment(c *gin.Context) {
	var paymentEntity entity.Payment
	if err := c.ShouldBindJSON(&paymentEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	_, err := s.us.GetPayment(paymentEntity.UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPaymentNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPaymentNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, errDesc, errException := s.us.DeleteOrderPayment(&paymentEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextPaymentNotFound) {
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
	payment, _ := s.us.GetPayment(paymentEntity.UUID)
	c.Status(http.StatusOK)
	response.NewSuccess(c, payment.DetailPayment(), success.PaymentSuccessfullyDeleteOrderPayment).JSON()
}
