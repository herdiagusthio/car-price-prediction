package prediction

import (
	"car-price-prediction/internal/domain"
	"fmt"

	onnx "github.com/yalue/onnxruntime_go"
)

// Ensure PredictionService implements domain.PredictionService interface
var _ domain.PredictionService = (*PredictionService)(nil)

// PredictionService encapsulates the ONNX model and prediction logic.
type PredictionService struct {
	modelPath string
}

// NewPredictionService creates a new prediction service with the given model path.
func NewPredictionService(modelPath string) *PredictionService {
	return &PredictionService{
		modelPath: modelPath,
	}
}

// Predict takes a UserInput, preprocesses it, runs the ONNX model, and returns a prediction result.
func (s *PredictionService) Predict(input domain.UserInput) (*domain.PredictionResult, error) {
	// Preprocess the input
	features, err := Transform(input)
	if err != nil {
		return nil, fmt.Errorf("preprocessing error: %w", err)
	}

	// Create a new input tensor with our data
	// The model expects a 2D tensor with shape [1, 64] (batch size of 1, 64 features)
	inputTensor, err := onnx.NewTensor[float32]([]int64{1, int64(ModelInputSize)}, features)
	if err != nil {
		return nil, fmt.Errorf("failed to create input tensor: %w", err)
	}
	defer inputTensor.Destroy()

	// Create an output tensor to hold the prediction result
	// The model outputs a single value representing the predicted price
	outputTensor, err := onnx.NewEmptyTensor[float32]([]int64{1, 1})
	if err != nil {
		return nil, fmt.Errorf("failed to create output tensor: %w", err)
	}
	defer outputTensor.Destroy()

	// Create a new session for this prediction using NewAdvancedSession
	// Use the correct tensor names from the model inspection
	session, err := onnx.NewAdvancedSession(
		s.modelPath,
		[]string{"float_input"}, // Correct input tensor name from model inspection
		[]string{"variable"}, // Correct output tensor name from model inspection
		[]onnx.ArbitraryTensor{inputTensor},
		[]onnx.ArbitraryTensor{outputTensor},
		nil, // Default options
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Destroy()

	// Run the model inference
	err = session.Run()
	if err != nil {
		return nil, fmt.Errorf("model inference error: %w", err)
	}

	// Extract the prediction from the output tensor
	outputData := outputTensor.GetData()
	if len(outputData) == 0 {
		return nil, fmt.Errorf("model produced no output")
	}

	// The first (and only) value in the output tensor is the predicted price
	predictedPrice := float32(outputData[0])

	return &domain.PredictionResult{
		PredictedPrice: predictedPrice,
	}, nil
}