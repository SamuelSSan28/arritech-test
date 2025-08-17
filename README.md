# Arritech User Management System

A fullstack web application for managing users with Vue.js frontend and Go backend, implementing CRUD operations with search and pagination functionality.

## Requirements Checklist

### Core Requirements
- ✅ MySQL database integration
- ✅ CRUD operations for users
- ✅ Search functionality
- ✅ Pagination
- ✅ Data validation and business rules

### Business Rules
- ✅ Email address must be unique
- ✅ Users must have an age greater than 18

### Bonus Tasks
- ✅ Request logging middleware
- ✅ Pagination functionality
- ✅ Dockerized backend application
- ✅ Clean architecture for testability

### Required Libraries
- ✅ Vue.js with Element Plus components (Table, Pagination)
- ✅ VeeValidate for form validation
- ✅ Go with Gin framework
- ✅ GORM for database operations
- ✅ Logrus for structured logging

## Quick Start

### Prerequisites
- Docker & Docker Compose
- Node.js 18+
- Go 1.21+

### Run with Docker

1. **Clone and start services**
   ```bash
   git clone <repository-url>
   cd arritech-test
   docker-compose up -d
   ```

2. **Run frontend**
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

3. **Access the application**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - API Documentation: http://localhost:8080/swagger/index.html

### Development Setup

#### Backend
```bash
cd backend
cp config.env.example .env
go mod download
go run cmd/server/main.go
```

#### Frontend
```bash
cd frontend
npm install
npm run dev
```

## API Endpoints

### Base URL: `http://localhost:8080/api/v1`

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users` | Get users list with pagination & search |
| GET | `/users/{id}` | Get user by ID |
| POST | `/users` | Create new user |
| PUT | `/users/{id}` | Update user |
| DELETE | `/users/{id}` | Delete user |

### Query Parameters
- `page`: Page number (default: 1)
- `per_page`: Items per page (default: 10, max: 100)
- `search`: Search term for name, email, or phone

``

## Technology Stack

- **Backend**: Go, Gin, GORM, MySQL
- **Frontend**: Vue.js 3, Element Plus, Vite
- **DevOps**: Docker, Docker Compose

## Testing

```bash
# Backend
cd backend && go test ./...

# Frontend
cd frontend && npm run test
```