package pricev1point00

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

// Prices is a struct defines the dependencies that will be used.
type Prices struct {
	us application.PriceAppInterface
}

// NewCountreis is constructor will initialize price handler.
func NewPrices(us application.PriceAppInterface) *Prices {
	return &Prices{
		us: us,
	}
}

// @Summary Create a new price
// @Description Create a new price.
// @Tags prices
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param price body entity.DetailPrice true "Price price"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/price [post]
// SavePrice is a function price to handle create a new price.
func (s *Prices) SavePrice(c *gin.Context) {
	var priceEntity entity.Price
	if err := c.ShouldBindJSON(&priceEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := priceEntity.ValidateSavePrice()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newPrice, errDesc, errException := s.us.SavePrice(&priceEntity)
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
	response.NewSuccess(c, newPrice.DetailPrice(), success.PriceSuccessfullyCreatePrice).JSON()
}

// @Summary Update price
// @Description Update an existing price.
// @Tags prices
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Price UUID"
// @Param name formData string true "Price name"
// @Param region formData string true "Price region"
// @Param latitude formData string true "Price latitude"
// @Param longitude formData string true "Price longitude"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/price/uuid [put]
// UpdatePrice is a function uses to handle update price by UUID.
func (s *Prices) UpdatePrice(c *gin.Context) {
	var priceEntity entity.Price
	if err := c.ShouldBindUri(&priceEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&priceEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := c.Param("uuid")
	_, err := s.us.GetPrice(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPriceNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPriceNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedPrice, errDesc, errException := s.us.UpdatePrice(UUID, &priceEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextPriceNotFound) {
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
	response.NewSuccess(c, updatedPrice.DetailPrice(), success.PriceSuccessfullyUpdatePrice).JSON()
}

// @Summary Delete price
// @Description Delete an existing price.
// @Tags prices
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Price UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/price/{uuid} [delete]
// DeletePrice is a function uses to handle delete price by UUID.
func (s *Prices) DeletePrice(c *gin.Context) {
	var priceEntity entity.Price
	if err := c.ShouldBindUri(&priceEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeletePrice(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPriceNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPriceNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.PriceSuccessfullyDeletePrice).JSON()
}

// @Summary Get prices
// @Description Get list of existing prices.
// @Tags prices
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
// @Router /api/v1/external/prices [get]
// GetPrices is a function uses to handle get price list.
func (s *Prices) GetPrices(c *gin.Context) {
	var price entity.Price
	var prices entity.Prices
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(price.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	prices, meta, err := s.us.GetPrices(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, prices.DetailPrices(), success.PriceSuccessfullyGetPriceList).WithMeta(meta).JSON()
}

// @Summary Get price
// @Description Get detail of existing price.
// @Tags prices
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Price UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/price/{uuid} [get]
// GetPrice is a function uses to handle get price detail by UUID.
func (s *Prices) GetPrice(c *gin.Context) {
	var priceEntity entity.Price
	if err := c.ShouldBindUri(&priceEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	price, err := s.us.GetPrice(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextPriceNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextPriceNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, price.DetailPrice(), success.PriceSuccessfullyGetPriceDetail).JSON()
}
