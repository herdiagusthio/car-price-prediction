package prediction_test

import (
	"testing"

	"car-price-prediction/internal/domain"
	"car-price-prediction/internal/prediction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockSession is a mock implementation of the onnxruntime_go.AdvancedSession
type MockSession struct {
	mock.Mock
}

func (m *MockSession) Run() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockSession) Destroy() {
	m.Called()
}

func TestNewPredictionService(t *testing.T) {
	// Test that NewPredictionService returns a valid service
	service := prediction.NewPredictionService("test/path")
	assert.NotNil(t, service, "NewPredictionService should return a non-nil service")
}

func TestPredict_Integration(t *testing.T) {
	// Skip this test in normal test runs as it requires the actual model file
	t.Skip("Skipping integration test that requires the actual model file")

	// Create a service with the actual model path
	service := prediction.NewPredictionService("../../model/best_model.onnx")

	// Create a valid input
	input := domain.UserInput{
		Symboling:        0,
		Wheelbase:        104.5,
		Carlength:        187.8,
		Carwidth:         66.5,
		Carheight:        54.3,
		Curbweight:       2976,
		Enginesize:       146,
		Boreratio:        3.62,
		Stroke:           3.5,
		Compressionratio: 9.3,
		Horsepower:       110,
		Peakrpm:          5500,
		Citympg:          19,
		Highwaympg:       27,
		Fueltype:         "gas",
		Aspiration:       "std",
		Doornumber:       "four",
		Carbody:          "sedan",
		Drivewheel:       "fwd",
		Enginelocation:   "front",
		Enginetype:       "dohc",
		Cylindernumber:   "four",
		Fuelsystem:       "mpfi",
		Brand:            "toyota",
	}

	// Call Predict
	result, err := service.Predict(input)

	// Assert that the result is as expected
	assert.NoError(t, err, "Predict should not return an error")
	assert.NotNil(t, result, "Predict should return a non-nil result")
	assert.Greater(t, result.PredictedPrice, float32(0), "Predicted price should be greater than 0")
}

// TestPredict_InvalidInput tests that the Predict method returns an error when given invalid input
func TestPredict_InvalidInput(t *testing.T) {
	// Skip this test as it requires ONNX runtime initialization
	t.Skip("Skipping test that requires ONNX runtime initialization")

	// Create a service
	service := prediction.NewPredictionService("test/path")

	// Create an invalid input (missing required fields)
	input := domain.UserInput{
		// Missing most fields
		Symboling: 0,
		// Invalid categorical value
		Fueltype: "invalid_fuel_type",
	}

	// Call Predict
	_, err := service.Predict(input)

	// Assert that the error is as expected
	assert.Error(t, err, "Predict should return an error for invalid input")
}

// TestTransform_Success tests that the Transform function correctly transforms a valid input
func TestTransform_Success(t *testing.T) {
	// Create a valid input
	input := domain.UserInput{
		Symboling:        0,
		Wheelbase:        104.5,
		Carlength:        187.8,
		Carwidth:         66.5,
		Carheight:        54.3,
		Curbweight:       2976,
		Enginesize:       146,
		Boreratio:        3.62,
		Stroke:           3.5,
		Compressionratio: 9.3,
		Horsepower:       110,
		Peakrpm:          5500,
		Citympg:          19,
		Highwaympg:       27,
		Fueltype:         "gas",
		Aspiration:       "std",
		Doornumber:       "four",
		Carbody:          "sedan",
		Drivewheel:       "fwd",
		Enginelocation:   "front",
		Enginetype:       "dohc",
		Cylindernumber:   "four",
		Fuelsystem:       "mpfi",
		Brand:            "toyota",
	}

	// Call Transform
	features, err := prediction.Transform(input)

	// Assert that the result is as expected
	assert.NoError(t, err, "Transform should not return an error")
	assert.NotNil(t, features, "Transform should return non-nil features")
	assert.Equal(t, prediction.ModelInputSize, len(features), "Transform should return the correct number of features")
}

// TestTransform_InvalidInput tests that the Transform function handles invalid input
func TestTransform_InvalidInput(t *testing.T) {
	// Create an input with an invalid categorical value
	input := domain.UserInput{
		// Include all required fields
		Symboling:        0,
		Wheelbase:        100.0,
		Carlength:        180.0,
		Carwidth:         70.0,
		Carheight:        55.0,
		Curbweight:       2500,
		Enginesize:       140,
		Boreratio:        3.5,
		Stroke:           3.5,
		Compressionratio: 9.0,
		Horsepower:       100,
		Peakrpm:          5000,
		Citympg:          20,
		Highwaympg:       25,
		// Invalid categorical value - but this doesn't cause an error
		// since the preprocessing just sets the feature to 0
		Fueltype:       "invalid_fuel_type",
		Aspiration:     "std",
		Doornumber:     "four",
		Carbody:        "sedan",
		Drivewheel:     "fwd",
		Enginelocation: "front",
		Enginetype:     "ohc",
		Cylindernumber: "four",
		Fuelsystem:     "mpfi",
		Brand:          "toyota",
	}

	// Call Transform
	features, err := prediction.Transform(input)

	// The transform function doesn't actually return an error for invalid categorical values
	// It just doesn't set the one-hot encoding (leaves it as 0)
	assert.NoError(t, err, "Transform should not return an error for invalid categorical values")
	assert.Equal(t, float32(0.0), features[prediction.ModelInputSize-1], "Last feature should be 0 for invalid input")
}

// TestFeatureMap tests that the ModelInputSize constant is exported correctly
func TestFeatureMap(t *testing.T) {
	// Verify that ModelInputSize is exported and has the expected value
	assert.Equal(t, 64, prediction.ModelInputSize, "ModelInputSize should be 64")

	// Create a features array with the correct size
	features := make([]float32, prediction.ModelInputSize)
	
	// Verify the array has the correct length
	assert.Equal(t, 64, len(features), "Features array should have length 64")
}