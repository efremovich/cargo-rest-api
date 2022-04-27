package driverv1point00

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

// TestSaveDriver_Success Test.
func TestSaveDriver_Success(t *testing.T) {
	var driverData entity.Driver
	var driverApp mock.DriverAppInterface
	driverHandler := NewDrivers(&driverApp)
	driverJSON := `{
		"name": "Варужан"
	}`
	UUID := uuid.New().String()
	UserUUID := uuid.New().String()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/drivers", driverHandler.SaveDriver)

	driverApp.SaveDriverFn = func(driver *entity.Driver) (*entity.Driver, map[string]string, error) {
		return &entity.Driver{
			UUID:     UUID,
			Name:     "Мамука",
			UserUUID: UserUUID,
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPost,
		"/api/v1/external/drivers",
		bytes.NewBufferString(driverJSON),
	)
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &driverData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, driverData.UUID, UUID)
	assert.EqualValues(t, driverData.UserUUID, UserUUID)
	assert.EqualValues(t, driverData.Name, "Мамука")
}

func TestSaveDriver_InvalidData(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON:  `{"name":33, "": ""}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"name": "", "": "jija",}`,
			statusCode: 422,
		},
	}

	for _, v := range samples {
		var driverApp mock.DriverAppInterface
		driverHandler := NewDrivers(&driverApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/drivers", driverHandler.SaveDriver)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/drivers", bytes.NewBufferString(v.inputJSON))
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

// TestUpdateDriver_Success Test.
func TestUpdateDriver_Success(t *testing.T) {
	var driverData entity.Driver
	var driverApp mock.DriverAppInterface
	driverHandler := NewDrivers(&driverApp)
	UUID := uuid.New().String()
	UserUUID := uuid.New().String()
	driverJSON := `{"name": "Мамука", "user_uuid": "` + UserUUID + `"}`

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/drivers/:uuid", driverHandler.UpdateDriver)

	driverApp.UpdateDriverFn = func(UUID string, driver *entity.Driver) (*entity.Driver, map[string]string, error) {
		return &entity.Driver{
			UUID:     UUID,
			UserUUID: UserUUID,
			Name:     "Мамука",
		}, nil, nil
	}

	driverApp.GetDriverFn = func(string) (*entity.Driver, error) {
		return &entity.Driver{
			UUID:     UUID,
			UserUUID: UserUUID,
			Name:     "Мамука",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPut,
		"/api/v1/external/drivers/"+UUID,
		bytes.NewBufferString(driverJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &driverData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, driverData.UUID, UUID)
	assert.EqualValues(t, driverData.UserUUID, UserUUID)
	assert.EqualValues(t, driverData.Name, "Мамука")
}

// TestGetDriver_Success Test.
func TestGetDriver_Success(t *testing.T) {
	var driverData entity.Driver
	var driverApp mock.DriverAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	driverHandler := NewDrivers(&driverApp)
	UUID := uuid.New().String()
	UserUUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/drivers/:uuid", driverHandler.GetDriver)

	driverApp.GetDriverFn = func(string) (*entity.Driver, error) {
		return &entity.Driver{
			UUID:     UUID,
			UserUUID: UserUUID,
			Name:     "Мамука",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/drivers/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &driverData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, driverData.UUID, UUID)
	assert.EqualValues(t, driverData.UserUUID, UserUUID)
	assert.EqualValues(t, driverData.Name, "Мамука")
}

// TestGetDrivers_Success Test.
func TestGetDrivers_Success(t *testing.T) {
	var driverApp mock.DriverAppInterface
	var driversData []entity.Driver
	var metaData repository.Meta
	driverHandler := NewDrivers(&driverApp)
	UUID := uuid.New().String()
	UserUUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/drivers", driverHandler.GetDrivers)
	driverApp.GetDriversFn = func(params *repository.Parameters) ([]*entity.Driver, *repository.Meta, error) {
		drivers := []*entity.Driver{
			{
				UUID:     UUID,
				UserUUID: UserUUID,
				Name:     "Мамука",
			},
			{
				UUID:     UUID,
				UserUUID: UserUUID,
				Name:     "Мамука",
			},
		}
		meta := repository.NewMeta(params, int64(len(drivers)))
		return drivers, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/drivers", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &driversData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(driversData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeleteDriver_Success Test.
func TestDeleteDriver_Success(t *testing.T) {
	var driverApp mock.DriverAppInterface
	driverHandler := NewDrivers(&driverApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/drivers/:uuid", driverHandler.DeleteDriver)

	driverApp.DeleteDriverFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/drivers/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeleteDriver_Failed_DriverNotFound Test.
func TestDeleteDriver_Failed_DriverNotFound(t *testing.T) {
	var driverApp mock.DriverAppInterface
	driverHandler := NewDrivers(&driverApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/drivers/:uuid", driverHandler.DeleteDriver)

	driverApp.DeleteDriverFn = func(UUID string) error {
		return exception.ErrorTextDriverNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/drivers/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
