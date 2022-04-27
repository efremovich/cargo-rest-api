package orderStatusTypev1point00

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

	"github.com/joho/godotenv"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestSaveOrderStatusType_Success Test.
func TestSaveOrderStatusType_Success(t *testing.T) {
	var orderStatusTypeData entity.OrderStatusType
	var orderStatusTypeApp mock.OrderStatusTypeAppInterface
	orderStatusTypeHandler := NewOrderStatusTypes(&orderStatusTypeApp)
	orderStatusTypeJSON := `{
		"type": "Паспорт"
	}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/orderStatusTypes", orderStatusTypeHandler.SaveOrderStatusType)

	orderStatusTypeApp.SaveOrderStatusTypeFn = func(orderStatusType *entity.OrderStatusType) (*entity.OrderStatusType, map[string]string, error) {
		return &entity.OrderStatusType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPost,
		"/api/v1/external/orderStatusTypes",
		bytes.NewBufferString(orderStatusTypeJSON),
	)
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &orderStatusTypeData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, orderStatusTypeData.UUID, UUID)
	assert.EqualValues(t, orderStatusTypeData.Type, "Паспорт")
}

func TestSaveOrderStatusType_InvalidData(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON:  `{"type":33, "": ""}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"type": "", "": "jija",}`,
			statusCode: 422,
		},
	}

	for _, v := range samples {
		var orderStatusTypeApp mock.OrderStatusTypeAppInterface
		orderStatusTypeHandler := NewOrderStatusTypes(&orderStatusTypeApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/orderStatusTypes", orderStatusTypeHandler.SaveOrderStatusType)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/orderStatusTypes", bytes.NewBufferString(v.inputJSON))
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

// TestUpdateOrderStatusType_Success Test.
func TestUpdateOrderStatusType_Success(t *testing.T) {
	var orderStatusTypeData entity.OrderStatusType
	var orderStatusTypeApp mock.OrderStatusTypeAppInterface
	orderStatusTypeHandler := NewOrderStatusTypes(&orderStatusTypeApp)
	orderStatusTypeJSON := `{"type": "Паспорт"}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/orderStatusTypes/:uuid", orderStatusTypeHandler.UpdateOrderStatusType)

	orderStatusTypeApp.UpdateOrderStatusTypeFn = func(UUID string, orderStatusType *entity.OrderStatusType) (*entity.OrderStatusType, map[string]string, error) {
		return &entity.OrderStatusType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil, nil
	}

	orderStatusTypeApp.GetOrderStatusTypeFn = func(string) (*entity.OrderStatusType, error) {
		return &entity.OrderStatusType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPut,
		"/api/v1/external/orderStatusTypes/"+UUID,
		bytes.NewBufferString(orderStatusTypeJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &orderStatusTypeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, orderStatusTypeData.UUID, UUID)
	assert.EqualValues(t, orderStatusTypeData.Type, "Паспорт")
}

// TestGetOrderStatusType_Success Test.
func TestGetOrderStatusType_Success(t *testing.T) {
	var orderStatusTypeData entity.OrderStatusType
	var orderStatusTypeApp mock.OrderStatusTypeAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	orderStatusTypeHandler := NewOrderStatusTypes(&orderStatusTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/orderStatusTypes/:uuid", orderStatusTypeHandler.GetOrderStatusType)

	orderStatusTypeApp.GetOrderStatusTypeFn = func(string) (*entity.OrderStatusType, error) {
		return &entity.OrderStatusType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/orderStatusTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &orderStatusTypeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, orderStatusTypeData.UUID, UUID)
	assert.EqualValues(t, orderStatusTypeData.Type, "Паспорт")
}

// TestGetOrderStatusTypes_Success Test.
func TestGetOrderStatusTypes_Success(t *testing.T) {
	var orderStatusTypeApp mock.OrderStatusTypeAppInterface
	var orderStatusTypesData []entity.OrderStatusType
	var metaData repository.Meta
	orderStatusTypeHandler := NewOrderStatusTypes(&orderStatusTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/orderStatusTypes", orderStatusTypeHandler.GetOrderStatusTypes)
	orderStatusTypeApp.GetOrderStatusTypesFn = func(params *repository.Parameters) ([]*entity.OrderStatusType, *repository.Meta, error) {
		orderStatusTypes := []*entity.OrderStatusType{
			{
				UUID: UUID,
				Type: "Паспорт",
			},
			{
				UUID: UUID,
				Type: "Свидетельство о рождении",
			},
		}
		meta := repository.NewMeta(params, int64(len(orderStatusTypes)))
		return orderStatusTypes, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/orderStatusTypes", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &orderStatusTypesData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(orderStatusTypesData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeleteOrderStatusType_Success Test.
func TestDeleteOrderStatusType_Success(t *testing.T) {
	var orderStatusTypeApp mock.OrderStatusTypeAppInterface
	orderStatusTypeHandler := NewOrderStatusTypes(&orderStatusTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/orderStatusTypes/:uuid", orderStatusTypeHandler.DeleteOrderStatusType)

	orderStatusTypeApp.DeleteOrderStatusTypeFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/orderStatusTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeleteOrderStatusType_Failed_OrderStatusTypeNotFound Test.
func TestDeleteOrderStatusType_Failed_OrderStatusTypeNotFound(t *testing.T) {
	var orderStatusTypeApp mock.OrderStatusTypeAppInterface
	orderStatusTypeHandler := NewOrderStatusTypes(&orderStatusTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/orderStatusTypes/:uuid", orderStatusTypeHandler.DeleteOrderStatusType)

	orderStatusTypeApp.DeleteOrderStatusTypeFn = func(UUID string) error {
		return exception.ErrorTextOrderStatusTypeNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/orderStatusTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
