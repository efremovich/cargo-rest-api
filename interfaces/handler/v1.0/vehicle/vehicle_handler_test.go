package vehiclev1point00

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

// TestSaveVehicle_Success Test.
func TestSaveVehicle_Success(t *testing.T) {
	var vehicleData entity.Vehicle
	var vehicleApp mock.VehicleAppInterface
	vehicleHandler := NewVehicles(&vehicleApp)
	vehicleJSON := `{
		"model": "Шевроле каптюр",
		"reg_code": "x554ыв95",
		"number_of_seats": 5,
		"class": "Люкс"
	}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/vehicles", vehicleHandler.SaveVehicle)

	vehicleApp.SaveVehicleFn = func(vehicle *entity.Vehicle) (*entity.Vehicle, map[string]string, error) {
		return &entity.Vehicle{
			UUID:          UUID,
			Model:         "Шевроле каптюр",
			RegCode:       "x554ыв95",
			NumberOfSeats: 5,
			Class:         "Люкс",
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/vehicles", bytes.NewBufferString(vehicleJSON))
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &vehicleData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, vehicleData.UUID, UUID)
	assert.EqualValues(t, vehicleData.Model, "Шевроле каптюр")
	assert.EqualValues(t, vehicleData.RegCode, "x554ыв95")
	assert.EqualValues(t, vehicleData.NumberOfSeats, 5)
	assert.EqualValues(t, vehicleData.Class, "Люкс")
}

func TestSaveVehicle_InvalidData(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON:  `{"model":"BMV", "": ""}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"model": "", "": "область",}`,
			statusCode: 422,
		},
	}

	for _, v := range samples {
		var vehicleApp mock.VehicleAppInterface
		vehicleHandler := NewVehicles(&vehicleApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/vehicles", vehicleHandler.SaveVehicle)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/vehicles", bytes.NewBufferString(v.inputJSON))
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

// TestUpdateVehicle_Success Test.
func TestUpdateVehicle_Success(t *testing.T) {
	var vehicleData entity.Vehicle
	var vehicleApp mock.VehicleAppInterface
	vehicleHandler := NewVehicles(&vehicleApp)
	vehicleJSON := `{
		"model": "Шевроле каптюр",
		"reg_code": "x554ыв95",
		"number_of_seats": 5,
		"class": "Люкс"
		  }`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/vehicles/:uuid", vehicleHandler.UpdateVehicle)

	vehicleApp.UpdateVehicleFn = func(UUID string, vehicle *entity.Vehicle) (*entity.Vehicle, map[string]string, error) {
		return &entity.Vehicle{
			UUID:          UUID,
			Model:         "Шевроле каптюр",
			RegCode:       "x554ыв95",
			NumberOfSeats: 5,
			Class:         "Люкс",
		}, nil, nil
	}

	vehicleApp.GetVehicleFn = func(string) (*entity.Vehicle, error) {
		return &entity.Vehicle{
			UUID:          UUID,
			Model:         "Шевроле каптюр",
			RegCode:       "x554ыв95",
			NumberOfSeats: 5,
			Class:         "Люкс",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPut, "/api/v1/external/vehicles/"+UUID, bytes.NewBufferString(vehicleJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &vehicleData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, vehicleData.UUID, UUID)
	assert.EqualValues(t, vehicleData.Model, "Шевроле каптюр")
	assert.EqualValues(t, vehicleData.RegCode, "x554ыв95")
	assert.EqualValues(t, vehicleData.NumberOfSeats, 5)
	assert.EqualValues(t, vehicleData.Class, "Люкс")
}

// TestGetVehicle_Success Test.
func TestGetVehicle_Success(t *testing.T) {
	var vehicleData entity.Vehicle
	var vehicleApp mock.VehicleAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	vehicleHandler := NewVehicles(&vehicleApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/vehicles/:uuid", vehicleHandler.GetVehicle)

	vehicleApp.GetVehicleFn = func(string) (*entity.Vehicle, error) {
		return &entity.Vehicle{
			UUID:          UUID,
			Model:         "Шевроле каптюр",
			RegCode:       "x554ыв95",
			NumberOfSeats: 5,
			Class:         "Люкс",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/vehicles/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &vehicleData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, vehicleData.UUID, UUID)
	assert.EqualValues(t, vehicleData.Model, "Шевроле каптюр")
	assert.EqualValues(t, vehicleData.RegCode, "x554ыв95")
	assert.EqualValues(t, vehicleData.NumberOfSeats, "5")
	assert.EqualValues(t, vehicleData.Class, "Люкс")
}

// TestGetVehicles_Success Test.
func TestGetVehicles_Success(t *testing.T) {
	var vehicleApp mock.VehicleAppInterface
	var vehiclesData []entity.Vehicle
	var metaData repository.Meta
	vehicleHandler := NewVehicles(&vehicleApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/vehicles", vehicleHandler.GetVehicles)
	vehicleApp.GetVehiclesFn = func(params *repository.Parameters) ([]*entity.Vehicle, *repository.Meta, error) {
		vehicles := []*entity.Vehicle{
			{
				UUID:          UUID,
				Model:         "Шевроле каптюр",
				RegCode:       "x554ыв95",
				NumberOfSeats: 5,
				Class:         "Люкс",
			},
			{
				Model:         "BMV",
				RegCode:       "x115ыв954",
				NumberOfSeats: 12,
				Class:         "Премиум",
			},
		}
		meta := repository.NewMeta(params, int64(len(vehicles)))
		return vehicles, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/vehicles", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &vehiclesData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(vehiclesData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeleteVehicle_Success Test.
func TestDeleteVehicle_Success(t *testing.T) {
	var vehicleApp mock.VehicleAppInterface
	vehicleHandler := NewVehicles(&vehicleApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/vehicles/:uuid", vehicleHandler.DeleteVehicle)

	vehicleApp.DeleteVehicleFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/vehicles/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeleteVehicle_Failed_VehicleNotFound Test.
func TestDeleteVehicle_Failed_VehicleNotFound(t *testing.T) {
	var vehicleApp mock.VehicleAppInterface
	vehicleHandler := NewVehicles(&vehicleApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/vehicles/:uuid", vehicleHandler.DeleteVehicle)

	vehicleApp.DeleteVehicleFn = func(UUID string) error {
		return exception.ErrorTextVehicleNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/vehicles/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
