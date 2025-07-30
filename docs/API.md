# Car Price Prediction API Documentation

This document provides details about the API endpoints for the Car Price Prediction service.

---

## POST /predict

Predicts the price of a car based on its features.

### Request

**Headers**

| Header        | Value              |
|---------------|--------------------|
| Content-Type  | application/json   |

**Body**

The request body must be a JSON object containing the car's features.

**Fields:**

| Field            | Type    | Description                               | Required |
|------------------|---------|-------------------------------------------|----------|
| `symboling`      | integer | Its assigned insurance risk rating        | Yes      |
| `wheelbase`      | float   | The distance between the centers of the front and rear wheels | Yes      |
| `carlength`      | float   | The length of the car                     | Yes      |
| `carwidth`       | float   | The width of the car                      | Yes      |
| `carheight`      | float   | The height of the car                     | Yes      |
| `curbweight`     | integer | The weight of the car without occupants or baggage | Yes      |
| `enginesize`     | integer | The size of the engine                    | Yes      |
| `boreratio`      | float   | The ratio of the cylinder bore's diameter to its stroke length | Yes      |
| `stroke`         | float   | The distance the piston travels in the cylinder | Yes      |
| `compressionratio`| float   | The ratio of the volume of the cylinder and the combustion chamber when the piston is at the bottom of its stroke, to the volume when the piston is at the top of its stroke | Yes      |
| `horsepower`     | integer | The power of the engine                   | Yes      |
| `peakrpm`        | integer | The engine's peak revolutions per minute  | Yes      |
| `citympg`        | integer | Miles per gallon in the city              | Yes      |
| `highwaympg`     | integer | Miles per gallon on the highway           | Yes      |
| `fueltype`       | string  | The type of fuel the car uses (e.g., "gas", "diesel") | Yes      |
| `aspiration`     | string  | The type of aspiration (e.g., "std", "turbo") | Yes      |
| `doornumber`     | string  | The number of doors (e.g., "two", "four") | Yes      |
| `carbody`        | string  | The body style of the car (e.g., "sedan", "hatchback") | Yes      |
| `drivewheel`     | string  | The type of drive wheels (e.g., "fwd", "rwd", "4wd") | Yes      |
| `enginelocation` | string  | The location of the engine (e.g., "front", "rear") | Yes      |
| `enginetype`     | string  | The type of engine (e.g., "ohc", "dohc") | Yes      |
| `cylindernumber` | string  | The number of cylinders (e.g., "four", "six") | Yes      |
| `fuelsystem`     | string  | The fuel system type (e.g., "mpfi", "2bbl") | Yes      |
| `brand`          | string  | The brand of the car                      | Yes      |

**Example Request Body:**

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

### Responses

**Success Response (200 OK)**

Returns the predicted price of the car.

```json
{
    "predicted_price": 13495.0
}
```

**Error Responses**

*   **400 Bad Request**: Returned if the request body is malformed or missing required fields.
```json
{
    "error": "Invalid input data"
}
```
*   **500 Internal Server Error**: Returned if there is an issue with the prediction model or server.
```json
{
    "error": "Prediction failed"
}
```

### Example `curl` Command

```bash
cURL -X POST http://localhost:8080/predict \\
-H "Content-Type: application/json" \\
-d '{
    "Symboling": 3,
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