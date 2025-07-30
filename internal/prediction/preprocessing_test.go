package prediction

import (
	"car-price-prediction/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform_Success(t *testing.T) {
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

	features, err := Transform(input)

	assert.NoError(t, err)
	assert.Len(t, features, ModelInputSize)

	// Check numerical features
	assert.Equal(t, float32(3), features[featureIndexMap["symboling"]])
	assert.Equal(t, float32(88.6), features[featureIndexMap["wheelbase"]])
	assert.Equal(t, float32(111), features[featureIndexMap["horsepower"]])

	// Check one-hot encoded features
	assert.Equal(t, float32(1.0), features[featureIndexMap["fueltype_gas"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["doornumber_two"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["drivewheel_rwd"]])
	assert.Equal(t, float32(0.0), features[featureIndexMap["drivewheel_fwd"]]) // Should not be set
}

func TestTransform_EdgeCases(t *testing.T) {
	// Test with zero values for numerical fields
	input := domain.UserInput{
		Symboling:        0,
		Wheelbase:        0,
		Carlength:        0,
		Carwidth:         0,
		Carheight:        0,
		Curbweight:       0,
		Enginesize:       0,
		Boreratio:        0,
		Stroke:           0,
		Compressionratio: 0,
		Horsepower:       0,
		Peakrpm:          0,
		Citympg:          0,
		Highwaympg:       0,
		Fueltype:         "diesel", // Not "gas", so fueltype_gas should be 0
		Aspiration:       "std",
		Doornumber:       "four", // Not "two", so doornumber_two should be 0
		Carbody:          "sedan",
		Drivewheel:       "fwd",
		Enginelocation:   "front",
		Enginetype:       "ohc",
		Cylindernumber:   "four",
		Fuelsystem:       "mpfi",
		Brand:            "toyota",
	}

	features, err := Transform(input)

	assert.NoError(t, err)
	assert.Len(t, features, ModelInputSize)

	// Check numerical features are all zero
	assert.Equal(t, float32(0), features[featureIndexMap["symboling"]])
	assert.Equal(t, float32(0), features[featureIndexMap["wheelbase"]])
	assert.Equal(t, float32(0), features[featureIndexMap["horsepower"]])

	// Check one-hot encoded features
	assert.Equal(t, float32(0.0), features[featureIndexMap["fueltype_gas"]])
	assert.Equal(t, float32(0.0), features[featureIndexMap["doornumber_two"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["drivewheel_fwd"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["carbody_sedan"]])
}

func TestTransform_CategoricalVariations(t *testing.T) {
	// Test with different categorical values
	input := domain.UserInput{
		// Only setting the categorical fields for this test
		Fueltype:       "gas",
		Aspiration:     "turbo",
		Doornumber:     "two",
		Carbody:        "sedan",
		Drivewheel:     "fwd",
		Enginelocation: "rear",
		Enginetype:     "ohc",
		Cylindernumber: "six",
		Fuelsystem:     "mpfi",
		Brand:          "toyota",
	}

	features, err := Transform(input)

	assert.NoError(t, err)
	assert.Len(t, features, ModelInputSize)

	// Check one-hot encoded features
	assert.Equal(t, float32(1.0), features[featureIndexMap["fueltype_gas"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["aspiration_turbo"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["doornumber_two"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["carbody_sedan"]])
	assert.Equal(t, float32(0.0), features[featureIndexMap["carbody_wagon"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["drivewheel_fwd"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["enginelocation_rear"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["enginetype_ohc"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["cylindernumber_six"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["fuelsystem_mpfi"]])
	assert.Equal(t, float32(1.0), features[featureIndexMap["brand_toyota"]])
}