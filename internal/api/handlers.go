package api

import (
	"car-price-prediction/internal/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PredictHandler godoc
// @Summary Predict car price
// @Description Predict the price of a car based on its features.
// @Accept  json
// @Produce  json
// @Param   input     body    domain.UserInput   true        "Car Features"
// @Success 200 {object} domain.PredictionResult
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /predict [post]
func PredictHandler(service domain.PredictionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind the request body to a UserInput struct
		var input domain.UserInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
			return
		}

		// Call the prediction service
		result, err := service.Predict(input)
		if err != nil {
			log.Printf("Prediction error: %v", err) // Log the actual error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Prediction failed: " + err.Error()})
			return
		}

		// Return the prediction result
		c.JSON(http.StatusOK, result)
	}
}