.PHONY: help test test-verbose test-coverage build clean run

# Default target
help:
	@echo "Available commands:"
	@echo "  test          - Run all tests"
	@echo "  test-verbose  - Run tests with verbose output"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  build         - Build the application"
	@echo "  clean         - Clean build artifacts"
	@echo "  run           - Run the application"
	@echo "  deps          - Download dependencies"
	@echo "  fmt           - Format code with gofmt"
	@echo "  lint          - Run golangci-lint"
	@echo "  vet           - Run go vet"

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Run tests
test: deps
	@echo "Running tests..."
	go test ./...

# Run tests with verbose output
test-verbose: deps
	@echo "Running tests with verbose output..."
	go test -v ./...

# Run tests with coverage
test-coverage: deps
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	@echo "Coverage report generated: coverage.out"
	@go tool cover -html=coverage.out -o coverage.html
	@echo "HTML coverage report generated: coverage.html"

# Build the application
build: deps
	@echo "Building application..."
	go build -o main cmd/server/main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f main
	rm -f coverage.out
	rm -f coverage.html

# Run the application
run: build
	@echo "Running application..."
	./main

# Install test dependencies
test-deps:
	@echo "Installing test dependencies..."
	go get github.com/stretchr/testify/assert
	go get github.com/stretchr/testify/mock
	go get github.com/stretchr/testify/require
	go get github.com/DATA-DOG/go-sqlmock

# Run specific test package
test-package:
	@echo "Usage: make test-package PKG=package_name"
	@echo "Example: make test-package PKG=./internal/domain/entity"
	@if [ -z "$(PKG)" ]; then echo "PKG variable is required"; exit 1; fi
	go test -v $(PKG)

# Run tests with race detection
test-race: deps
	@echo "Running tests with race detection..."
	go test -race ./...

# Run benchmarks
bench: deps
	@echo "Running benchmarks..."
	go test -bench=. ./...

# Check for security vulnerabilities
security: deps
	@echo "Checking for security vulnerabilities..."
	go list -json -deps ./... | nancy sleuth

# Generate mocks (if using mockery)
mocks:
	@echo "Generating mocks..."
	@if command -v mockery > /dev/null; then \
		mockery --all --keeptree --output=./mocks; \
	else \
		echo "mockery not found. Install with: go install github.com/vektra/mockery/v2@latest"; \
	fi

# Show test coverage in terminal
test-coverage-term: deps
	@echo "Running tests with coverage..."
	go test -cover ./...

# Run integration tests
test-integration: deps
	@echo "Running integration tests..."
	go test -tags=integration ./...

# Run unit tests only
test-unit: deps
	@echo "Running unit tests only..."
	go test -tags=unit ./...

# Pre-commit checks
pre-commit: fmt vet test
	@echo "Pre-commit checks completed successfully!"

# CI pipeline
ci: deps fmt vet test test-coverage
	@echo "CI pipeline completed successfully!" 