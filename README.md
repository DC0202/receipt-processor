
# Receipt Processor

The **Receipt Processor** is a Go-based application designed to process receipts, calculate points based on specific criteria, and provide APIs to retrieve receipt details and health status. The application includes backend validation to ensure data integrity and prevents duplicate processing of receipts.

---

## Features

- **Receipt Processing**:
  - Processes receipt data and calculates reward points.
  - Validates all receipt fields against strict criteria.
  - Rejects any receipt data with invalid or extra fields.
  - Prevents duplicate receipt processing using hashing.
- **APIs**:
  - `/receipts/process`: Process a receipt (POST).
  - `/receipts/{id}/points`: Retrieve points for a receipt, with optional detailed explanation (GET).
  - `/health`: Health check endpoint (GET).
- **Structured Logging**:
  - Advanced logging with configurable log levels (`DEBUG`, `INFO`, `WARN`, `ERROR`).
- **Environment Configurations**:
  - Configurable via `.env` file or environment variables.
- **Docker Support**:
  - Fully containerized for easy deployment.

---

## Table of Contents

1. [Requirements](#requirements)
2. [Setup Instructions](#setup-instructions)
3. [Environment Variables](#environment-variables)
4. [Running Locally](#running-locally)
5. [Running with Docker](#running-with-docker)
6. [APIs and Usage](#apis-and-usage)
7. [Receipt Validation Rules](#receipt-validation-rules)
8. [Preventing Duplicate Receipts](#preventing-duplicate-receipts)
9. [Testing](#testing)
10. [Logging](#logging)
11. [Deployment](#deployment)
12. [License](#license)

---

## Requirements

- Go `v1.23+`
- Docker (optional for containerized deployment)
- `make` (optional for easier commands)

---

## Setup Instructions

### Clone the Repository

```bash
git clone <repository-url>
cd receipt-processor
```

### Install Dependencies

```bash
go mod tidy
```

### Environment Variables

The application uses a `.env` file for configuration. Below are the required variables:

```env
APP_PORT=8000       # Port to run the application
LOG_LEVEL=error     # Log level (debug, info, warn, error)
DATABASE_URL=localhost # Database URL (if applicable)
```

Make sure to copy the `.env` file into the root of your project.

### Running Locally

Start the Application:

```bash
go run main.go
```

Access the APIs:

- Process a receipt: POST `/receipts/process`
- Get points: GET `/receipts/{id}/points`
- Health check: GET `/health`

### Running with Docker

Build the Docker Image:

```bash
docker build -t receipt-processor .
```

Run the Docker Container:

```bash
docker run -p 8080:8080 --env-file .env receipt-processor
```

Access APIs:

Use `http://localhost:8080/` for APIs while the container is running.

---

## APIs and Usage

### 1. Process Receipt (POST `/receipts/process`)

Description: This endpoint processes a receipt and returns a unique receipt ID. Receipt data is validated against predefined rules, and duplicate receipts are identified using hashing.

#### Request:

```json
{
  "retailer": "Target",
  "purchaseDate": "2024-11-24",
  "purchaseTime": "14:00",
  "items": [
    {
      "shortDescription": "Shampoo",
      "price": "5.99"
    },
    {
      "shortDescription": "Conditioner",
      "price": "6.49"
    }
  ],
  "total": "12.48"
}
```

#### Response:

```json
{
  "id": "receipt12345"
}
```

### 2. Get Points (GET `/receipts/{id}/points`)

Description: This endpoint retrieves the points for a specific receipt. An optional query parameter `detailed` can be used to return a detailed explanation of how the points were calculated.

#### Path Parameters:

- `id`: The unique receipt ID generated during the `/receipts/process` call.

#### Query Parameters:

- `detailed` (optional): If set to true, the response includes a detailed explanation of the points calculation.

#### Request:

```bash
GET /receipts/{id}/points?detailed=true
```

#### Response Without `detailed`:

```json
{
  "points": 100
}
```

#### Response With `detailed`:

```json
{
  "points": 100,
  "explanation": "Breakdown:
50 points - total is a round dollar amount
25 points - total is a multiple of 0.25
..."
}
```

### 3. Health Check (GET `/health`)

Description: This endpoint checks the health status of the application and returns a simple `OK` response.

#### Request:

```bash
GET /health
```

#### Response:

```json
"OK"
```

---

## Receipt Validation Rules

The application validates receipt data using the following rules:

- **Retailer Name**:
  - Must be alphanumeric, spaces, dashes (-), or ampersands (&).
  - Invalid retailer names will result in an error.
- **Items**:
  - Each item must have a `shortDescription` and a `price`.
  - `shortDescription` must contain only alphanumeric characters, spaces, or dashes.
  - `price` must be a valid decimal with two places and must not be negative.
- **Purchase Date**:
  - Must follow the YYYY-MM-DD format.
  - Cannot be a future date.
- **Purchase Time**:
  - Must follow the HH:MM format (24-hour clock).
- **Total**:
  - Must be a valid decimal with two places.
  - Must be non-negative.
  - The sum of all item prices must match the total.
- **No Extra Fields**:
  - The receipt data must only include retailer, purchaseDate, purchaseTime, items, and total.
  - Any additional fields will cause validation to fail.

### Backend Validation

- All fields are validated using regex patterns and additional logic for correctness.

#### Error Response Example:

```json
{
  "error": "Invalid retailer name: Retailer field is required and must be alphanumeric."
}
```

---

## Preventing Duplicate Receipts

To prevent duplicate receipt processing, the application uses hashing:

- **SHA-1 Hashing**:
  - The backend generates a unique SHA-1 hash for each receipt based on its stringified data.
  - This ensures that even slight changes in the data (e.g., different total) result in a different hash.

### Deduplication

- If the hash already exists in the receipt store, the server responds with the existing receipt ID without reprocessing.

#### Response for Duplicate Receipt:

```json
{
  "id": "receipt12345"
}
```

---

## Testing

To run the tests:

```bash
go test ./...
```

---

## Logging

Logs are configured using logrus and are written to both `app.log` and the console. The log level can be configured via the `LOG_LEVEL` environment variable.

### Log Levels:

- `DEBUG`: Detailed debugging logs.
- `INFO`: General application information.
- `WARN`: Potential issues.
- `ERROR`: Errors that need attention.

---

## Deployment

### Local Deployment

- Ensure `go` is installed.
- Run `go run main.go` to start the application.

### Docker Deployment

- Build the Docker image: `docker build -t receipt-processor .`
- Run the container: `docker run -p 8080:8080 --env-file .env receipt-processor`

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.
