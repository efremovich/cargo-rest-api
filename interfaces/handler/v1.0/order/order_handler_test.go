package orderv1point00

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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// TestSaveOrder_Success Test.
func TestSaveOrder_Success(t *testing.T) {
	var orderData entity.Order
	var orderApp mock.OrderAppInterface
	orderHandler := NewOrders(&orderApp)
	UUID := uuid.New().String()

	orderDate := time.Now()
	paymentUUID := uuid.New().String()
	tripUUID := uuid.New().String()
	externalUUID := uuid.New().String()
	statusUUID := uuid.New().String()

	orderJSON := `{
  "order_date": "` + orderDate.Format(time.RFC3339) + `",
    "payment_uuid":"` + paymentUUID + `",
    "trip_uuid":"` + tripUUID + `",
    "external_uuid":"` + externalUUID + `",
    "seat":"2s",
    "status_uuid":"` + statusUUID + `"
	}`
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/order", orderHandler.SaveOrder)

	orderApp.SaveOrderFn = func(order *entity.Order) (*entity.Order, map[string]string, error) {
		return &entity.Order{
			UUID:         UUID,
			OrderDate:    orderDate,
			TripUUID:     tripUUID,
			ExternalUUID: externalUUID,
			Seat:         "2s",
			StatusUUID:   statusUUID,
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPost,
		"/api/v1/external/order",
		bytes.NewBufferString(orderJSON),
	)
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, err := json.Marshal(response["data"])

	if err != nil {
		t.Errorf("%v\n", err)
	}

	err = json.Unmarshal(data, &orderData)

	if err != nil {
		t.Errorf("%v\n", err)
	}

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, orderData.UUID, UUID)
	assert.EqualValues(t, orderData.OrderDate.Format(time.RFC3339), orderDate.Format(time.RFC3339))
	assert.EqualValues(t, orderData.TripUUID, tripUUID)
	assert.EqualValues(t, orderData.Seat, "2s")
	assert.EqualValues(t, orderData.StatusUUID, statusUUID)
}

func TestSaveOrder_InvalidData(t *testing.T) {
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
		var orderApp mock.OrderAppInterface
		orderHandler := NewOrders(&orderApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/order", orderHandler.SaveOrder)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/order", bytes.NewBufferString(v.inputJSON))
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

// TestUpdateOrder_Success Test.
func TestUpdateOrder_Success(t *testing.T) {
	var orderData entity.Order
	var orderApp mock.OrderAppInterface
	orderHandler := NewOrders(&orderApp)

	UUID := uuid.New().String()

	orderDate := time.Now()
	paymentUUID := uuid.New().String()
	tripUUID := uuid.New().String()
	externalUUID := uuid.New().String()
	statusUUID := uuid.New().String()

	orderJSON := `{
		"uuid":"` + UUID + `",
  	"order_date": "` + orderDate.Format(time.RFC3339) + `",
    "payment_uuid":"` + paymentUUID + `",
    "trip_uuid":"` + tripUUID + `",
    "external_uuid":"` + externalUUID + `",
    "seat":"2s",
    "status_uuid":"` + statusUUID + `"
	}`

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/order/:uuid", orderHandler.UpdateOrder)

	orderApp.UpdateOrderFn = func(UUID string, order *entity.Order) (*entity.Order, map[string]string, error) {
		return &entity.Order{
			UUID:         UUID,
			OrderDate:    orderDate,
			TripUUID:     tripUUID,
			ExternalUUID: externalUUID,
			Seat:         "2s",
			StatusUUID:   statusUUID,
		}, nil, nil
	}

	orderApp.GetOrderFn = func(string) (*entity.Order, error) {
		return &entity.Order{
			UUID:         UUID,
			OrderDate:    orderDate,
			TripUUID:     tripUUID,
			ExternalUUID: externalUUID,
			Seat:         "2s",
			StatusUUID:   statusUUID,
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPut,
		"/api/v1/external/order/"+UUID,
		bytes.NewBufferString(orderJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &orderData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, orderData.UUID, UUID)
	assert.EqualValues(t, orderData.OrderDate.Format(time.RFC3339), orderDate.Format(time.RFC3339))
	assert.EqualValues(t, orderData.TripUUID, tripUUID)
	assert.EqualValues(t, orderData.Seat, "2s")
	assert.EqualValues(t, orderData.StatusUUID, statusUUID)
}

// TestGetOrder_Success Test.
func TestGetOrder_Success(t *testing.T) {
	var orderData entity.Order
	var orderApp mock.OrderAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	orderHandler := NewOrders(&orderApp)
	UUID := uuid.New().String()

	orderDate := time.Now()
	tripUUID := uuid.New().String()
	externalUUID := uuid.New().String()
	statusUUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/order/:uuid", orderHandler.GetOrder)

	orderApp.GetOrderFn = func(string) (*entity.Order, error) {
		return &entity.Order{
			UUID:         UUID,
			OrderDate:    orderDate,
			TripUUID:     tripUUID,
			ExternalUUID: externalUUID,
			Seat:         "2s",
			StatusUUID:   statusUUID,
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/order/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &orderData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, orderData.UUID, UUID)
	assert.EqualValues(t, orderData.OrderDate.Format(time.RFC3339), orderDate.Format(time.RFC3339))
	assert.EqualValues(t, orderData.TripUUID, tripUUID)
	assert.EqualValues(t, orderData.Seat, "2s")
	assert.EqualValues(t, orderData.StatusUUID, statusUUID)
}

// TestGetOrders_Success Test.
func TestGetOrders_Success(t *testing.T) {
	var orderApp mock.OrderAppInterface
	var ordersData []entity.Order
	var metaData repository.Meta
	orderHandler := NewOrders(&orderApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/orders", orderHandler.GetOrders)
	orderApp.GetOrdersFn = func(params *repository.Parameters) ([]*entity.Order, *repository.Meta, error) {
		orders := []*entity.Order{
			{
				UUID:         UUID,
				OrderDate:    time.Now(),
				TripUUID:     uuid.New().String(),
				ExternalUUID: uuid.New().String(),
				Seat:         "2s",
				StatusUUID:   uuid.New().String(),
			},
			{
				UUID:         UUID,
				OrderDate:    time.Now(),
				TripUUID:     uuid.New().String(),
				ExternalUUID: uuid.New().String(),
				Seat:         "2s",
				StatusUUID:   uuid.New().String(),
			},
		}
		meta := repository.NewMeta(params, int64(len(orders)))
		return orders, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/orders", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &ordersData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(ordersData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeleteOrder_Success Test.
func TestDeleteOrder_Success(t *testing.T) {
	var orderApp mock.OrderAppInterface
	orderHandler := NewOrders(&orderApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/order/:uuid", orderHandler.DeleteOrder)

	orderApp.DeleteOrderFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/order/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeleteOrder_Failed_OrderNotFound Test.
func TestDeleteOrder_Failed_OrderNotFound(t *testing.T) {
	var orderApp mock.OrderAppInterface
	orderHandler := NewOrders(&orderApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/order/:uuid", orderHandler.DeleteOrder)

	orderApp.DeleteOrderFn = func(UUID string) error {
		return exception.ErrorTextOrderNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/order/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
