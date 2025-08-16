package http

import (
	"net/http"
	"strconv"

	"arritech-user-management/internal/domain/entity"
	"arritech-user-management/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService service.UserService
	validator   *validator.Validate
	logger      *logrus.Logger
}

func NewUserHandler(userService service.UserService, validator *validator.Validate, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator,
		logger:      logger,
	}
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string            `json:"error"`
	Details map[string]string `json:"details,omitempty"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body entity.CreateUserRequest true "User information"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req entity.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
		})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = getValidationMessage(err)
		}
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Validation failed",
			Details: validationErrors,
		})
		return
	}

	user, err := h.userService.CreateUser(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "email already exists" || err.Error() == "user must be older than 18 years" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}
		h.logger.WithError(err).Error("Failed to create user")
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User created successfully",
		Data:    user,
	})
}

// GetUser retrieves a user by ID
// @Summary Get user by ID
// @Description Get user information by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	user, err := h.userService.GetUser(c.Request.Context(), uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
			return
		}
		h.logger.WithError(err).Error("Failed to get user")
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "User retrieved successfully",
		Data:    user,
	})
}

// UpdateUser updates a user
// @Summary Update user
// @Description Update user information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body entity.UpdateUserRequest true "User information to update"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	var req entity.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request format"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = getValidationMessage(err)
		}
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Validation failed",
			Details: validationErrors,
		})
		return
	}

	user, err := h.userService.UpdateUser(c.Request.Context(), uint(id), req)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
			return
		}
		if err.Error() == "email already exists" || err.Error() == "user must be older than 18 years" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}
		h.logger.WithError(err).Error("Failed to update user")
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "User updated successfully",
		Data:    user,
	})
}

// DeleteUser deletes a user
// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	err = h.userService.DeleteUser(c.Request.Context(), uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
			return
		}
		h.logger.WithError(err).Error("Failed to delete user")
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "User deleted successfully",
	})
}

// ListUsers retrieves users with pagination and search
// @Summary List users
// @Description Get list of users with pagination, search and sorting functionality
// @Tags users
// @Accept json
// @Produce json
// @Param search query string false "Search term"
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Param sort_by query string false "Sort field (name, email, age, phone, created_at, updated_at)" default(created_at)
// @Param sort_dir query string false "Sort direction (asc, desc)" default(desc)
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	// Log all query parameters received
	h.logger.WithFields(logrus.Fields{
		"raw_query":    c.Request.URL.RawQuery,
		"sort_by":      c.Query("sortBy"),
		"sort_dir":     c.Query("sortDir"),
		"sort_by_alt":  c.Query("sort_by"),
		"sort_dir_alt": c.Query("sort_dir"),
		"page":         c.Query("page"),
		"per_page":     c.Query("per_page"),
		"search":       c.Query("search"),
	}).Info("Received query parameters for ListUsers")

	// Log all available query parameters
	queryParams := c.Request.URL.Query()
	h.logger.WithField("all_query_params", queryParams).Info("All available query parameters")

	var params entity.UserSearchParams
	if err := c.ShouldBindQuery(&params); err != nil {
		h.logger.WithError(err).Error("Failed to bind query parameters")
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid query parameters"})
		return
	}

	// Log the bound parameters
	h.logger.WithFields(logrus.Fields{
		"bound_sort_by":  params.SortBy,
		"bound_sort_dir": params.SortDir,
		"bound_page":     params.Page,
		"bound_per_page": params.PerPage,
		"bound_search":   params.Search,
	}).Info("Bound query parameters")

	// Set defaults
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PerPage == 0 {
		params.PerPage = 10
	}
	if params.SortBy == "" {
		params.SortBy = "created_at"
	}
	if params.SortDir == "" {
		params.SortDir = "desc"
	}

	// Log parameters after setting defaults
	h.logger.WithFields(logrus.Fields{
		"final_sort_by":  params.SortBy,
		"final_sort_dir": params.SortDir,
		"final_page":     params.Page,
		"final_per_page": params.PerPage,
	}).Info("Parameters after setting defaults")

	// Validate sort_by field
	validSortFields := map[string]bool{
		"name": true, "email": true, "age": true, "phone": true,
		"created_at": true, "updated_at": true,
	}
	if !validSortFields[params.SortBy] {
		h.logger.WithField("invalid_sort_by", params.SortBy).Error("Invalid sort field provided")
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid sort field"})
		return
	}

	if err := h.validator.Struct(params); err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = getValidationMessage(err)
		}
		h.logger.WithError(err).WithField("validation_errors", validationErrors).Error("Validation failed")
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Validation failed",
			Details: validationErrors,
		})
		return
	}

	h.logger.WithFields(logrus.Fields{
		"validated_sort_by":  params.SortBy,
		"validated_sort_dir": params.SortDir,
	}).Info("Parameters validated successfully, calling service")

	result, err := h.userService.ListUsers(c.Request.Context(), params)
	if err != nil {
		h.logger.WithError(err).Error("Failed to list users")
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to list users"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Users retrieved successfully",
		Data:    result,
	})
}

// getValidationMessage returns a user-friendly validation message
func getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Must be a valid email address"
	case "min":
		if err.Type().String() == "int" {
			return "Must be at least " + err.Param()
		}
		return "Must be at least " + err.Param() + " characters long"
	case "max":
		if err.Type().String() == "int" {
			return "Must be at most " + err.Param()
		}
		return "Must be at most " + err.Param() + " characters long"
	default:
		return "Invalid value"
	}
}
