package documentTypev1point00

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

// TestSaveDocumentType_Success Test.
func TestSaveDocumentType_Success(t *testing.T) {
	var documentTypeData entity.DocumentType
	var documentTypeApp mock.DocumentTypeAppInterface
	documentTypeHandler := NewDocumentTypes(&documentTypeApp)
	documentTypeJSON := `{
		"type": "Паспорт"
	}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.POST("/documentTypes", documentTypeHandler.SaveDocumentType)

	documentTypeApp.SaveDocumentTypeFn = func(documentType *entity.DocumentType) (*entity.DocumentType, map[string]string, error) {
		return &entity.DocumentType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/documentTypes", bytes.NewBufferString(documentTypeJSON))
	c.Request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &documentTypeData)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.EqualValues(t, documentTypeData.UUID, UUID)
	assert.EqualValues(t, documentTypeData.Type, "Паспорт")
}

func TestSaveDocumentType_InvalidData(t *testing.T) {
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
		var documentTypeApp mock.DocumentTypeAppInterface
		documentTypeHandler := NewDocumentTypes(&documentTypeApp)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		v1 := r.Group("/api/v1/external/")
		v1.POST("/documentTypes", documentTypeHandler.SaveDocumentType)

		var err error
		c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/documentTypes", bytes.NewBufferString(v.inputJSON))
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

// TestUpdateDocumentType_Success Test.
func TestUpdateDocumentType_Success(t *testing.T) {
	var documentTypeData entity.DocumentType
	var documentTypeApp mock.DocumentTypeAppInterface
	documentTypeHandler := NewDocumentTypes(&documentTypeApp)
	documentTypeJSON := `{"type": "Паспорт"}`
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.PUT("/documentTypes/:uuid", documentTypeHandler.UpdateDocumentType)

	documentTypeApp.UpdateDocumentTypeFn = func(UUID string, documentType *entity.DocumentType) (*entity.DocumentType, map[string]string, error) {
		return &entity.DocumentType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil, nil
	}

	documentTypeApp.GetDocumentTypeFn = func(string) (*entity.DocumentType, error) {
		return &entity.DocumentType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(
		http.MethodPut,
		"/api/v1/external/documentTypes/"+UUID,
		bytes.NewBufferString(documentTypeJSON),
	)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &documentTypeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, documentTypeData.UUID, UUID)
	assert.EqualValues(t, documentTypeData.Type, "Паспорт")
}

// TestGetDocumentType_Success Test.
func TestGetDocumentType_Success(t *testing.T) {
	var documentTypeData entity.DocumentType
	var documentTypeApp mock.DocumentTypeAppInterface

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	documentTypeHandler := NewDocumentTypes(&documentTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/documentTypes/:uuid", documentTypeHandler.GetDocumentType)

	documentTypeApp.GetDocumentTypeFn = func(string) (*entity.DocumentType, error) {
		return &entity.DocumentType{
			UUID: UUID,
			Type: "Паспорт",
		}, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/documentTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])

	_ = json.Unmarshal(data, &documentTypeData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, documentTypeData.UUID, UUID)
	assert.EqualValues(t, documentTypeData.Type, "Паспорт")
}

// TestGetDocumentTypes_Success Test.
func TestGetDocumentTypes_Success(t *testing.T) {
	var documentTypeApp mock.DocumentTypeAppInterface
	var documentTypesData []entity.DocumentType
	var metaData repository.Meta
	documentTypeHandler := NewDocumentTypes(&documentTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.GET("/documentTypes", documentTypeHandler.GetDocumentTypes)
	documentTypeApp.GetDocumentTypesFn = func(params *repository.Parameters) ([]*entity.DocumentType, *repository.Meta, error) {
		documentTypes := []*entity.DocumentType{
			{
				UUID: UUID,
				Type: "Паспорт",
			},
			{
				UUID: UUID,
				Type: "Свидетельство о рождении",
			},
		}
		meta := repository.NewMeta(params, int64(len(documentTypes)))
		return documentTypes, meta, nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/api/v1/external/documentTypes", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	response := encoder.ResponseDecoder(w.Body)
	data, _ := json.Marshal(response["data"])
	meta, _ := json.Marshal(response["meta"])

	_ = json.Unmarshal(data, &documentTypesData)
	_ = json.Unmarshal(meta, &metaData)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.EqualValues(t, 2, len(documentTypesData))
	assert.EqualValues(t, 1, metaData.Page)
	assert.EqualValues(t, 5, metaData.PerPage)
	assert.EqualValues(t, 2, metaData.Total)
}

// TestDeleteDocumentType_Success Test.
func TestDeleteDocumentType_Success(t *testing.T) {
	var documentTypeApp mock.DocumentTypeAppInterface
	documentTypeHandler := NewDocumentTypes(&documentTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/documentTypes/:uuid", documentTypeHandler.DeleteDocumentType)

	documentTypeApp.DeleteDocumentTypeFn = func(UUID string) error {
		return nil
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodDelete, "/api/v1/external/documentTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

// TestDeleteDocumentType_Failed_DocumentTypeNotFound Test.
func TestDeleteDocumentType_Failed_DocumentTypeNotFound(t *testing.T) {
	var documentTypeApp mock.DocumentTypeAppInterface
	documentTypeHandler := NewDocumentTypes(&documentTypeApp)
	UUID := uuid.New().String()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1/external/")
	v1.DELETE("/documentTypes/:uuid", documentTypeHandler.DeleteDocumentType)

	documentTypeApp.DeleteDocumentTypeFn = func(UUID string) error {
		return exception.ErrorTextDocumentTypeNotFound
	}

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/api/v1/external/documentTypes/"+UUID, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
