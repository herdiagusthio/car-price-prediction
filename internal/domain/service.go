package domain

// PredictionService defines the interface for prediction services.
type PredictionService interface {
	// Predict takes a UserInput and returns a PredictionResult or an error.
	Predict(input UserInput) (*PredictionResult, error)
}