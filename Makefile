.PHONY: help dev dev-backend dev-frontend build clean docker-up docker-down test

# Help command
help:
	@echo "Available commands:"
	@echo "  dev          - Start full development environment"
	@echo "  dev-backend  - Start backend development server"
	@echo "  dev-frontend - Start frontend development server"
	@echo "  build        - Build all applications"
	@echo "  docker-up    - Start Docker services"
	@echo "  docker-down  - Stop Docker services"
	@echo "  test         - Run all tests"
	@echo "  clean        - Clean build artifacts"
	@echo ""
	@echo "Database commands:"
	@echo "  db-migrate   - Run database migrations"
	@echo "  db-seed      - Seed database with sample data"
	@echo "  db-reset     - Reset database (remove volumes, migrate, seed)"

# Development commands
dev: docker-up dev-frontend	

dev-backend:
	@echo "Starting backend development server..."
	cd backend && go run cmd/server/main.go

dev-frontend:
	@echo "Starting frontend development server..."
	cd frontend && npm install && npm run dev

# Build commands
build: build-backend build-frontend

build-backend:
	@echo "Building backend..."
	cd backend && go build -o main cmd/server/main.go

build-frontend:
	@echo "Building frontend..."
	cd frontend && npm install && npm run build

# Docker commands
docker-up:
	@echo "Starting Docker services..."
	docker-compose up -d

docker-down:
	@echo "Stopping Docker services..."
	docker-compose down

docker-logs:
	@echo "Showing Docker logs..."
	docker-compose logs -f

# Test commands
test: test-backend test-frontend

test-backend:
	@echo "Running backend tests..."
	cd backend && go test ./...

test-frontend:
	@echo "Running frontend tests..."
	cd frontend && npm test

# Clean commands
clean:
	@echo "Cleaning build artifacts..."
	cd backend && rm -f main
	cd frontend && rm -rf dist
	docker-compose down --volumes --remove-orphans

# Database commands
db-migrate:
	@echo "Running database migrations..."
	cd backend && go run cmd/server/main.go --migrate-only

db-seed:
	@echo "Seeding database with sample data..."
	docker exec -i arritech-mysql mysql -u arritech -parritech123 arritech_users < backend/scripts/seed.sql

db-reset: docker-down
	@echo "Resetting database (removing volumes)..."
	docker-compose down --volumes --remove-orphans
	@echo "Starting fresh database..."
	docker-compose up -d
	@echo "Waiting for database to be ready..."
	sleep 10
	@echo "Running migrations..."
	$(MAKE) db-migrate
	@echo "Seeding database..."
	$(MAKE) db-seed
	@echo "Database reset complete!"

# Utility commands
install-deps:
	@echo "Installing backend dependencies..."
	cd backend && go mod download
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

format:
	@echo "Formatting code..."
	cd backend && go fmt ./...
	cd frontend && npm run lint

# Production commands
prod-build:
	@echo "Building for production..."
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml build

prod-up:
	@echo "Starting production services..."
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d 