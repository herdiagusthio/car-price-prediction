package api

import (
	"bytes"
	"car-price-prediction/internal/domain"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// mockPredictionService is a mock implementation of the prediction service for testing.
type mockPredictionService struct{}

// Predict implements the prediction service interface for testing.
func (m *mockPredictionService) Predict(input domain.UserInput) (*domain.PredictionResult, error) {
	// Return a fixed prediction for testing
	return &domain.PredictionResult{PredictedPrice: 15000.0}, nil
}

// setupTestServer initializes a test server with a mock prediction service.
func setupTestServer() *httptest.Server {
	// Use a mock prediction service for testing
	mockService := &mockPredictionService{}

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a router with the mock service
	router := SetupRouter(mockService)

	return httptest.NewServer(router)
}

func TestPredictHandler_Success(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// Create a valid input
	input := domain.UserInput{
		Symboling:        3,
		Wheelbase:        88.6,
		Carlength:        168.8,
		Carwidth:         64.1,
		Carheight:        48.8,
		Curbweight:       2548,
		Enginesize:       130,
		Boreratio:        3.47,
		Stroke:           2.68,
		Compressionratio: 9.0,
		Horsepower:       111,
		Peakrpm:          5000,
		Citympg:          21,
		Highwaympg:       27,
		Fueltype:         "gas",
		Aspiration:       "std",
		Doornumber:       "two",
		Carbody:          "convertible",
		Drivewheel:       "rwd",
		Enginelocation:   "front",
		Enginetype:       "dohc",
		Cylindernumber:   "four",
		Fuelsystem:       "mpfi",
		Brand:            "alfa-romero",
	}

	// Convert to JSON
	body, _ := json.Marshal(input)

	// Send the request
	resp, err := http.Post(server.URL+"/predict", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Parse the response
	var result domain.PredictionResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)

	// Check the result
	assert.Equal(t, float32(15000.0), result.PredictedPrice)
}

func TestPredictHandler_BadRequest_InvalidJSON(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// Send an invalid JSON
	body := []byte(`{"invalid":json}`)

	// Send the request
	resp, err := http.Post(server.URL+"/predict", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestPredictHandler_BadRequest_MissingFields(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// Create an input with missing required fields
	input := map[string]interface{}{
		"symboling": 3,
		// Missing all other required fields
	}

	// Convert to JSON
	body, _ := json.Marshal(input)

	// Send the request
	resp, err := http.Post(server.URL+"/predict", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}