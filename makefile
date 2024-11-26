# Variables
APP_NAME=receipt-processor
BUILD_DIR=bin
DOCKER_IMAGE=$(APP_NAME)

# Default goal
.DEFAULT_GOAL := help

# Targets

## Build the Go application
build:
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd

## Run the application locally
run:
	@echo "Running the application..."
	cd cmd && go run main.go

## Clean the build directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

## Run tests
test:
	@echo "Running tests..."
	go test ./internal/config
	go test ./internal/handler
	go test ./internal/model
	go test ./internal/services
	go test ./internal/utility
	go test ./pkg/hash

## Build a Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t receipt-processor .

## Run the Docker container
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 receipt-processor

## Remove the Docker image
docker-clean:
	@echo "Removing Docker image..."
	docker rmi receipt-processor

## Display available commands
help:
	@echo "Available make targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":[^#]*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

# Add descriptions to commands for `make help`
build: ## Build the Go application
run: ## Run the application locally
clean: ## Clean the build directory
test: ## Run tests
docker-build: ## Build a Docker image
docker-run: ## Run the Docker container
docker-clean: ## Remove the Docker image
help: ## Display available commands
