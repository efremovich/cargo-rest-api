package routev1point00

import (
	"bytes"
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
	"cargo-rest-api/infrastructure/message/exception"
	"cargo-rest-api/pkg/encoder"
	"cargo-rest-api/pkg/util"
	"cargo-rest-api/tests/mock"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// TestSaveRoute_Success Test.
func TestSaveRoute_Success(t *testing.T) {
	var routeData entity.Route
	var routeApp mock.RouteAppInterface
	routeHandler := NewRoutes(&routeApp)
	UUID := uuid.New().String()
	FromUUID := uuid.New().String()
	ToUUID := uuid.New().String()
	PriceUUID := uuid.New().String()
	PassengerTypeUUID := uuid.New().String()

	routeJSON := `{
		"frome_uuid": "` + FromUUID + `",
    "to_uuid":"` + ToUUID + `",
    "distance":124,
    "distance_time":134,
    "price":[
       {"uuid": "` + PriceUUID + `", "passenger_type_uuid": "` + PassengerTypeUUID + `", "price":1}
    ]
	}`

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/route", routeHandler.SaveRoute)

	routeApp.SaveRouteFn = func(route *entity.Route) (*entity.Route, map[string]string, error) {
		return &entity.Route{
			UUID:         UUID,
			FromUUID:     FromUUID,
			ToUUID:       ToUUID,
			Distance:     124,
			DistanceTime: 134,
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPost,
		"/api/v1/external/route",
		bytes.NewBufferString(routeJSON),
	)
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &routeData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, routeData.UUID, UUID)
	assert.EqualValues(t, routeData.FromUUID, FromUUID)
	assert.EqualValues(t, routeData.ToUUID, ToUUID)
	assert.EqualValues(t, routeData.Distance, 124)
	assert.EqualValues(t, routeData.DistanceTime, 134)
}

func TestSaveRoute_InvalidData(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON:  `{"uuid":33, "": ""}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"from_uuid": "", "": "jija",}`,
			statusCode: 422,
		},
	}

	for _, v := range samples {
		var routeApp mock.RouteAppInterface
		routeHandler := NewRoutes(&routeApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/route", routeHandler.SaveRoute)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/route", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		r.ServeHTTP(w, c.Request)

		validationErr := make(map[string]string)
		response := encoder.ResponseDecoder(w.Body)
		data, _ := json.Marshal(response["data"])

		err = json.Unmarshal(data, &validationErr)
		if err != nil {
			t.Errorf("error unmarshalling error %s\n", err)
		}
		assert.Equal(t, w.Code, v.statusCode)
	}
}

