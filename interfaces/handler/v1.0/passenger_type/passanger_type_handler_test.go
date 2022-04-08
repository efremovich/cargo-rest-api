package passengerTypev1point00

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

// TestSavePassengerType_Success Test.
func TestSavePassengerType_Success(t *testing.T) {
	var passengerTypeData entity.PassengerType
	var passengerTypeApp mock.PassengerTypeAppInterface
	passengerTypeHandler := NewPassengerTypes(&passengerTypeApp)
	passengerTypeJSON := `{
		"type": "Взрослый",
	}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/passengerTypes", passengerTypeHandler.SavePassengerType)

	passengerTypeApp.SavePassengerTypeFn = func(passengerType *entity.PassengerType) (*entity.PassengerType, map[string]string, error) {
		return &entity.PassengerType{
			UUID: UUID,
			Type: "Взрослый",
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/passengerTypes", bytes.NewBufferString(passengerTypeJSON))
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &passengerTypeData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, passengerTypeData.UUID, UUID)
	assert.EqualValues(t, passengerTypeData.Type, "Взрослый")
}

func TestSavePassengerType_InvalidData(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON:  `{"type":33, "": ""}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"type": "", "": "область",}`,
			statusCode: 422,
		},
	}

	for _, v := range samples {
		var passengerTypeApp mock.PassengerTypeAppInterface
		passengerTypeHandler := NewPassengerTypes(&passengerTypeApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/passengerTypes", passengerTypeHandler.SavePassengerType)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/passengerTypes", bytes.NewBufferString(v.inputJSON))
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

// TestUpdatePassengerType_Success Test.
func TestUpdatePassengerType_Success(t *testing.T) {
	var passengerTypeData entity.PassengerType
	var passengerTypeApp mock.PassengerTypeAppInterface
	passengerTypeHandler := NewPassengerTypes(&passengerTypeApp)
	passengerTypeJSON := `{"type": "Взрослый"}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/passengerTypes/:uuid", passengerTypeHandler.UpdatePassengerType)

	passengerTypeApp.UpdatePassengerTypeFn = func(UUID string, passengerType *entity.PassengerType) (*entity.PassengerType, map[string]string, error) {
		return &entity.PassengerType{
			UUID: UUID,
			Type: "Взрослый",
		}, nil, nil
	}

	passengerTypeApp.GetPassengerTypeFn = func(string) (*entity.PassengerType, error) {
		return &entity.PassengerType{
			UUID: UUID,
			Type: "Взрослый",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPut, "/api/v1/external/passengerTypes/"+UUID, bytes.NewBufferString(passengerTypeJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &passengerTypeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, passengerTypeData.UUID, UUID)
	assert.EqualValues(t, passengerTypeData.Type, "Взрослый")
}

// TestGetPassengerType_Success Test.
func TestGetPassengerType_Success(t *testing.T) {
	var passengerTypeData entity.PassengerType
	var passengerTypeApp mock.PassengerTypeAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	passengerTypeHandler := NewPassengerTypes(&passengerTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/passengerTypes/:uuid", passengerTypeHandler.GetPassengerType)

	passengerTypeApp.GetPassengerTypeFn = func(string) (*entity.PassengerType, error) {
		return &entity.PassengerType{
			UUID: UUID,
			Type: "Взрослый",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/passengerTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &passengerTypeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, passengerTypeData.UUID, UUID)
	assert.EqualValues(t, passengerTypeData.Type, "Взрослый")
}

// TestGetPassengerTypes_Success Test.
func TestGetPassengerTypes_Success(t *testing.T) {
	var passengerTypeApp mock.PassengerTypeAppInterface
	var passengerTypesData []entity.PassengerType
	var metaData repository.Meta
	passengerTypeHandler := NewPassengerTypes(&passengerTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/passengerTypes", passengerTypeHandler.GetPassengerTypes)
	passengerTypeApp.GetPassengerTypesFn = func(params *repository.Parameters) ([]*entity.PassengerType, *repository.Meta, error) {
		passengerTypes := []*entity.PassengerType{
			{
				UUID: UUID,
				Type: "Взрослый",
			},
			{
				UUID: UUID,
				Type: "Детсткий",
			},
		}
		meta := repository.NewMeta(params, int64(len(passengerTypes)))
		return passengerTypes, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/passengerTypes", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &passengerTypesData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(passengerTypesData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeletePassengerType_Success Test.
func TestDeletePassengerType_Success(t *testing.T) {
	var passengerTypeApp mock.PassengerTypeAppInterface
	passengerTypeHandler := NewPassengerTypes(&passengerTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/passengerTypes/:uuid", passengerTypeHandler.DeletePassengerType)

	passengerTypeApp.DeletePassengerTypeFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/passengerTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeletePassengerType_Failed_PassengerTypeNotFound Test.
func TestDeletePassengerType_Failed_PassengerTypeNotFound(t *testing.T) {
	var passengerTypeApp mock.PassengerTypeAppInterface
	passengerTypeHandler := NewPassengerTypes(&passengerTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/passengerTypes/:uuid", passengerTypeHandler.DeletePassengerType)

	passengerTypeApp.DeletePassengerTypeFn = func(UUID string) error {
		return exception.ErrorTextPassengerTypeNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/passengerTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
