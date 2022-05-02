package routev1point00

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

// Routes is a struct defines the dependencies that will be used.
type Routes struct {
	us application.RouteAppInterface
}

// NewCountreis is constructor will initialize route handler.
func NewRoutes(us application.RouteAppInterface) *Routes {
	return &Routes{
		us: us,
	}
}

// @Summary Create a new route
// @Description Create a new route.
// @Tags routes
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param route body entity.DetailRoute true "Route route"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/route [post]
// SaveRoute is a function route to handle create a new route.
func (s *Routes) SaveRoute(c *gin.Context) {
	var routeEntity entity.Route
	if err := c.ShouldBindJSON(&routeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	validateErr := routeEntity.ValidateSaveRoute()
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}
	newRoute, errDesc, errException := s.us.SaveRoute(&routeEntity)
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
	response.NewSuccess(c, newRoute.DetailRoute(), success.RouteSuccessfullyCreateRoute).
		JSON()
}

// @Summary Update route
// @Description Update an existing route.
// @Tags routes
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param route body entity.DetailRoute true "Route route"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/route/uuid [put]
// UpdateRoute is a function uses to handle update route by UUID.
func (s *Routes) UpdateRoute(c *gin.Context) {
	var routeEntity entity.Route
	if err := c.ShouldBindUri(&routeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&routeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	UUID := routeEntity.UUID
	_, err := s.us.GetRoute(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextRouteNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextRouteNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedRoute, errDesc, errException := s.us.UpdateRoute(UUID, &routeEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextRouteNotFound) {
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
	response.NewSuccess(c, updatedRoute.DetailRoute(), success.RouteSuccessfullyUpdateRoute).
		JSON()
}

// @Summary Delete route
// @Description Delete an existing route.
// @Tags routes
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Route UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/route/{uuid} [delete]
// DeleteRoute is a function uses to handle delete route by UUID.
func (s *Routes) DeleteRoute(c *gin.Context) {
	var routeEntity entity.Route
	if err := c.ShouldBindUri(&routeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	err := s.us.DeleteRoute(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextRouteNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextRouteNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, nil, success.RouteSuccessfullyDeleteRoute).JSON()
}

// @Summary Get routes
// @Description Get list of existing routes.
// @Tags routes
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
// @Router /api/v1/external/routes [get]
// GetRoutes is a function uses to handle get route list.
func (s *Routes) GetRoutes(c *gin.Context) {
	var route entity.Route
	var routes entity.Routes
	var err error
	parameters := repository.NewGinParameters(c)
	validateErr := parameters.ValidateParameter(route.FilterableFields()...)
	if len(validateErr) > 0 {
		exceptionData := response.TranslateErrorForm(c, validateErr)
		c.Set("data", exceptionData)
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	routes, meta, err := s.us.GetRoutes(parameters)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response.NewSuccess(c, routes.DetailRoutes(), success.RouteSuccessfullyGetRouteList).
		WithMeta(meta).
		JSON()
}

// @Summary Get route
// @Description Get detail of existing route.
// @Tags routes
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param uuid path string true "Route UUID"
// @Success 200 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/route/{uuid} [get]
// GetRoute is a function uses to handle get route detail by UUID.
func (s *Routes) GetRoute(c *gin.Context) {
	var routeEntity entity.Route
	if err := c.ShouldBindUri(&routeEntity.UUID); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, exception.ErrorTextBadRequest)
		return
	}

	UUID := c.Param("uuid")
	route, err := s.us.GetRoute(UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextRouteNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextRouteNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response.NewSuccess(c, route.DetailRoute(), success.RouteSuccessfullyGetRouteDetail).
		JSON()
}

// @Summary Add price to a route
// @Description Add price to a route.
// @Tags routes
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param route body entity.DetailRoute true "Route route"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/route/price_add [post]
// SaveRoute is a function route to handle create a new route.
func (s *Routes) AddRoutePrice(c *gin.Context) {
	var routeEntity entity.Route
	if err := c.ShouldBindJSON(&routeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	_, err := s.us.GetRoute(routeEntity.UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextRouteNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextRouteNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, errDesc, errException := s.us.AddRoutePrice(&routeEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextRouteNotFound) {
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
	route, err := s.us.GetRoute(routeEntity.UUID)
	c.Status(http.StatusOK)
	response.NewSuccess(c, route.DetailRoute(), success.RouteSuccessfullyAddRoutePrice).JSON()
}

// @Summary Delete price to a route
// @Description Delete price to a route.
// @Tags routes
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language code" Enums(en, ru) default(en)
// @Param Set-Request-Id header string false "Request id"
// @Security BasicAuth
// @Security JWTAuth
// @Param route body entity.DetailRoute true "Route route"
// @Success 201 {object} response.successOutput
// @Failure 400 {object} response.errorOutput
// @Failure 401 {object} response.errorOutput
// @Failure 403 {object} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/v1/external/route/price_del [post]
// SaveRoute is a function route to handle create a new route.
func (s *Routes) DeleteRoutePrice(c *gin.Context) {
	var routeEntity entity.Route
	if err := c.ShouldBindJSON(&routeEntity); err != nil {
		_ = c.AbortWithError(http.StatusUnprocessableEntity, exception.ErrorTextUnprocessableEntity)
		return
	}

	_, err := s.us.GetRoute(routeEntity.UUID)
	if err != nil {
		if errors.Is(err, exception.ErrorTextRouteNotFound) {
			_ = c.AbortWithError(http.StatusNotFound, exception.ErrorTextRouteNotFound)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, errDesc, errException := s.us.DeleteRoutePrice(&routeEntity)
	if errException != nil {
		c.Set("data", errDesc)
		if errors.Is(errException, exception.ErrorTextRouteNotFound) {
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
	route, err := s.us.GetRoute(routeEntity.UUID)
	c.Status(http.StatusOK)
	response.NewSuccess(c, route.DetailRoute(), success.RouteSuccessfullyDeleteRoutePrice).JSON()
}
