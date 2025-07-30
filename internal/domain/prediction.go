package domain

// UserInput represents the JSON request body for the prediction API.
// It contains all the car features needed for price prediction.
type UserInput struct {
	// Numerical features
	Symboling        int     `json:"Symboling"` // Removed binding:"required"
	Wheelbase        float32 `json:"wheelbase" binding:"required"`
	Carlength        float32 `json:"carlength" binding:"required"`
	Carwidth         float32 `json:"carwidth" binding:"required"`
	Carheight        float32 `json:"carheight" binding:"required"`
	Curbweight       int     `json:"curbweight" binding:"required"`
	Enginesize       int     `json:"enginesize" binding:"required"`
	Boreratio        float32 `json:"boreratio" binding:"required"`
	Stroke           float32 `json:"stroke" binding:"required"`
	Compressionratio float32 `json:"compressionratio" binding:"required"`
	Horsepower       int     `json:"horsepower" binding:"required"`
	Peakrpm          int     `json:"peakrpm" binding:"required"`
	Citympg          int     `json:"citympg" binding:"required"`
	Highwaympg       int     `json:"highwaympg" binding:"required"`

	// Categorical features
	Fueltype       string `json:"fueltype" binding:"required"`
	Aspiration     string `json:"aspiration" binding:"required"`
	Doornumber     string `json:"doornumber" binding:"required"`
	Carbody        string `json:"carbody" binding:"required"`
	Drivewheel     string `json:"drivewheel" binding:"required"`
	Enginelocation string `json:"enginelocation" binding:"required"`
	Enginetype     string `json:"enginetype" binding:"required"`
	Cylindernumber string `json:"cylindernumber" binding:"required"`
	Fuelsystem     string `json:"fuelsystem" binding:"required"`
	Brand          string `json:"brand" binding:"required"`
}

// PredictionResult represents the JSON response body for the prediction API.
type PredictionResult struct {
	PredictedPrice float32 `json:"predicted_price"`
}

// ErrorResponse represents the JSON response body for an error.
type ErrorResponse struct {
	Error string `json:"error"`
}