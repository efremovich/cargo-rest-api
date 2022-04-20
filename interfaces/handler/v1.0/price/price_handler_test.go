package pricev1point00

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

// TestSavePrice_Success Test.
func TestSavePrice_Success(t *testing.T) {
	var priceData entity.Price
	var priceApp mock.PriceAppInterface
	priceHandler := NewPrices(&priceApp)
	priceJSON := `{"passenger_type_uuid": "503f4ab8-5bf2-409c-a469-8da4b614232c","price": 150.00}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/price", priceHandler.SavePrice)

	priceApp.SavePriceFn = func(price *entity.Price) (*entity.Price, map[string]string, error) {
		return &entity.Price{
			UUID:              UUID,
			PassengerTypeUUID: "503f4ab8-5bf2-409c-a469-8da4b614232c",
			Price:             150.00,
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/price", bytes.NewBufferString(priceJSON))
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &priceData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, priceData.UUID, UUID)
	assert.EqualValues(t, priceData.PassengerTypeUUID, "503f4ab8-5bf2-409c-a469-8da4b614232c")
	assert.EqualValues(t, priceData.Price, 150.00)
}

func TestSavePrice_InvalidData(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON:  `{"passenger_type_uuid":"", "price": ""}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"passenger_type_uuid": "22", "price": "Двадцать два",}`,
			statusCode: 422,
		},
	}

	for _, v := range samples {
		var priceApp mock.PriceAppInterface
		priceHandler := NewPrices(&priceApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/price", priceHandler.SavePrice)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/price", bytes.NewBufferString(v.inputJSON))
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

// TestUpdatePrice_Success Test.
func TestUpdatePrice_Success(t *testing.T) {
	var priceData entity.Price
	var priceApp mock.PriceAppInterface
	priceHandler := NewPrices(&priceApp)
	priceJSON := `{"passenger_type_uuid": "503f4ab8-5bf2-409c-a469-8da4b614232c", "price":150.22}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/price/:uuid", priceHandler.UpdatePrice)

	priceApp.UpdatePriceFn = func(UUID string, price *entity.Price) (*entity.Price, map[string]string, error) {
		return &entity.Price{
			UUID:              UUID,
			PassengerTypeUUID: "503f4ab8-5bf2-409c-a469-8da4b614232c",
			Price:             150.22,
		}, nil, nil
	}

	priceApp.GetPriceFn = func(string) (*entity.Price, error) {
		return &entity.Price{
			UUID:              UUID,
			PassengerTypeUUID: "503f4ab8-5bf2-409c-a469-8da4b614232c",
			Price:             150.22,
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPut, "/api/v1/external/price/"+UUID, bytes.NewBufferString(priceJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &priceData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, priceData.UUID, UUID)
	assert.EqualValues(t, priceData.PassengerTypeUUID, "503f4ab8-5bf2-409c-a469-8da4b614232c")
	assert.EqualValues(t, priceData.Price, 150.22)
}

// TestGetPrice_Success Test.
func TestGetPrice_Success(t *testing.T) {
	var priceData entity.Price
	var priceApp mock.PriceAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	priceHandler := NewPrices(&priceApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/price/:uuid", priceHandler.GetPrice)

	priceApp.GetPriceFn = func(string) (*entity.Price, error) {
		return &entity.Price{
			UUID:              UUID,
			PassengerTypeUUID: "503f4ab8-5bf2-409c-a469-8da4b614232c",
			Price:             150.22,
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/price/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &priceData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, priceData.UUID, UUID)
	assert.EqualValues(t, priceData.PassengerTypeUUID, "503f4ab8-5bf2-409c-a469-8da4b614232c")
	assert.EqualValues(t, priceData.Price, 150.22)
}

// TestGetPrices_Success Test.
func TestGetPrices_Success(t *testing.T) {
	var priceApp mock.PriceAppInterface
	var pricesData []entity.Price
	var metaData repository.Meta
	priceHandler := NewPrices(&priceApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/prices", priceHandler.GetPrices)
	priceApp.GetPricesFn = func(params *repository.Parameters) ([]*entity.Price, *repository.Meta, error) {
		prices := []*entity.Price{
			{
				UUID:              UUID,
				PassengerTypeUUID: "503f4ab8-5bf2-409c-a469-8da4b614232c",
				Price:             150.22,
			},
			{
				UUID:              UUID,
				PassengerTypeUUID: "503f4ab8-5bf2-409c-a469-8da4b614232c",
				Price:             150.22,
			},
		}
		meta := repository.NewMeta(params, int64(len(prices)))
		return prices, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/prices", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &pricesData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(pricesData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeletePrice_Success Test.
func TestDeletePrice_Success(t *testing.T) {
	var priceApp mock.PriceAppInterface
	priceHandler := NewPrices(&priceApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/price/:uuid", priceHandler.DeletePrice)

	priceApp.DeletePriceFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/price/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeletePrice_Failed_PriceNotFound Test.
func TestDeletePrice_Failed_PriceNotFound(t *testing.T) {
	var priceApp mock.PriceAppInterface
	priceHandler := NewPrices(&priceApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/price/:uuid", priceHandler.DeletePrice)

	priceApp.DeletePriceFn = func(UUID string) error {
		return exception.ErrorTextPriceNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/price/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
