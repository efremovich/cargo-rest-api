package passengerv1point00

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

	"github.com/joho/godotenv"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestSavePassenger_Success Test.
func TestSavePassenger_Success(t *testing.T) {
	var passengerData entity.Passenger
	var passengerApp mock.PassengerAppInterface
	passengerHandler := NewPassengers(&passengerApp)
	passengerJSON := `{"first_name": "Владимир",
      "last_name": "Ульянов",
      "patronomic": "Ильич",
      "birthday": "1870-04-22T00:00:00Z",
      "passport_series": "0401",
      "passport_number": "564247",
      "user_uuid": "64f8b70d-d84f-4dde-a066-5dcb2f1f402a",
      "passenger_type_uuid":"7f3eb88e-98bd-4f5b-8a8c-34aaed1c7ffd"}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/passenger", passengerHandler.SavePassenger)

	passengerApp.SavePassengerFn = func(passenger *entity.Passenger) (*entity.Passenger, map[string]string, error) {
		return &entity.Passenger{
			UUID:              UUID,
			FirstName:         "Владимир",
			LastName:          "Ульянов",
			Patronomic:        "Ильич",
			BirthDay:          time.Date(1870, time.April, 22, 0, 0, 0, 0, time.UTC),
			PassportSeries:    "0401",
			PassportNumber:    "564247",
			UserUUID:          "c7912516-5bb4-4a5b-abee-d7bd242d93c6",
			PassengerTypeUUID: "503f4ab8-5bf2-409c-a469-8da4b614232c",
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPost,
		"/api/v1/external/passenger",
		bytes.NewBufferString(passengerJSON),
	)
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &passengerData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, passengerData.UUID, UUID)
	assert.EqualValues(t, passengerData.FirstName, "Владимир")
	assert.EqualValues(t, passengerData.LastName, "Ульянов")
	assert.EqualValues(t, passengerData.Patronomic, "Ильич")
	assert.EqualValues(t, passengerData.BirthDay, time.Date(1870, time.April, 22, 0, 0, 0, 0, time.UTC))
	assert.EqualValues(t, passengerData.PassportSeries, "0401")
	assert.EqualValues(t, passengerData.PassportNumber, "564247")
	assert.EqualValues(t, passengerData.UserUUID, "c7912516-5bb4-4a5b-abee-d7bd242d93c6")
	assert.EqualValues(t, passengerData.PassengerTypeUUID, "503f4ab8-5bf2-409c-a469-8da4b614232c")
}

func TestSavePassenger_InvalidData(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON: `{
          "first_name": "",
          "last_name": "",
          "patronoic": "",
          "birthday": "1890-02-22T00:00:00Z",
          "pasport_series": "0401",
          "pasport_number": "564247",
          "user_uuid": "",
          "passenger_type_uuid":"7f3eb88e-98bd-4f5b-8a8c-34aaed1c7ffd"
        }`,
			statusCode: 422,
		},
		{
			inputJSON: `{
          "first_name": 33,
          "last_name": 22,
          "patronoic": "",
          "birthday": "1890-02-22T00:00:00Z",
          "pasport_series": "0401",
          "pasport_number": "564247",
          "user_uuid": true,
          "passenger_type_uuid":"7f3eb88e-98bd-4f5b-8a8c-34aaed1c7ffd"
        }`,
			statusCode: 422,
		},
	}

	for _, v := range samples {
		var passengerApp mock.PassengerAppInterface
		passengerHandler := NewPassengers(&passengerApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/passenger", passengerHandler.SavePassenger)

		var err error
		c.Request, err = http.NewRequest(
			http.MethodPost,
			"/api/v1/external/passenger",
			bytes.NewBufferString(v.inputJSON),
		)
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

// TestUpdatePassenger_Success Test.
func TestUpdatePassenger_Success(t *testing.T) {
	var passengerData entity.Passenger
	var passengerApp mock.PassengerAppInterface
	passengerHandler := NewPassengers(&passengerApp)
	passengerJSON := `{
      "first_name": "Владимир",
      "last_name": "Ульянов",
      "patronoic": "Ильич",
      "birthday": "1890-02-22T00:00:00Z",
      "pasport_series": "0401",
      "pasport_number": "564247",
      "user_uuid": "64f8b70d-d84f-4dde-a066-5dcb2f1f402a",
      "passenger_type_uuid":"7f3eb88e-98bd-4f5b-8a8c-34aaed1c7ffd"
  }`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/passenger/:uuid", passengerHandler.UpdatePassenger)

	passengerApp.UpdatePassengerFn = func(UUID string, passenger *entity.Passenger) (*entity.Passenger, map[string]string, error) {
		return &entity.Passenger{
			UUID:              UUID,
			FirstName:         "Владимир",
			LastName:          "Ульянов",
			Patronomic:        "Ильич",
			BirthDay:          time.Date(1870, time.April, 22, 0, 0, 0, 0, time.UTC),
			PassportSeries:    "0401",
			PassportNumber:    "564247",
			UserUUID:          "c7912516-5bb4-4a5b-abee-d7bd242d93c6",
			PassengerTypeUUID: "503f4ab8-5bf2-409c-a469-8da4b614232c",
		}, nil, nil
	}

	passengerApp.GetPassengerFn = func(string) (*entity.Passenger, error) {
		return &entity.Passenger{
			UUID:              UUID,
			FirstName:         "Владимир",
			LastName:          "Ульянов",
			Patronomic:        "Ильич",
			BirthDay:          time.Date(1870, time.April, 22, 0, 0, 0, 0, time.UTC),
			PassportSeries:    "0401",
			PassportNumber:    "564247",
			UserUUID:          "c7912516-5bb4-4a5b-abee-d7bd242d93c6",
			PassengerTypeUUID: "503f4ab8-5bf2-409c-a469-8da4b614232c",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPut,
		"/api/v1/external/passenger/"+UUID,
		bytes.NewBufferString(passengerJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &passengerData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, passengerData.UUID, UUID)
	assert.EqualValues(t, passengerData.FirstName, "Владимир")
	assert.EqualValues(t, passengerData.LastName, "Ульянов")
	assert.EqualValues(t, passengerData.Patronomic, "Ильич")
	assert.EqualValues(t, passengerData.BirthDay, time.Date(1870, time.April, 22, 0, 0, 0, 0, time.UTC))
	assert.EqualValues(t, passengerData.PassportSeries, "0401")
	assert.EqualValues(t, passengerData.PassportNumber, "564247")
	assert.EqualValues(t, passengerData.UserUUID, "c7912516-5bb4-4a5b-abee-d7bd242d93c6")
	assert.EqualValues(t, passengerData.PassengerTypeUUID, "503f4ab8-5bf2-409c-a469-8da4b614232c")
}

// TestGetPassenger_Success Test.
func TestGetPassenger_Success(t *testing.T) {
	var passengerData entity.Passenger
	var passengerApp mock.PassengerAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	passengerHandler := NewPassengers(&passengerApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/passenger/:uuid", passengerHandler.GetPassenger)

	passengerApp.GetPassengerFn = func(string) (*entity.Passenger, error) {
		return &entity.Passenger{
			UUID:              UUID,
			FirstName:         "Владимир",
			LastName:          "Ульянов",
			Patronomic:        "Ильич",
			BirthDay:          time.Date(1870, time.April, 22, 0, 0, 0, 0, time.UTC),
			PassportSeries:    "0401",
			PassportNumber:    "564247",
			UserUUID:          "c7912516-5bb4-4a5b-abee-d7bd242d93c6",
			PassengerTypeUUID: "503f4ab8-5bf2-409c-a469-8da4b614232c",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/passenger/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &passengerData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, passengerData.UUID, UUID)
	assert.EqualValues(t, passengerData.FirstName, "Владимир")
	assert.EqualValues(t, passengerData.LastName, "Ульянов")
	assert.EqualValues(t, passengerData.Patronomic, "Ильич")
	assert.EqualValues(t, passengerData.BirthDay, time.Date(1870, time.April, 22, 0, 0, 0, 0, time.UTC))
	assert.EqualValues(t, passengerData.PassportSeries, "0401")
	assert.EqualValues(t, passengerData.PassportNumber, "564247")
	assert.EqualValues(t, passengerData.UserUUID, "c7912516-5bb4-4a5b-abee-d7bd242d93c6")
	assert.EqualValues(t, passengerData.PassengerTypeUUID, "503f4ab8-5bf2-409c-a469-8da4b614232c")
}

// TestGetPassengers_Success Test.
func TestGetPassengers_Success(t *testing.T) {
	var passengerApp mock.PassengerAppInterface
	var passengersData []entity.Passenger
	var metaData repository.Meta
	passengerHandler := NewPassengers(&passengerApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/passengers", passengerHandler.GetPassengers)
	passengerApp.GetPassengersFn = func(params *repository.Parameters) ([]*entity.Passenger, *repository.Meta, error) {
		passengers := []*entity.Passenger{
			{
				UUID:              UUID,
				FirstName:         "Владимир",
				LastName:          "Ульянов",
				Patronomic:        "Ильич",
				BirthDay:          time.Date(1870, time.April, 22, 0, 0, 0, 0, time.UTC),
				PassportSeries:    "0401",
				PassportNumber:    "564247",
				UserUUID:          "c7912516-5bb4-4a5b-abee-d7bd242d93c6",
				PassengerTypeUUID: "503f4ab8-5bf2-409c-a469-8da4b614232c",
			},
			{
				UUID:              UUID,
				FirstName:         "Николай",
				LastName:          "Чехидзе",
				Patronomic:        "Семёнович",
				BirthDay:          time.Date(1864, time.April, 9, 0, 0, 0, 0, time.UTC),
				PassportSeries:    "5501",
				PassportNumber:    "014247",
				UserUUID:          "c7912516-5bb4-4a5b-abee-d7bd242d93c6",
				PassengerTypeUUID: "04e9b29e-064b-4a13-8bab-074b14ae465d",
			},
		}
		meta := repository.NewMeta(params, int64(len(passengers)))
		return passengers, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/passengers", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &passengersData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(passengersData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeletePassenger_Success Test.
func TestDeletePassenger_Success(t *testing.T) {
	var passengerApp mock.PassengerAppInterface
	passengerHandler := NewPassengers(&passengerApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/passenger/:uuid", passengerHandler.DeletePassenger)

	passengerApp.DeletePassengerFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/passenger/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeletePassenger_Failed_PassengerNotFound Test.
func TestDeletePassenger_Failed_PassengerNotFound(t *testing.T) {
	var passengerApp mock.PassengerAppInterface
	passengerHandler := NewPassengers(&passengerApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/passenger/:uuid", passengerHandler.DeletePassenger)

	passengerApp.DeletePassengerFn = func(UUID string) error {
		return exception.ErrorTextPassengerNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/passenger/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
