package tripv1point00

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

// TestSaveTrip_Success Test.
func TestSaveTrip_Success(t *testing.T) {
	var tripData entity.Trip
	var tripApp mock.TripAppInterface
	tripHandler := NewTrips(&tripApp)
	UUID := uuid.New().String()
	RouteUUID := uuid.New().String()
	VehicleUUID := uuid.New().String()
	DriverUUID := uuid.New().String()
	RegularityTypeUUID := uuid.New().String()

	tripJSON := `{
		"route_uuid": "` + RouteUUID + `",
    "vehicle_uuid":"` + VehicleUUID + `",
    "departure_time":"2022-04-22T11:00:00Z",
    "arravial_time":"2022-04-22T19:30:00Z",
    "regularityTipe_uuid":"` + RegularityTypeUUID + `",
    "driver_uuid":"` + DriverUUID + `"
	}`
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/trip", tripHandler.SaveTrip)

	tripApp.SaveTripFn = func(trip *entity.Trip) (*entity.Trip, map[string]string, error) {
		return &entity.Trip{
			UUID:               UUID,
			RouteUUID:          RouteUUID,
			VehicleUUID:        VehicleUUID,
			DepartureTime:      time.Date(2022, time.April, 22, 11, 0, 0, 0, time.UTC),
			ArravialTive:       time.Date(2022, time.April, 22, 19, 30, 0, 0, time.UTC),
			RegularityTypeUUID: RegularityTypeUUID,
			DriverUUID:         DriverUUID,
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPost,
		"/api/v1/external/trip",
		bytes.NewBufferString(tripJSON),
	)
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &tripData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, tripData.UUID, UUID)
	assert.EqualValues(t, tripData.RouteUUID, RouteUUID)
	assert.EqualValues(t, tripData.VehicleUUID, VehicleUUID)
	assert.EqualValues(t, tripData.RegularityTypeUUID, RegularityTypeUUID)
	assert.EqualValues(t, tripData.DepartureTime, time.Date(2022, time.April, 22, 11, 0, 0, 0, time.UTC))
	assert.EqualValues(t, tripData.ArravialTive, time.Date(2022, time.April, 22, 19, 30, 0, 0, time.UTC))
	assert.EqualValues(t, tripData.DriverUUID, DriverUUID)
}

