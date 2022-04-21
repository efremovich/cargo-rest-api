package sityv1point00

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

// TestSaveSity_Success Test.
func TestSaveSity_Success(t *testing.T) {
	var sityData entity.Sity
	var sityApp mock.SityAppInterface
	sityHandler := NewSities(&sityApp)
	sityJSON := `{
		"name": "Самарканд",
		"region": "Самардкандская область",
		"latitude": "74.54",
		"longitude": "55.444"
	}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/sities", sityHandler.SaveSity)

	sityApp.SaveSityFn = func(sity *entity.Sity) (*entity.Sity, map[string]string, error) {
		return &entity.Sity{
			UUID:      UUID,
			Name:      "Самарканд",
			Region:    "Самаркандская область",
			Latitude:  "74.54",
			Longitude: "55.444",
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/sities", bytes.NewBufferString(sityJSON))
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &sityData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, sityData.UUID, UUID)
	assert.EqualValues(t, sityData.Name, "Самарканд")
	assert.EqualValues(t, sityData.Region, "Самаркандская область")
	assert.EqualValues(t, sityData.Latitude, "74.54")
	assert.EqualValues(t, sityData.Longitude, "55.444")
}

func TestSaveSity_InvalidData(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON:  `{"name":"", "region": ""}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"name": "", "region": "область",}`,
			statusCode: 422,
		},
	}

	for _, v := range samples {
		var sityApp mock.SityAppInterface
		sityHandler := NewSities(&sityApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/sities", sityHandler.SaveSity)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/sities", bytes.NewBufferString(v.inputJSON))
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

// TestUpdateSity_Success Test.
func TestUpdateSity_Success(t *testing.T) {
	var sityData entity.Sity
	var sityApp mock.SityAppInterface
	sityHandler := NewSities(&sityApp)
	sityJSON := `{
			"name": "Самарканд",
			"region": "Самардкандская область",
			"latitude": "74.54",
			"longitude": "55.444"
		  }`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/sities/:uuid", sityHandler.UpdateSity)

	sityApp.UpdateSityFn = func(UUID string, sity *entity.Sity) (*entity.Sity, map[string]string, error) {
		return &entity.Sity{
			UUID:      UUID,
			Name:      "Самарканд",
			Region:    "Самаркандская область",
			Latitude:  "74.54",
			Longitude: "55.444",
		}, nil, nil
	}

	sityApp.GetSityFn = func(string) (*entity.Sity, error) {
		return &entity.Sity{
			UUID:      UUID,
			Name:      "Самарканд",
			Region:    "Самаркандская область",
			Latitude:  "74.54",
			Longitude: "55.444",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPut, "/api/v1/external/sities/"+UUID, bytes.NewBufferString(sityJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &sityData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, sityData.UUID, UUID)
	assert.EqualValues(t, sityData.Name, "Самарканд")
	assert.EqualValues(t, sityData.Region, "Самаркандская область")
	assert.EqualValues(t, sityData.Latitude, "74.54")
	assert.EqualValues(t, sityData.Longitude, "55.444")
}

// TestGetSity_Success Test.
func TestGetSity_Success(t *testing.T) {
	var sityData entity.Sity
	var sityApp mock.SityAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	sityHandler := NewSities(&sityApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/sities/:uuid", sityHandler.GetSity)

	sityApp.GetSityFn = func(string) (*entity.Sity, error) {
		return &entity.Sity{
			UUID:      UUID,
			Name:      "Самарканд",
			Region:    "Самаркандская область",
			Latitude:  "74.54",
			Longitude: "55.444",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/sities/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &sityData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, sityData.UUID, UUID)
	assert.EqualValues(t, sityData.Name, "Самарканд")
	assert.EqualValues(t, sityData.Region, "Самаркандская область")
	assert.EqualValues(t, sityData.Latitude, "74.54")
	assert.EqualValues(t, sityData.Longitude, "55.444")
}

// TestGetSities_Success Test.
func TestGetSities_Success(t *testing.T) {
	var sityApp mock.SityAppInterface
	var sitiesData []entity.Sity
	var metaData repository.Meta
	sityHandler := NewSities(&sityApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/sities", sityHandler.GetSities)
	sityApp.GetSitiesFn = func(params *repository.Parameters) ([]*entity.Sity, *repository.Meta, error) {
		sities := []*entity.Sity{
			{
				UUID:      UUID,
				Name:      "Самара",
				Region:    "Самарская область",
				Latitude:  "33.44",
				Longitude: "65.568",
			},
			{
				UUID:      UUID,
				Name:      "Тверь",
				Region:    "Тверская область",
				Latitude:  "23.5488",
				Longitude: "35.456",
			},
		}
		meta := repository.NewMeta(params, int64(len(sities)))
		return sities, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/sities", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &sitiesData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(sitiesData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeleteSity_Success Test.
func TestDeleteSity_Success(t *testing.T) {
	var sityApp mock.SityAppInterface
	sityHandler := NewSities(&sityApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/sities/:uuid", sityHandler.DeleteSity)

	sityApp.DeleteSityFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/sities/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeleteSity_Failed_SityNotFound Test.
func TestDeleteSity_Failed_SityNotFound(t *testing.T) {
	var sityApp mock.SityAppInterface
	sityHandler := NewSities(&sityApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/sities/:uuid", sityHandler.DeleteSity)

	sityApp.DeleteSityFn = func(UUID string) error {
		return exception.ErrorTextSityNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/sities/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
