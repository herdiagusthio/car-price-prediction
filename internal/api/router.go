package api

import (
	"car-price-prediction/internal/domain"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter configures the Gin router and defines the API endpoints.
func SetupRouter(service domain.PredictionService) *gin.Engine {
	// Create a new Gin router with default middleware.
	r := gin.Default()

	// Define the /predict endpoint.
	r.POST("/predict", PredictHandler(service))

	// Add Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}