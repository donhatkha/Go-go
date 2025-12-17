.PHONY: help build run test clean docker-build docker-up docker-down docker-logs install

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

build: ## Build the application
	@echo "Building application..."
	go build -o bin/gogo-app main.go

run: ## Run the application
	@echo "Running application..."
	go run main.go

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	go clean

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t gogo-app .

docker-up: ## Start Docker containers
	@echo "Starting Docker containers..."
	docker-compose up -d

docker-down: ## Stop Docker containers
	@echo "Stopping Docker containers..."
	docker-compose down

docker-logs: ## View Docker logs
	docker-compose logs -f

docker-restart: ## Restart Docker containers
	@echo "Restarting Docker containers..."
	docker-compose restart

dev: ## Run in development mode
	@echo "Starting development server..."
	go run main.go

format: ## Format code
	@echo "Formatting code..."
	go fmt ./...

lint: ## Lint code
	@echo "Linting code..."
	go vet ./...
