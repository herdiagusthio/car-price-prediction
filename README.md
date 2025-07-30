# Car Price Prediction API

A high-performance, production-ready Go API that predicts car prices using a pre-trained ONNX model.

## Features

- RESTful API built with Gin framework
- Clean Architecture for maintainability and testability
- ONNX Runtime for efficient model inference
- Comprehensive test suite with high code coverage
- Input validation and error handling

## Prerequisites

- Go 1.18 or higher
- ONNX Runtime library
  - macOS: `brew install onnxruntime`
  - Linux: `sudo apt-get install libonnxruntime`

## Project Structure

```
/
├── cmd/
│   └── api/            # Entry point, server initialization
├── internal/
│   ├── api/            # Gin handlers, routing, and middleware
│   ├── domain/         # Core business objects (structs)
│   ├── prediction/     # Business logic for prediction
│   └── config/         # Configuration loading
├── model/
│   └── best_model.onnx # The ONNX model file
└── docs/
    └── development/    # Development documentation
```

## Getting Started

### Building the Project

```bash
# Clone the repository
git clone https://github.com/yourusername/car-price-prediction.git
cd car-price-prediction

# Build the project
go build -o car-price-api ./cmd/api
```

### Running the API

```bash
# Run the API server
./car-price-api
```

The API will start on port 8080 by default.

## API Usage

### Predicting Car Price

**Endpoint:** `POST /predict`

**Request Body:**

```json
{
    "symboling": 3,
    "wheelbase": 88.6,
    "carlength": 168.8,
    "carwidth": 64.1,
    "carheight": 48.8,
    "curbweight": 2548,
    "enginesize": 130,
    "boreratio": 3.47,
    "stroke": 2.68,
    "compressionratio": 9.0,
    "horsepower": 111,
    "peakrpm": 5000,
    "citympg": 21,
    "highwaympg": 27,
    "fueltype": "gas",
    "aspiration": "std",
    "doornumber": "two",
    "carbody": "convertible",
    "drivewheel": "rwd",
    "enginelocation": "front",
    "enginetype": "dohc",
    "cylindernumber": "four",
    "fuelsystem": "mpfi",
    "brand": "alfa-romero"
}
```

**Success Response (200 OK):**

```json
{
    "predicted_price": 13495.50
}
```

**Example using curl:**

```bash
curl -X POST http://localhost:8080/predict \
  -H "Content-Type: application/json" \
  -d '{
    "symboling": 3,
    "wheelbase": 88.6,
    "carlength": 168.8,
    "carwidth": 64.1,
    "carheight": 48.8,
    "curbweight": 2548,
    "enginesize": 130,
    "boreratio": 3.47,
    "stroke": 2.68,
    "compressionratio": 9.0,
    "horsepower": 111,
    "peakrpm": 5000,
    "citympg": 21,
    "highwaympg": 27,
    "fueltype": "gas",
    "aspiration": "std",
    "doornumber": "two",
    "carbody": "convertible",
    "drivewheel": "rwd",
    "enginelocation": "front",
    "enginetype": "dohc",
    "cylindernumber": "four",
    "fuelsystem": "mpfi",
    "brand": "alfa-romero"
}'
```

## Running Tests

```bash
# Run all tests with coverage
go test -cover ./...

# Generate a coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.