# Arritech User Management System

A fullstack web application for managing users with Vue.js frontend and Go backend, implementing CRUD operations with search and pagination functionality.

## 🎯 Project Overview

This project implements a user management system as requested in the Arritech Fullstack Developer evaluation. It provides a complete solution for managing user information with the following features:

### ✨ Features

- **User Management**: Complete CRUD operations (Create, Read, Update, Delete)
- **Search Functionality**: Search users by name, email, or phone number
- **Pagination**: Efficient data loading with customizable page sizes
- **Data Validation**: Both frontend and backend validation with meaningful error messages
- **Responsive Design**: Mobile-friendly interface using Element Plus components
- **Real-time Feedback**: Loading states, success/error messages, and confirmation dialogs
- **Business Rules**: Email uniqueness and age validation (must be > 18)

### 🏗️ Architecture

The project follows clean architecture principles with clear separation of concerns:

#### Backend (Go)
- **Clean Architecture**: Domain entities, repositories, services, and handlers
- **Repository Pattern**: Abstraction layer for data persistence (easily switchable between MySQL/MongoDB)
- **Dependency Injection**: Loose coupling between components
- **Middleware**: Request logging and CORS handling
- **Health Checks**: Monitoring endpoint for application status

#### Frontend (Vue.js)
- **Component-based Architecture**: Reusable and maintainable components
- **State Management**: Reactive state with Vue 3 Composition API
- **Service Layer**: Centralized API communication
- **Form Validation**: Client-side validation with VeeValidate
- **Responsive Design**: Mobile-first approach with Element Plus

## 📋 Requirements Fulfilled

### Core Requirements
- ✅ Vue.js frontend application
- ✅ Go backend with JSON API
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

## 🚀 Getting Started

### Prerequisites

- **Docker & Docker Compose**: For running the application
- **Node.js 18+**: For frontend development
- **Go 1.21+**: For backend development
- **MySQL 8.0+**: For local database (optional with Docker)

### Quick Start with Docker

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd arritech-test
   ```

2. **Start the backend services**
   ```bash
   docker-compose up -d
   ```

3. **Verify services are running**
   ```bash
   # Check backend health
   curl http://localhost:8080/health
   
   # Check MySQL connection
   docker-compose logs mysql
   ```

4. **Set up and run the frontend**
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

5. **Access the application**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - API Documentation: http://localhost:8080/health

### Development Setup

#### Backend Development

1. **Navigate to backend directory**
   ```bash
   cd backend
   ```

2. **Copy environment configuration**
   ```bash
   cp config.env.example .env
   ```

3. **Start MySQL database**
   ```bash
   docker run -d \
     --name mysql-dev \
     -e MYSQL_ROOT_PASSWORD=password \
     -e MYSQL_DATABASE=arritech_users \
     -p 3306:3306 \
     mysql:8.0
   ```

4. **Install dependencies and run**
   ```bash
   go mod download
   go run cmd/server/main.go
   ```

#### Frontend Development

1. **Navigate to frontend directory**
   ```bash
   cd frontend
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Start development server**
   ```bash
   npm run dev
   ```

4. **Build for production**
   ```bash
   npm run build
   ```

## 📖 API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Endpoints

#### Get Users List
```http
GET /users?page=1&per_page=10&search=query
```

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `per_page` (optional): Items per page (default: 10, max: 100)
- `search` (optional): Search term for name, email, or phone

**Response:**
```json
{
  "message": "Users retrieved successfully",
  "data": {
    "users": [...],
    "total": 100,
    "page": 1,
    "per_page": 10,
    "total_pages": 10
  }
}
```

#### Get User by ID
```http
GET /users/{id}
```

#### Create User
```http
POST /users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 25,
  "phone": "+1234567890",
  "address": "123 Main St"
}
```

#### Update User
```http
PUT /users/{id}
Content-Type: application/json

{
  "name": "John Smith",
  "age": 26
}
```

#### Delete User
```http
DELETE /users/{id}
```

### Error Responses
```json
{
  "error": "Error message",
  "details": {
    "field": "Validation error message"
  }
}
```

## 🧪 Testing

### Backend Tests
```bash
cd backend
go test ./...
```

### Frontend Tests
```bash
cd frontend
npm run test
```

### API Testing with curl

**Create a user:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 25,
    "phone": "+1234567890"
  }'
```

**Get users list:**
```bash
curl http://localhost:8080/api/v1/users?page=1&per_page=10
```

## 🔧 Configuration

### Backend Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | localhost | MySQL host |
| `DB_PORT` | 3306 | MySQL port |
| `DB_USER` | root | MySQL username |
| `DB_PASSWORD` | password | MySQL password |
| `DB_NAME` | arritech_users | Database name |
| `SERVER_PORT` | 8080 | API server port |
| `GIN_MODE` | debug | Gin mode (debug/release) |
| `LOG_LEVEL` | info | Log level (debug/info/warn/error) |
| `LOG_FORMAT` | json | Log format (json/text) |

### Frontend Configuration

The frontend is configured via `vite.config.js` and includes:
- API proxy to backend (localhost:8080)
- Build optimization
- Development server settings

## 📁 Project Structure

```
arritech-test/
├── backend/                 # Go backend application
│   ├── cmd/server/         # Application entry point
│   ├── internal/           # Private application code
│   │   ├── domain/         # Domain entities and interfaces
│   │   ├── handler/        # HTTP handlers
│   │   ├── repository/     # Data access layer
│   │   └── service/        # Business logic layer
│   ├── pkg/                # Public packages
│   │   ├── database/       # Database configuration
│   │   ├── logger/         # Logging setup
│   │   └── middleware/     # HTTP middleware
│   ├── scripts/            # Database scripts
│   ├── go.mod              # Go dependencies
│   └── Dockerfile          # Docker configuration
├── frontend/               # Vue.js frontend application
│   ├── src/
│   │   ├── components/     # Reusable components
│   │   ├── views/          # Page components
│   │   ├── services/       # API services
│   │   ├── router/         # Vue Router configuration
│   │   └── utils/          # Utility functions
│   ├── package.json        # Node.js dependencies
│   └── vite.config.js      # Vite configuration
├── docker-compose.yml      # Docker services orchestration
└── README.md              # Project documentation
```

## 🛠️ Technology Stack

### Backend
- **Go 1.21**: Programming language
- **Gin**: HTTP web framework
- **GORM**: ORM library
- **MySQL**: Database
- **Logrus**: Structured logging
- **Docker**: Containerization

### Frontend
- **Vue.js 3**: Progressive JavaScript framework
- **Element Plus**: UI component library
- **Vue Router**: Client-side routing
- **Axios**: HTTP client
- **VeeValidate**: Form validation
- **Vite**: Build tool and dev server

### DevOps
- **Docker & Docker Compose**: Container orchestration
- **MySQL 8.0**: Database server
- **Alpine Linux**: Lightweight container base

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the ISC License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

For support and questions, please contact the development team or create an issue in the repository.

---

**Developed for Arritech Fullstack Developer Evaluation** 