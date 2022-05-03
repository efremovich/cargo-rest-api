package paymentv1point00

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

// TestSavePayment_Success Test.
func TestSavePayment_Success(t *testing.T) {
	var paymentData entity.Payment
	var paymentApp mock.PaymentAppInterface
	paymentHandler := NewPayments(&paymentApp)
	UUID := uuid.New().String()
	FromUUID := uuid.New().String()
	ToUUID := uuid.New().String()
	PriceUUID := uuid.New().String()
	PassengerTypeUUID := uuid.New().String()

	paymentJSON := `{
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
	v1.POST("/payment", paymentHandler.SavePayment)

	paymentApp.SavePaymentFn = func(payment *entity.Payment) (*entity.Payment, map[string]string, error) {
		return &entity.Payment{
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
		"/api/v1/external/payment",
		bytes.NewBufferString(paymentJSON),
	)
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &paymentData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, paymentData.UUID, UUID)
	assert.EqualValues(t, paymentData.FromUUID, FromUUID)
	assert.EqualValues(t, paymentData.ToUUID, ToUUID)
	assert.EqualValues(t, paymentData.Distance, 124)
	assert.EqualValues(t, paymentData.DistanceTime, 134)
}

func TestSavePayment_InvalidData(t *testing.T) {
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
		var paymentApp mock.PaymentAppInterface
		paymentHandler := NewPayments(&paymentApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/payment", paymentHandler.SavePayment)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/payment", bytes.NewBufferString(v.inputJSON))
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

// TestUpdatePayment_Success Test.
func TestUpdatePayment_Success(t *testing.T) {
	var paymentData entity.Payment
	var paymentApp mock.PaymentAppInterface
	paymentHandler := NewPayments(&paymentApp)

	UUID := uuid.New().String()
	FromUUID := uuid.New().String()
	ToUUID := uuid.New().String()
	PriceUUID := uuid.New().String()
	PassengerTypeUUID := uuid.New().String()

	paymentJSON := `{
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
	v1.PUT("/payment/:uuid", paymentHandler.UpdatePayment)

	paymentApp.UpdatePaymentFn = func(UUID string, payment *entity.Payment) (*entity.Payment, map[string]string, error) {
		return &entity.Payment{
			UUID:         UUID,
			FromUUID:     FromUUID,
			ToUUID:       ToUUID,
			Distance:     124,
			DistanceTime: 134,
		}, nil, nil
	}

	paymentApp.GetPaymentFn = func(string) (*entity.Payment, error) {
		return &entity.Payment{
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
		"/api/v1/external/payment/"+UUID,
		bytes.NewBufferString(paymentJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &paymentData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, paymentData.UUID, UUID)
	assert.EqualValues(t, paymentData.FromUUID, FromUUID)
	assert.EqualValues(t, paymentData.ToUUID, ToUUID)
	assert.EqualValues(t, paymentData.Distance, 124)
	assert.EqualValues(t, paymentData.DistanceTime, 134)
}

// TestGetPayment_Success Test.
func TestGetPayment_Success(t *testing.T) {
	var paymentData entity.Payment
	var paymentApp mock.PaymentAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	paymentHandler := NewPayments(&paymentApp)
	UUID := uuid.New().String()
	FromUUID := uuid.New().String()
	ToUUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/payment/:uuid", paymentHandler.GetPayment)

	paymentApp.GetPaymentFn = func(string) (*entity.Payment, error) {
		return &entity.Payment{
			UUID:         UUID,
			FromUUID:     FromUUID,
			ToUUID:       ToUUID,
			Distance:     124,
			DistanceTime: 134,
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/payment/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &paymentData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, paymentData.UUID, UUID)
	assert.EqualValues(t, paymentData.FromUUID, FromUUID)
	assert.EqualValues(t, paymentData.ToUUID, ToUUID)
	assert.EqualValues(t, paymentData.Distance, 124)
	assert.EqualValues(t, paymentData.DistanceTime, 134)
}

// TestGetPayments_Success Test.
func TestGetPayments_Success(t *testing.T) {
	var paymentApp mock.PaymentAppInterface
	var paymentsData []entity.Payment
	var metaData repository.Meta
	paymentHandler := NewPayments(&paymentApp)
	UUID := uuid.New().String()
	FromUUID := uuid.New().String()
	ToUUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/payments", paymentHandler.GetPayments)
	paymentApp.GetPaymentsFn = func(params *repository.Parameters) ([]*entity.Payment, *repository.Meta, error) {
		payments := []*entity.Payment{
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
		meta := repository.NewMeta(params, int64(len(payments)))
		return payments, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/payments", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &paymentsData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(paymentsData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeletePayment_Success Test.
func TestDeletePayment_Success(t *testing.T) {
	var paymentApp mock.PaymentAppInterface
	paymentHandler := NewPayments(&paymentApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/payment/:uuid", paymentHandler.DeletePayment)

	paymentApp.DeletePaymentFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/payment/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeletePayment_Failed_PaymentNotFound Test.
func TestDeletePayment_Failed_PaymentNotFound(t *testing.T) {
	var paymentApp mock.PaymentAppInterface
	paymentHandler := NewPayments(&paymentApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/payment/:uuid", paymentHandler.DeletePayment)

	paymentApp.DeletePaymentFn = func(UUID string) error {
		return exception.ErrorTextPaymentNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/payment/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestPayments_AddPaymentPrice(t *testing.T) {
	var paymentData entity.Payment
	var paymentApp mock.PaymentAppInterface
	paymentHandler := NewPayments(&paymentApp)

	UUID := uuid.New().String()
	priceUUID := uuid.New().String()

	paymentJSON := `{"UUID":"` + UUID + `", "prices":[{"uuid":"` + priceUUID + `"}]}`

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/payment/price_add/", paymentHandler.AddPaymentPrice)

	paymentApp.AddPaymentPriceFn = func(payment *entity.Payment) (*entity.Payment, map[string]string, error) {
		return &entity.Payment{
			UUID:   UUID,
			Prices: []*entity.Price{{UUID: priceUUID}},
		}, nil, nil
	}

	paymentApp.GetPaymentFn = func(string) (*entity.Payment, error) {
		return &entity.Payment{
			UUID:   UUID,
			Prices: []*entity.Price{{UUID: priceUUID}},
		}, nil
	}
	var err error

	c.Request, err = http.NewRequest(
		http.MethodPost,
		"/api/v1/external/payment/price_add/",
		bytes.NewBufferString(paymentJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &paymentData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, paymentData.UUID, UUID)

}

func TestPayments_DeletePaymentPrice(t *testing.T) {
	var paymentData entity.Payment
	var paymentApp mock.PaymentAppInterface
	paymentHandler := NewPayments(&paymentApp)

	UUID := uuid.New().String()
	priceUUID := uuid.New().String()

	paymentJSON := `{"UUID":"` + UUID + `", "prices":[{"uuid":"` + priceUUID + `"}]}`

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/payment/price_del/", paymentHandler.DeletePaymentPrice)

	paymentApp.DeletePaymentPriceFn = func(payment *entity.Payment) (*entity.Payment, map[string]string, error) {
		return &entity.Payment{
			UUID:   UUID,
			Prices: []*entity.Price{{UUID: priceUUID}},
		}, nil, nil
	}

	paymentApp.GetPaymentFn = func(string) (*entity.Payment, error) {
		return &entity.Payment{
			UUID:   UUID,
			Prices: []*entity.Price{{UUID: priceUUID}},
		}, nil
	}
	var err error

	c.Request, err = http.NewRequest(
		http.MethodPost,
		"/api/v1/external/payment/price_del/",
		bytes.NewBufferString(paymentJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &paymentData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, paymentData.UUID, UUID)

}
