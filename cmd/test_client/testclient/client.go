// Package testclient provides functions to test the car price prediction API.
package testclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// UserInput represents the JSON request body for the prediction API.
// It contains all the car features needed for price prediction.
type UserInput struct {
	// Numerical features
	Symboling        int     `json:"symboling"`
	Wheelbase        float32 `json:"wheelbase"`
	Carlength        float32 `json:"carlength"`
	Carwidth         float32 `json:"carwidth"`
	Carheight        float32 `json:"carheight"`
	Curbweight       int     `json:"curbweight"`
	Enginesize       int     `json:"enginesize"`
	Boreratio        float32 `json:"boreratio"`
	Stroke           float32 `json:"stroke"`
	Compressionratio float32 `json:"compressionratio"`
	Horsepower       int     `json:"horsepower"`
	Peakrpm          int     `json:"peakrpm"`
	Citympg          int     `json:"citympg"`
	Highwaympg       int     `json:"highwaympg"`

	// Categorical features
	Fueltype       string `json:"fueltype"`
	Aspiration     string `json:"aspiration"`
	Doornumber     string `json:"doornumber"`
	Carbody        string `json:"carbody"`
	Drivewheel     string `json:"drivewheel"`
	Enginelocation string `json:"enginelocation"`
	Enginetype     string `json:"enginetype"`
	Cylindernumber string `json:"cylindernumber"`
	Fuelsystem     string `json:"fuelsystem"`
	Brand          string `json:"brand"`
}

// PredictionResult represents the JSON response body for the prediction API.
type PredictionResult struct {
	PredictedPrice float32 `json:"predicted_price"`
}

// ErrorResponse represents an error response from the API.
type ErrorResponse struct {
	Error string `json:"error"`
}

// MinimalClient sends a minimal request to the car price prediction API.
func MinimalClient() {
	// Create a minimal JSON payload with just the required fields
	payload := map[string]interface{}{
		"symboling":        0,
		"wheelbase":        104.5,
		"carlength":        187.8,
		"carwidth":         66.5,
		"carheight":        54.3,
		"curbweight":       2976,
		"enginesize":       146,
		"boreratio":        3.62,
		"stroke":           3.5,
		"compressionratio": 9.3,
		"horsepower":       110,
		"peakrpm":          5500,
		"citympg":          19,
		"highwaympg":       27,
		"fueltype":         "gas",
		"aspiration":       "std",
		"doornumber":       "four",
		"carbody":          "sedan",
		"drivewheel":       "fwd",
		"enginelocation":   "front",
		"enginetype":       "dohc",
		"cylindernumber":   "four",
		"fuelsystem":       "mpfi",
		"brand":            "toyota",
	}

	// Convert the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	// Print the JSON payload for debugging
	fmt.Printf("JSON payload: %s\n", string(jsonData))

	// Send the request
	resp, err := http.Post("http://localhost:8080/predict", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	// Print the response
	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Response body: %s\n", string(respBody))
}

// FullClient sends a full request to the car price prediction API.
func FullClient() {
	// Create a sample input
	input := UserInput{
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

	// Convert the input to JSON
	inputJSON, err := json.Marshal(input)
	if err != nil {
		fmt.Printf("Error marshaling input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Sending request with payload: %s\n", string(inputJSON))

	// Send the request to the API
	resp, err := http.Post("http://localhost:8080/predict", "application/json", bytes.NewBuffer(inputJSON))
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Response body: %s\n", string(respBody))

	// Parse the response based on the status code
	if resp.StatusCode == http.StatusOK {
		var result PredictionResult
		if err := json.Unmarshal(respBody, &result); err != nil {
			fmt.Printf("Error parsing prediction result: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Predicted price: $%.2f\n", result.PredictedPrice)
	} else {
		var errorResp ErrorResponse
		if err := json.Unmarshal(respBody, &errorResp); err != nil {
			fmt.Printf("Error parsing error response: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Error from API: %s\n", errorResp.Error)
	}
}