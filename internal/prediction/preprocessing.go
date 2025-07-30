package prediction

import (
	"car-price-prediction/internal/domain"
	"fmt"
	"strings"
)

// Transform converts a UserInput struct into a feature vector for the ONNX model.
// It handles both numerical features and categorical features (via one-hot encoding).
func Transform(input domain.UserInput) ([]float32, error) {
	// Initialize the feature vector with zeros
	features := make([]float32, ModelInputSize)

	// Set numerical features directly
	features[featureIndexMap["symboling"]] = float32(input.Symboling)
	features[featureIndexMap["wheelbase"]] = input.Wheelbase
	features[featureIndexMap["carlength"]] = input.Carlength
	features[featureIndexMap["carwidth"]] = input.Carwidth
	features[featureIndexMap["carheight"]] = input.Carheight
	features[featureIndexMap["curbweight"]] = float32(input.Curbweight)
	features[featureIndexMap["enginesize"]] = float32(input.Enginesize)
	features[featureIndexMap["boreratio"]] = input.Boreratio
	features[featureIndexMap["stroke"]] = input.Stroke
	features[featureIndexMap["compressionratio"]] = input.Compressionratio
	features[featureIndexMap["horsepower"]] = float32(input.Horsepower)
	features[featureIndexMap["peakrpm"]] = float32(input.Peakrpm)
	features[featureIndexMap["citympg"]] = float32(input.Citympg)
	features[featureIndexMap["highwaympg"]] = float32(input.Highwaympg)

	// Set categorical features using one-hot encoding
	// Fuel type
	if input.Fueltype == "gas" {
		setOneHot(features, "fueltype", input.Fueltype)
	}

	// Aspiration
	if input.Aspiration == "turbo" {
		setOneHot(features, "aspiration", input.Aspiration)
	}

	// Door number
	if input.Doornumber == "two" {
		setOneHot(features, "doornumber", input.Doornumber)
	}

	// Car body
	setOneHot(features, "carbody", input.Carbody)

	// Drive wheel
	setOneHot(features, "drivewheel", input.Drivewheel)

	// Engine location
	if input.Enginelocation == "rear" {
		setOneHot(features, "enginelocation", input.Enginelocation)
	}

	// Engine type
	setOneHot(features, "enginetype", input.Enginetype)

	// Cylinder number
	setOneHot(features, "cylindernumber", input.Cylindernumber)

	// Fuel system
	setOneHot(features, "fuelsystem", input.Fuelsystem)

	// Brand
	setOneHot(features, "brand", input.Brand)

	return features, nil
}

// setOneHot sets the appropriate one-hot encoded feature to 1.0.
// It constructs a feature key in the format "featureName_value" and sets that index to 1.0.
func setOneHot(features []float32, featureName, value string) {
	// Construct the feature key (e.g., "carbody_sedan")
	featureKey := fmt.Sprintf("%s_%s", featureName, strings.ToLower(value))

	// Check if this feature exists in our map
	if idx, exists := featureIndexMap[featureKey]; exists {
		features[idx] = 1.0
	}
	// If the feature doesn't exist in our map, it remains 0.0 (already initialized)
}