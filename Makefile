# Makefile for Go Project

# Define variables
GO = go
BINARY_NAME = go-joke-w-service
BUILD_DIR = build

# Default target
all: build

# Build the binary
build:
	$(GO) mod download
    $(GO) mod tidy
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/main.go

# Build the Docker image
docker-build:
	docker build -t go-joke-w-service .

# Run Docker Compose
docker-up:
	docker-compose up

# Run tests
test:
	$(GO) test ./...

# Run the application
run:
	$(GO) run ./cmd/main.go

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)

# Install dependencies
install:
	$(GO) mod download
	$(GO) mod tidy
