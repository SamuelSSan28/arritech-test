package mysql

import (
	"context"
	"fmt"
	"math"
	"strings"

	"arritech-user-management/internal/domain/entity"
	"arritech-user-management/internal/domain/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new MySQL user repository
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "email") {
			return fmt.Errorf("email already exists")
		}
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "email") {
			return fmt.Errorf("email already exists")
		}
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&entity.User{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (r *userRepository) List(ctx context.Context, params entity.UserSearchParams) (*entity.UserListResponse, error) {
	var users []entity.User
	var total int64

	// Log received parameters
	logrus.WithFields(logrus.Fields{
		"sort_by":  params.SortBy,
		"sort_dir": params.SortDir,
		"page":     params.Page,
		"per_page": params.PerPage,
		"search":   params.Search,
	}).Info("Repository: Starting List operation with parameters")

	// Set default pagination values
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PerPage == 0 {
		params.PerPage = 10
	}

	// Set default sorting
	if params.SortBy == "" {
		params.SortBy = "created_at"
	}
	if params.SortDir == "" {
		params.SortDir = "desc"
	}

	logrus.WithFields(logrus.Fields{
		"final_sort_by":  params.SortBy,
		"final_sort_dir": params.SortDir,
		"final_page":     params.Page,
		"final_per_page": params.PerPage,
	}).Info("Repository: Parameters after setting defaults")

	query := r.db.WithContext(ctx).Model(&entity.User{})

	// Apply search filter
	if params.Search != "" {
		searchTerm := "%" + params.Search + "%"
		query = query.Where("name LIKE ? OR email LIKE ? OR phone LIKE ?", searchTerm, searchTerm, searchTerm)
		logrus.WithField("search_term", searchTerm).Info("Repository: Applied search filter")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		logrus.WithError(err).Error("Repository: Failed to count users")
		return nil, fmt.Errorf("failed to count users: %w", err)
	}

	logrus.WithField("total_users", total).Info("Repository: Total users count retrieved")

	// Apply sorting
	sortField := getSortField(params.SortBy)
	logrus.WithFields(logrus.Fields{
		"original_sort_by":  params.SortBy,
		"mapped_sort_field": sortField,
		"sort_direction":    params.SortDir,
	}).Info("Repository: Applying sorting")

	// Build the ORDER BY clause
	if params.SortBy == "age" {
		// For age sorting, we need to order by date_of_birth in reverse order
		if params.SortDir == "asc" {
			query = query.Order("date_of_birth DESC") // Older people first
			logrus.Info("Repository: Applied age sorting ASC (date_of_birth DESC)")
		} else {
			query = query.Order("date_of_birth ASC") // Younger people first
			logrus.Info("Repository: Applied age sorting DESC (date_of_birth ASC)")
		}
	} else {
		// Sanitize the sort direction
		dir := strings.ToUpper(params.SortDir)
		if dir != "ASC" && dir != "DESC" {
			dir = "DESC"
			logrus.WithField("original_dir", params.SortDir).Warn("Repository: Invalid sort direction, defaulting to DESC")
		}
		orderClause := fmt.Sprintf("%s %s", sortField, dir)
		query = query.Order(orderClause)
		logrus.WithField("order_clause", orderClause).Info("Repository: Applied standard sorting")
	}

	// Apply pagination
	offset := (params.Page - 1) * params.PerPage
	logrus.WithFields(logrus.Fields{
		"offset": offset,
		"limit":  params.PerPage,
	}).Info("Repository: Applying pagination")

	if err := query.Offset(offset).Limit(params.PerPage).Find(&users).Error; err != nil {
		logrus.WithError(err).Error("Repository: Failed to find users")
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	logrus.WithField("users_returned", len(users)).Info("Repository: Successfully retrieved users from database")

	totalPages := int(math.Ceil(float64(total) / float64(params.PerPage)))

	return &entity.UserListResponse{
		Users:      users,
		Total:      total,
		Page:       params.Page,
		PerPage:    params.PerPage,
		TotalPages: totalPages,
	}, nil
}

// getSortField returns the database field name for sorting
func getSortField(sortBy string) string {
	logrus.WithField("input_sort_by", sortBy).Info("Repository: getSortField called")

	result := ""
	switch sortBy {
	case "name":
		result = "name"
	case "email":
		result = "email"
	case "age":
		result = "date_of_birth" // Sort by date of birth for age
	case "phone":
		result = "phone"
	case "created_at":
		result = "created_at"
	case "updated_at":
		result = "updated_at"
	case "id":
		result = "id"
	default:
		result = "created_at"
	}

	logrus.WithFields(logrus.Fields{
		"input_sort_by": sortBy,
		"mapped_field":  result,
	}).Info("Repository: getSortField mapping completed")

	return result
}

func (r *userRepository) EmailExists(ctx context.Context, email string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email)

	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}

	if err := query.Count(&count).Error; err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}

	return count > 0, nil
}
