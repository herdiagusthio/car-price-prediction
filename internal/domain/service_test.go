package domain_test

import (
	"testing"

	"car-price-prediction/internal/domain"
)

// TestPredictionServiceInterface ensures that the PredictionService interface
// has the expected methods.
func TestPredictionServiceInterface(t *testing.T) {
	// This is a compile-time test to ensure that the interface has the expected methods.
	// If the interface changes, this test will fail to compile.
	var _ domain.PredictionService = (*mockPredictionService)(nil)
}

// mockPredictionService is a mock implementation of the PredictionService interface.
type mockPredictionService struct{}

// Predict implements the PredictionService interface.
func (m *mockPredictionService) Predict(input domain.UserInput) (*domain.PredictionResult, error) {
	return &domain.PredictionResult{
		PredictedPrice: 10000.0,
	}, nil
}