func TestSaveTrip_InvalidData(t *testing.T) {
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
		var tripApp mock.TripAppInterface
		tripHandler := NewTrips(&tripApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/trip", tripHandler.SaveTrip)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/trip", bytes.NewBufferString(v.inputJSON))
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

// TestUpdateTrip_Success Test.
func TestUpdateTrip_Success(t *testing.T) {
	var tripData entity.Trip
	var tripApp mock.TripAppInterface
	tripHandler := NewTrips(&tripApp)

	UUID := uuid.New().String()
	RouteUUID := uuid.New().String()
	VehicleUUID := uuid.New().String()
	DriverUUID := uuid.New().String()
	RegularityTypeUUID := uuid.New().String()

	tripJSON := `{
		"uuid":"` + UUID + `",
		"route_uuid": "` + RouteUUID + `",
    "vehicle_uuid":"` + VehicleUUID + `",
    "departure_time":"2022-04-22T11:00:00Z",
    "arravial_time":"2022-04-22T19:30:00Z",
    "regularity_type_uuid":"` + RegularityTypeUUID + `",
    "driver_uuid":"` + DriverUUID + `"
	}`

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/trip/:uuid", tripHandler.UpdateTrip)

	tripApp.UpdateTripFn = func(UUID string, trip *entity.Trip) (*entity.Trip, map[string]string, error) {
		return &entity.Trip{
			UUID:               UUID,
			RouteUUID:          RouteUUID,
			VehicleUUID:        VehicleUUID,
			DepartureTime:      time.Date(2022, time.April, 22, 11, 0, 0, 0, time.UTC),
			ArravialTive:       time.Date(2022, time.April, 22, 19, 30, 0, 0, time.UTC),
			RegularityTypeUUID: RegularityTypeUUID,
			DriverUUID:         DriverUUID,
		}, nil, nil
	}

	tripApp.GetTripFn = func(string) (*entity.Trip, error) {
		return &entity.Trip{
			UUID:               UUID,
			RouteUUID:          RouteUUID,
			VehicleUUID:        VehicleUUID,
			DepartureTime:      time.Date(2022, time.April, 22, 11, 0, 0, 0, time.UTC),
			ArravialTive:       time.Date(2022, time.April, 22, 20, 22, 0, 0, time.UTC),
			RegularityTypeUUID: RegularityTypeUUID,
			DriverUUID:         DriverUUID,
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPut,
		"/api/v1/external/trip/"+UUID,
		bytes.NewBufferString(tripJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &tripData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, tripData.UUID, UUID)
	assert.EqualValues(t, tripData.RouteUUID, RouteUUID)
	assert.EqualValues(t, tripData.VehicleUUID, VehicleUUID)
	assert.EqualValues(t, tripData.RegularityTypeUUID, RegularityTypeUUID)
	assert.EqualValues(t, tripData.DepartureTime, time.Date(2022, time.April, 22, 11, 0, 0, 0, time.UTC))
	assert.EqualValues(t, tripData.ArravialTive, time.Date(2022, time.April, 22, 19, 30, 0, 0, time.UTC))
	assert.EqualValues(t, tripData.DriverUUID, DriverUUID)
}

// TestGetTrip_Success Test.
func TestGetTrip_Success(t *testing.T) {
	var tripData entity.Trip
	var tripApp mock.TripAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	tripHandler := NewTrips(&tripApp)
	UUID := uuid.New().String()
	RouteUUID := uuid.New().String()
	VehicleUUID := uuid.New().String()
	DriverUUID := uuid.New().String()
	RegularityTypeUUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/trip/:uuid", tripHandler.GetTrip)

	tripApp.GetTripFn = func(string) (*entity.Trip, error) {
		return &entity.Trip{
			UUID:               UUID,
			RouteUUID:          RouteUUID,
			VehicleUUID:        VehicleUUID,
			DepartureTime:      time.Date(2022, time.April, 22, 11, 0, 0, 0, time.UTC),
			ArravialTive:       time.Date(2022, time.April, 22, 19, 30, 0, 0, time.UTC),
			RegularityTypeUUID: RegularityTypeUUID,
			DriverUUID:         DriverUUID,
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/trip/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &tripData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, tripData.UUID, UUID)
	assert.EqualValues(t, tripData.RouteUUID, RouteUUID)
	assert.EqualValues(t, tripData.VehicleUUID, VehicleUUID)
	assert.EqualValues(t, tripData.RegularityTypeUUID, RegularityTypeUUID)
	assert.EqualValues(t, tripData.DepartureTime, time.Date(2022, time.April, 22, 11, 0, 0, 0, time.UTC))
	assert.EqualValues(t, tripData.ArravialTive, time.Date(2022, time.April, 22, 19, 30, 0, 0, time.UTC))
	assert.EqualValues(t, tripData.DriverUUID, DriverUUID)
}

// TestGetTrips_Success Test.
func TestGetTrips_Success(t *testing.T) {
	var tripApp mock.TripAppInterface
	var tripsData []entity.Trip
	var metaData repository.Meta
	tripHandler := NewTrips(&tripApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/trips", tripHandler.GetTrips)
	tripApp.GetTripsFn = func(params *repository.Parameters) ([]*entity.Trip, *repository.Meta, error) {
		trips := []*entity.Trip{
			{
				UUID:               UUID,
				RouteUUID:          uuid.New().String(),
				VehicleUUID:        uuid.New().String(),
				DepartureTime:      time.Date(2022, time.April, 22, 11, 0, 0, 0, time.UTC),
				ArravialTive:       time.Date(2022, time.April, 22, 19, 30, 0, 0, time.UTC),
				RegularityTypeUUID: uuid.New().String(),
				DriverUUID:         uuid.New().String(),
			},
			{
				UUID:               UUID,
				RouteUUID:          uuid.New().String(),
				VehicleUUID:        uuid.New().String(),
				DepartureTime:      time.Date(2022, time.April, 22, 11, 0, 0, 0, time.UTC),
				ArravialTive:       time.Date(2022, time.April, 22, 19, 30, 0, 0, time.UTC),
				RegularityTypeUUID: uuid.New().String(),
				DriverUUID:         uuid.New().String(),
			},
		}
		meta := repository.NewMeta(params, int64(len(trips)))
		return trips, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/trips", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &tripsData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(tripsData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeleteTrip_Success Test.
func TestDeleteTrip_Success(t *testing.T) {
	var tripApp mock.TripAppInterface
	tripHandler := NewTrips(&tripApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/trip/:uuid", tripHandler.DeleteTrip)

	tripApp.DeleteTripFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/trip/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeleteTrip_Failed_TripNotFound Test.
func TestDeleteTrip_Failed_TripNotFound(t *testing.T) {
	var tripApp mock.TripAppInterface
	tripHandler := NewTrips(&tripApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/trip/:uuid", tripHandler.DeleteTrip)

	tripApp.DeleteTripFn = func(UUID string) error {
		return exception.ErrorTextTripNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/trip/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
