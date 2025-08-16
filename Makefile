.PHONY: help test test-verbose test-coverage swagger swagger-force up down logs rebuild frontend-install frontend-build frontend-dev frontend-build-prod

# Default target
help:
	@echo "Available commands:"
	@echo ""
	@echo "Full Stack (Backend + Frontend + MySQL):"
	@echo "  up            - Start everything (backend + frontend + MySQL)"
	@echo "  down          - Stop everything"
	@echo "  logs          - Show backend logs"
	@echo "  rebuild       - Rebuild backend container"
	@echo ""
	@echo "Frontend:"
	@echo "  frontend-install  - Install frontend dependencies"
	@echo "  frontend-build    - Build frontend for production"
	@echo "  frontend-dev      - Start frontend dev server"
	@echo "  frontend-build-prod - Build frontend and copy to backend"
	@echo ""
	@echo "Testing:"
	@echo "  test          - Run backend tests"
	@echo "  test-verbose  - Run backend tests with verbose output"
	@echo "  test-coverage - Run backend tests with coverage report"
	@echo ""
	@echo "Swagger:"
	@echo "  swagger       - Generate Swagger documentation"
	@echo "  swagger-force - Force reinstall swag and generate docs"

# Full Stack (Backend + Frontend + MySQL)
up:
	@echo "Starting full stack development environment..."
	@echo "Backend: http://localhost:8080 (with hot reload)"
	@echo "Frontend: http://localhost:5173 (Vite dev server)"
	@echo "Swagger: http://localhost:8080/swagger/index.html"
	@echo "MySQL: localhost:3306"
	@echo ""
	@echo "Starting backend + MySQL..."
	docker-compose up -d
	@echo ""
	@echo "Starting frontend dev server..."
	@make frontend-dev

# Stop everything
down:
	@echo "Stopping full stack..."
	docker-compose down
	@echo "Full stack stopped"

# Show backend logs
logs:
	@echo "Showing backend logs..."
	docker-compose logs -f backend

# Rebuild backend container
rebuild:
	@echo "Rebuilding backend container..."
	docker-compose build backend
	docker-compose up -d

# Run backend tests
test:
	@echo "Running backend tests..."
	cd backend && go test ./...

# Run backend tests with verbose output
test-verbose:
	@echo "Running backend tests with verbose output..."
	cd backend && go test -v ./...

# Run backend tests with coverage
test-coverage:
	@echo "Running backend tests with coverage..."
	cd backend && go test -coverprofile=coverage.out ./...
	@echo "Coverage report generated: backend/coverage.out"
	cd backend && go tool cover -html=coverage.out -o coverage.html
	@echo "HTML coverage report generated: backend/coverage.html"

# Frontend commands
frontend-install:
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

frontend-build:
	@echo "Building frontend for production..."
	cd frontend && npm run build

frontend-dev:
	@echo "Starting frontend development server..."
	cd frontend && npm run dev

frontend-build-prod: frontend-build
	@echo "Building frontend and preparing for production..."
	@echo "Frontend built successfully!"

# Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	cd backend && if command -v swag > /dev/null; then \
		swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal; \
	else \
		echo "swag not found. Installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal; \
	fi
	@echo "Swagger documentation generated successfully!"

# Swagger documentation (force reinstall)
swagger-force:
	@echo "Force installing swag and generating Swagger documentation..."
	cd backend && go install github.com/swaggo/swag/cmd/swag@latest
	cd backend && swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
	@echo "Swagger documentation generated successfully!" 