// TestUpdateRoute_Success Test.
func TestUpdateRoute_Success(t *testing.T) {
	var routeData entity.Route
	var routeApp mock.RouteAppInterface
	routeHandler := NewRoutes(&routeApp)

	UUID := uuid.New().String()
	FromUUID := uuid.New().String()
	ToUUID := uuid.New().String()
	PriceUUID := uuid.New().String()
	PassengerTypeUUID := uuid.New().String()

	routeJSON := `{
		"uuid":"` + UUID + `",
		"frome_uuid": "` + FromUUID + `",
    "to_uuid":"` + ToUUID + `",
    "distance":124,
    "distance_time":134,
    "price":[
       {"uuid": "` + PriceUUID + `", "passenger_type_uuid": "` + PassengerTypeUUID + `", "price":1}
    ]
	}`

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/route/:uuid", routeHandler.UpdateRoute)

	routeApp.UpdateRouteFn = func(UUID string, route *entity.Route) (*entity.Route, map[string]string, error) {
		return &entity.Route{
			UUID:         UUID,
			FromUUID:     FromUUID,
			ToUUID:       ToUUID,
			Distance:     124,
			DistanceTime: 134,
		}, nil, nil
	}

	routeApp.GetRouteFn = func(string) (*entity.Route, error) {
		return &entity.Route{
			UUID:         UUID,
			FromUUID:     FromUUID,
			ToUUID:       ToUUID,
			Distance:     124,
			DistanceTime: 134,
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPut,
		"/api/v1/external/route/"+UUID,
		bytes.NewBufferString(routeJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &routeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, routeData.UUID, UUID)
	assert.EqualValues(t, routeData.FromUUID, FromUUID)
	assert.EqualValues(t, routeData.ToUUID, ToUUID)
	assert.EqualValues(t, routeData.Distance, 124)
	assert.EqualValues(t, routeData.DistanceTime, 134)
}

// TestGetRoute_Success Test.
func TestGetRoute_Success(t *testing.T) {
	var routeData entity.Route
	var routeApp mock.RouteAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	routeHandler := NewRoutes(&routeApp)
	UUID := uuid.New().String()
	FromUUID := uuid.New().String()
	ToUUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/route/:uuid", routeHandler.GetRoute)

	routeApp.GetRouteFn = func(string) (*entity.Route, error) {
		return &entity.Route{
			UUID:         UUID,
			FromUUID:     FromUUID,
			ToUUID:       ToUUID,
			Distance:     124,
			DistanceTime: 134,
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/route/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &routeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, routeData.UUID, UUID)
	assert.EqualValues(t, routeData.FromUUID, FromUUID)
	assert.EqualValues(t, routeData.ToUUID, ToUUID)
	assert.EqualValues(t, routeData.Distance, 124)
	assert.EqualValues(t, routeData.DistanceTime, 134)
}

// TestGetRoutes_Success Test.
func TestGetRoutes_Success(t *testing.T) {
	var routeApp mock.RouteAppInterface
	var routesData []entity.Route
	var metaData repository.Meta
	routeHandler := NewRoutes(&routeApp)
	UUID := uuid.New().String()
	FromUUID := uuid.New().String()
	ToUUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/routes", routeHandler.GetRoutes)
	routeApp.GetRoutesFn = func(params *repository.Parameters) ([]*entity.Route, *repository.Meta, error) {
		routes := []*entity.Route{
			{
				UUID:         UUID,
				FromUUID:     FromUUID,
				ToUUID:       ToUUID,
				Distance:     124,
				DistanceTime: 134,
			},
			{
				UUID:         UUID,
				FromUUID:     FromUUID,
				ToUUID:       ToUUID,
				Distance:     124,
				DistanceTime: 134,
			},
		}
		meta := repository.NewMeta(params, int64(len(routes)))
		return routes, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/routes", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &routesData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(routesData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeleteRoute_Success Test.
func TestDeleteRoute_Success(t *testing.T) {
	var routeApp mock.RouteAppInterface
	routeHandler := NewRoutes(&routeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/route/:uuid", routeHandler.DeleteRoute)

	routeApp.DeleteRouteFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/route/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeleteRoute_Failed_RouteNotFound Test.
func TestDeleteRoute_Failed_RouteNotFound(t *testing.T) {
	var routeApp mock.RouteAppInterface
	routeHandler := NewRoutes(&routeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/route/:uuid", routeHandler.DeleteRoute)

	routeApp.DeleteRouteFn = func(UUID string) error {
		return exception.ErrorTextRouteNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/route/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestRoutes_AddRoutePrice(t *testing.T) {
	var routeData entity.Route
	var routeApp mock.RouteAppInterface
	routeHandler := NewRoutes(&routeApp)

	UUID := uuid.New().String()
	priceUUID := uuid.New().String()

	routeJSON := `{"UUID":"` + UUID + `", "prices":[{"uuid":"` + priceUUID + `"}]}`

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/route/price_add/", routeHandler.AddRoutePrice)

	routeApp.AddRoutePriceFn = func(route *entity.Route) (*entity.Route, map[string]string, error) {
		return &entity.Route{
			UUID:   UUID,
			Prices: []*entity.Price{{UUID: priceUUID}},
		}, nil, nil
	}

	routeApp.GetRouteFn = func(string) (*entity.Route, error) {
		return &entity.Route{
			UUID:   UUID,
			Prices: []*entity.Price{{UUID: priceUUID}},
		}, nil
	}
	var err error

	c.Request, err = http.NewRequest(
		http.MethodPost,
		"/api/v1/external/route/price_add/",
		bytes.NewBufferString(routeJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &routeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, routeData.UUID, UUID)

}

func TestRoutes_DeleteRoutePrice(t *testing.T) {
	var routeData entity.Route
	var routeApp mock.RouteAppInterface
	routeHandler := NewRoutes(&routeApp)

	UUID := uuid.New().String()
	priceUUID := uuid.New().String()

	routeJSON := `{"UUID":"` + UUID + `", "prices":[{"uuid":"` + priceUUID + `"}]}`

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/route/price_del/", routeHandler.DeleteRoutePrice)

	routeApp.DeleteRoutePriceFn = func(route *entity.Route) (*entity.Route, map[string]string, error) {
		return &entity.Route{
			UUID:   UUID,
			Prices: []*entity.Price{{UUID: priceUUID}},
		}, nil, nil
	}

	routeApp.GetRouteFn = func(string) (*entity.Route, error) {
		return &entity.Route{
			UUID:   UUID,
			Prices: []*entity.Price{{UUID: priceUUID}},
		}, nil
	}
	var err error

	c.Request, err = http.NewRequest(
		http.MethodPost,
		"/api/v1/external/route/price_del/",
		bytes.NewBufferString(routeJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &routeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, routeData.UUID, UUID)

}
