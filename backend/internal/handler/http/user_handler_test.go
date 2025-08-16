package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"arritech-user-management/internal/domain/entity"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of the UserService interface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, req entity.CreateUserRequest) (*entity.User, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserService) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserService) ListUsers(ctx context.Context, params entity.UserSearchParams) (*entity.UserListResponse, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UserListResponse), args.Error(1)
}

func (m *MockUserService) UpdateUser(ctx context.Context, id uint, req entity.UpdateUserRequest) (*entity.User, error) {
	args := m.Called(ctx, id, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func setupTestHandler() (*UserHandler, *MockUserService) {
	mockService := &MockUserService{}
	validator := validator.New()

	// Create a real logger for testing
	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel) // Reduce noise in tests
	handler := NewUserHandler(mockService, validator, logger)

	return handler, mockService
}

func setupTestRouter(handler *UserHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", handler.CreateUser)
			users.GET("", handler.ListUsers)
			users.GET("/:id", handler.GetUser)
			users.PUT("/:id", handler.UpdateUser)
			users.DELETE("/:id", handler.DeleteUser)
		}
	}

	return router
}

func TestUserHandler_CreateUser(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		setupMock      func(*MockUserService)
	}{
		{
			name: "Successful user creation",
			requestBody: entity.CreateUserRequest{
				Name:        "Test User",
				Email:       "test@example.com",
				DateOfBirth: "1990-01-01",
				Phone:       "1234567890",
				Address:     "Test Address",
			},
			expectedStatus: http.StatusCreated,
			setupMock: func(mockService *MockUserService) {
				mockService.On("CreateUser", mock.Anything, mock.AnythingOfType("entity.CreateUserRequest")).Return(&entity.User{
					ID:          1,
					Name:        "Test User",
					Email:       "test@example.com",
					DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
					Phone:       "1234567890",
					Address:     "Test Address",
				}, nil)
			},
		},
		{
			name: "Invalid request body",
			requestBody: map[string]interface{}{
				"name":  123, // Invalid type
				"email": "invalid-email",
			},
			expectedStatus: http.StatusBadRequest,
			setupMock: func(mockService *MockUserService) {
				// No mock setup needed for validation failure
			},
		},
		{
			name: "Service error",
			requestBody: entity.CreateUserRequest{
				Name:        "Test User",
				Email:       "test@example.com",
				DateOfBirth: "1990-01-01",
			},
			expectedStatus: http.StatusInternalServerError,
			setupMock: func(mockService *MockUserService) {
				mockService.On("CreateUser", mock.Anything, mock.AnythingOfType("entity.CreateUserRequest")).Return(nil, errors.New("service error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler, mockService := setupTestHandler()
			router := setupTestRouter(handler)

			tt.setupMock(mockService)

			requestBody, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_GetUser(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		setupMock      func(*MockUserService)
	}{
		{
			name:           "Successful user retrieval",
			userID:         "1",
			expectedStatus: http.StatusOK,
			setupMock: func(mockService *MockUserService) {
				mockService.On("GetUser", mock.Anything, uint(1)).Return(&entity.User{
					ID:          1,
					Name:        "Test User",
					Email:       "test@example.com",
					DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
				}, nil)
			},
		},
		{
			name:           "Invalid user ID",
			userID:         "invalid",
			expectedStatus: http.StatusBadRequest,
			setupMock: func(mockService *MockUserService) {
				// No mock setup needed for validation failure
			},
		},
		{
			name:           "User not found",
			userID:         "999",
			expectedStatus: http.StatusNotFound,
			setupMock: func(mockService *MockUserService) {
				mockService.On("GetUser", mock.Anything, uint(999)).Return(nil, errors.New("user not found"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler, mockService := setupTestHandler()
			router := setupTestRouter(handler)

			tt.setupMock(mockService)

			req, _ := http.NewRequest("GET", "/api/v1/users/"+tt.userID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_ListUsers(t *testing.T) {
	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		setupMock      func(*MockUserService)
	}{
		{
			name:           "Successful user listing",
			queryParams:    "?page=1&per_page=10&sortBy=name&sortDir=asc",
			expectedStatus: http.StatusOK,
			setupMock: func(mockService *MockUserService) {
				mockService.On("ListUsers", mock.Anything, mock.AnythingOfType("entity.UserSearchParams")).Return(&entity.UserListResponse{
					Users: []entity.User{
						{
							ID:          1,
							Name:        "Alice",
							Email:       "alice@example.com",
							DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
						},
						{
							ID:          2,
							Name:        "Bob",
							Email:       "bob@example.com",
							DateOfBirth: time.Date(1991, 1, 1, 0, 0, 0, 0, time.UTC),
						},
					},
					Total:      2,
					Page:       1,
					PerPage:    10,
					TotalPages: 1,
				}, nil)
			},
		},
		{
			name:           "Invalid sort field",
			queryParams:    "?sortBy=invalid&sortDir=asc",
			expectedStatus: http.StatusBadRequest,
			setupMock: func(mockService *MockUserService) {
				// No mock setup needed for validation failure
			},
		},
		{
			name:           "Service error",
			queryParams:    "?page=1&per_page=10",
			expectedStatus: http.StatusInternalServerError,
			setupMock: func(mockService *MockUserService) {
				mockService.On("ListUsers", mock.Anything, mock.AnythingOfType("entity.UserSearchParams")).Return(nil, errors.New("service error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler, mockService := setupTestHandler()
			router := setupTestRouter(handler)

			tt.setupMock(mockService)

			req, _ := http.NewRequest("GET", "/api/v1/users"+tt.queryParams, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_UpdateUser(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		requestBody    interface{}
		expectedStatus int
		setupMock      func(*MockUserService)
	}{
		{
			name:   "Successful user update",
			userID: "1",
			requestBody: entity.UpdateUserRequest{
				Name:  stringPtr("Updated Name"),
				Email: stringPtr("updated@example.com"),
			},
			expectedStatus: http.StatusOK,
			setupMock: func(mockService *MockUserService) {
				mockService.On("UpdateUser", mock.Anything, uint(1), mock.AnythingOfType("entity.UpdateUserRequest")).Return(&entity.User{
					ID:          1,
					Name:        "Updated Name",
					Email:       "updated@example.com",
					DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
				}, nil)
			},
		},
		{
			name:           "Invalid user ID",
			userID:         "invalid",
			requestBody:    entity.UpdateUserRequest{},
			expectedStatus: http.StatusBadRequest,
			setupMock: func(mockService *MockUserService) {
				// No mock setup needed for validation failure
			},
		},
		{
			name:   "Service error",
			userID: "1",
			requestBody: entity.UpdateUserRequest{
				Name: stringPtr("Updated Name"),
			},
			expectedStatus: http.StatusInternalServerError,
			setupMock: func(mockService *MockUserService) {
				mockService.On("UpdateUser", mock.Anything, uint(1), mock.AnythingOfType("entity.UpdateUserRequest")).Return(nil, errors.New("service error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler, mockService := setupTestHandler()
			router := setupTestRouter(handler)

			tt.setupMock(mockService)

			requestBody, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("PUT", "/api/v1/users/"+tt.userID, bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		setupMock      func(*MockUserService)
	}{
		{
			name:           "Successful user deletion",
			userID:         "1",
			expectedStatus: http.StatusOK,
			setupMock: func(mockService *MockUserService) {
				mockService.On("DeleteUser", mock.Anything, uint(1)).Return(nil)
			},
		},
		{
			name:           "Invalid user ID",
			userID:         "invalid",
			expectedStatus: http.StatusBadRequest,
			setupMock: func(mockService *MockUserService) {
				// No mock setup needed for validation failure
			},
		},
		{
			name:           "User not found",
			userID:         "999",
			expectedStatus: http.StatusNotFound,
			setupMock: func(mockService *MockUserService) {
				mockService.On("DeleteUser", mock.Anything, uint(999)).Return(errors.New("user not found"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler, mockService := setupTestHandler()
			router := setupTestRouter(handler)

			tt.setupMock(mockService)

			req, _ := http.NewRequest("DELETE", "/api/v1/users/"+tt.userID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}

func TestGetValidationMessage(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		tag      string
		param    string
		expected string
	}{
		{
			name:     "Required field",
			field:    "Name",
			tag:      "required",
			param:    "",
			expected: "This field is required",
		},
		{
			name:     "Email validation",
			field:    "Email",
			tag:      "email",
			param:    "",
			expected: "Must be a valid email address",
		},
		{
			name:     "Min length string",
			field:    "Name",
			tag:      "min",
			param:    "2",
			expected: "Must be at least 2 characters long",
		},
		{
			name:     "Min value int",
			field:    "Age",
			tag:      "min",
			param:    "18",
			expected: "Must be at least 18 characters long",
		},
		{
			name:     "Max length string",
			field:    "Name",
			tag:      "max",
			param:    "100",
			expected: "Must be at most 100 characters long",
		},
		{
			name:     "Max value int",
			field:    "Age",
			tag:      "max",
			param:    "120",
			expected: "Must be at most 120 characters long",
		},
		{
			name:     "Unknown tag",
			field:    "Field",
			tag:      "unknown",
			param:    "",
			expected: "Invalid value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock validator.FieldError
			fieldError := &mockFieldError{
				field: tt.field,
				tag:   tt.tag,
				param: tt.param,
			}

			result := getValidationMessage(fieldError)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// mockFieldError is a mock implementation of validator.FieldError for testing
type mockFieldError struct {
	field string
	tag   string
	param string
}

func (m *mockFieldError) Tag() string                               { return m.tag }
func (m *mockFieldError) ActualTag() string                         { return m.tag }
func (m *mockFieldError) Namespace() string                         { return m.field }
func (m *mockFieldError) StructNamespace() string                   { return m.field }
func (m *mockFieldError) Field() string                             { return m.field }
func (m *mockFieldError) StructField() string                       { return m.field }
func (m *mockFieldError) Value() interface{}                        { return nil }
func (m *mockFieldError) Param() string                             { return m.param }
func (m *mockFieldError) Kind() reflect.Kind                        { return reflect.String }
func (m *mockFieldError) Type() reflect.Type                        { return reflect.TypeOf("") }
func (m *mockFieldError) Translate(translator ut.Translator) string { return "" }
func (m *mockFieldError) Error() string                             { return m.field + " " + m.tag }

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}
