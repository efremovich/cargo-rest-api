package regularityTypev1point00

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

// TestSaveRegularityType_Success Test.
func TestSaveRegularityType_Success(t *testing.T) {
	var regularityTypeData entity.RegularityType
	var regularityTypeApp mock.RegularityTypeAppInterface
	regularityTypeHandler := NewRegularityTypes(&regularityTypeApp)
	regularityTypeJSON := `{
		"type": "Паспорт"
	}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/regularityTypes", regularityTypeHandler.SaveRegularityType)

	regularityTypeApp.SaveRegularityTypeFn = func(regularityType *entity.RegularityType) (*entity.RegularityType, map[string]string, error) {
		return &entity.RegularityType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPost,
		"/api/v1/external/regularityTypes",
		bytes.NewBufferString(regularityTypeJSON),
	)
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &regularityTypeData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, regularityTypeData.UUID, UUID)
	assert.EqualValues(t, regularityTypeData.Type, "Паспорт")
}

func TestSaveRegularityType_InvalidData(t *testing.T) {
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
		var regularityTypeApp mock.RegularityTypeAppInterface
		regularityTypeHandler := NewRegularityTypes(&regularityTypeApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/regularityTypes", regularityTypeHandler.SaveRegularityType)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/regularityTypes", bytes.NewBufferString(v.inputJSON))
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

// TestUpdateRegularityType_Success Test.
func TestUpdateRegularityType_Success(t *testing.T) {
	var regularityTypeData entity.RegularityType
	var regularityTypeApp mock.RegularityTypeAppInterface
	regularityTypeHandler := NewRegularityTypes(&regularityTypeApp)
	regularityTypeJSON := `{"type": "Паспорт"}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/regularityTypes/:uuid", regularityTypeHandler.UpdateRegularityType)

	regularityTypeApp.UpdateRegularityTypeFn = func(UUID string, regularityType *entity.RegularityType) (*entity.RegularityType, map[string]string, error) {
		return &entity.RegularityType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil, nil
	}

	regularityTypeApp.GetRegularityTypeFn = func(string) (*entity.RegularityType, error) {
		return &entity.RegularityType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPut,
		"/api/v1/external/regularityTypes/"+UUID,
		bytes.NewBufferString(regularityTypeJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &regularityTypeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, regularityTypeData.UUID, UUID)
	assert.EqualValues(t, regularityTypeData.Type, "Паспорт")
}

// TestGetRegularityType_Success Test.
func TestGetRegularityType_Success(t *testing.T) {
	var regularityTypeData entity.RegularityType
	var regularityTypeApp mock.RegularityTypeAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	regularityTypeHandler := NewRegularityTypes(&regularityTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/regularityTypes/:uuid", regularityTypeHandler.GetRegularityType)

	regularityTypeApp.GetRegularityTypeFn = func(string) (*entity.RegularityType, error) {
		return &entity.RegularityType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/regularityTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &regularityTypeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, regularityTypeData.UUID, UUID)
	assert.EqualValues(t, regularityTypeData.Type, "Паспорт")
}

// TestGetRegularityTypes_Success Test.
func TestGetRegularityTypes_Success(t *testing.T) {
	var regularityTypeApp mock.RegularityTypeAppInterface
	var regularityTypesData []entity.RegularityType
	var metaData repository.Meta
	regularityTypeHandler := NewRegularityTypes(&regularityTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/regularityTypes", regularityTypeHandler.GetRegularityTypes)
	regularityTypeApp.GetRegularityTypesFn = func(params *repository.Parameters) ([]*entity.RegularityType, *repository.Meta, error) {
		regularityTypes := []*entity.RegularityType{
			{
				UUID: UUID,
				Type: "Паспорт",
			},
			{
				UUID: UUID,
				Type: "Свидетельство о рождении",
			},
		}
		meta := repository.NewMeta(params, int64(len(regularityTypes)))
		return regularityTypes, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/regularityTypes", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &regularityTypesData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(regularityTypesData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeleteRegularityType_Success Test.
func TestDeleteRegularityType_Success(t *testing.T) {
	var regularityTypeApp mock.RegularityTypeAppInterface
	regularityTypeHandler := NewRegularityTypes(&regularityTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/regularityTypes/:uuid", regularityTypeHandler.DeleteRegularityType)

	regularityTypeApp.DeleteRegularityTypeFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/regularityTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeleteRegularityType_Failed_RegularityTypeNotFound Test.
func TestDeleteRegularityType_Failed_RegularityTypeNotFound(t *testing.T) {
	var regularityTypeApp mock.RegularityTypeAppInterface
	regularityTypeHandler := NewRegularityTypes(&regularityTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/regularityTypes/:uuid", regularityTypeHandler.DeleteRegularityType)

	regularityTypeApp.DeleteRegularityTypeFn = func(UUID string) error {
		return exception.ErrorTextRegularityTypeNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/regularityTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
