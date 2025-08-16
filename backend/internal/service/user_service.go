package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"arritech-user-management/internal/domain/entity"
	"arritech-user-management/internal/domain/repository"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(ctx context.Context, req entity.CreateUserRequest) (*entity.User, error)
	GetUser(ctx context.Context, id uint) (*entity.User, error)
	UpdateUser(ctx context.Context, id uint, req entity.UpdateUserRequest) (*entity.User, error)
	DeleteUser(ctx context.Context, id uint) error
	ListUsers(ctx context.Context, params entity.UserSearchParams) (*entity.UserListResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
	logger   *logrus.Logger
}

func NewUserService(userRepo repository.UserRepository, logger *logrus.Logger) UserService {
	return &userService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *userService) CreateUser(ctx context.Context, req entity.CreateUserRequest) (*entity.User, error) {
	s.logger.WithField("email", req.Email).Info("Creating new user")

	// Business rule: Email must be unique
	exists, err := s.userRepo.EmailExists(ctx, req.Email, 0)
	if err != nil {
		s.logger.WithError(err).Error("Failed to check email existence")
		return nil, fmt.Errorf("failed to validate email: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("email already exists")
	}

	// Parse and validate date of birth
	dateOfBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return nil, fmt.Errorf("invalid date of birth format, use YYYY-MM-DD: %w", err)
	}

	// Business rule: User must be older than 18 years
	age := calculateAge(dateOfBirth)
	if age <= 18 {
		return nil, fmt.Errorf("user must be older than 18 years")
	}

	// Sanitize input
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	req.Name = strings.TrimSpace(req.Name)

	user := &entity.User{
		Name:        req.Name,
		Email:       req.Email,
		DateOfBirth: dateOfBirth,
		Phone:       req.Phone,
		Address:     req.Address,
	}

	// Set computed age for response
	user.Age = age

	if err := s.userRepo.Create(ctx, user); err != nil {
		s.logger.WithError(err).Error("Failed to create user")
		return nil, err
	}

	s.logger.WithField("user_id", user.ID).Info("User created successfully")
	return user, nil
}

func (s *userService) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	s.logger.WithField("user_id", id).Info("Getting user")

	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.WithError(err).WithField("user_id", id).Error("Failed to get user")
		return nil, err
	}

	// Set computed age
	user.Age = user.CalculateAge()

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, id uint, req entity.UpdateUserRequest) (*entity.User, error) {
	s.logger.WithField("user_id", id).Info("Updating user")

	// Get existing user
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.WithError(err).WithField("user_id", id).Error("Failed to get user for update")
		return nil, err
	}

	// Business rule: Email must be unique (if being updated)
	if req.Email != nil {
		email := strings.ToLower(strings.TrimSpace(*req.Email))
		exists, err := s.userRepo.EmailExists(ctx, email, id)
		if err != nil {
			s.logger.WithError(err).Error("Failed to check email existence")
			return nil, fmt.Errorf("failed to validate email: %w", err)
		}
		if exists {
			return nil, fmt.Errorf("email already exists")
		}
		user.Email = email
	}

	// Business rule: Date of birth validation (if being updated)
	if req.DateOfBirth != nil {
		dateOfBirth, err := time.Parse("2006-01-02", *req.DateOfBirth)
		if err != nil {
			return nil, fmt.Errorf("invalid date of birth format, use YYYY-MM-DD: %w", err)
		}

		age := calculateAge(dateOfBirth)
		if age <= 18 {
			return nil, fmt.Errorf("user must be older than 18 years")
		}

		user.DateOfBirth = dateOfBirth
	}

	// Update other fields
	if req.Name != nil {
		user.Name = strings.TrimSpace(*req.Name)
	}
	if req.Phone != nil {
		user.Phone = *req.Phone
	}
	if req.Address != nil {
		user.Address = *req.Address
	}

	// Set computed age
	user.Age = user.CalculateAge()

	if err := s.userRepo.Update(ctx, user); err != nil {
		s.logger.WithError(err).WithField("user_id", id).Error("Failed to update user")
		return nil, err
	}

	s.logger.WithField("user_id", id).Info("User updated successfully")
	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	s.logger.WithField("user_id", id).Info("Deleting user")

	if err := s.userRepo.Delete(ctx, id); err != nil {
		s.logger.WithError(err).WithField("user_id", id).Error("Failed to delete user")
		return err
	}

	s.logger.WithField("user_id", id).Info("User deleted successfully")
	return nil
}

func (s *userService) ListUsers(ctx context.Context, params entity.UserSearchParams) (*entity.UserListResponse, error) {
	s.logger.WithFields(logrus.Fields{
		"search":   params.Search,
		"page":     params.Page,
		"per_page": params.PerPage,
		"sort_by":  params.SortBy,
		"sort_dir": params.SortDir,
	}).Info("Service: Listing users with parameters")

	result, err := s.userRepo.List(ctx, params)
	if err != nil {
		s.logger.WithError(err).Error("Service: Failed to list users from repository")
		return nil, err
	}

	s.logger.WithFields(logrus.Fields{
		"total_users": result.Total,
		"sort_by":     params.SortBy,
		"sort_dir":    params.SortDir,
	}).Info("Service: Repository returned users, calculating ages")

	// Calculate age for all users
	for i := range result.Users {
		result.Users[i].Age = result.Users[i].CalculateAge()
	}

	s.logger.WithField("total_users", result.Total).Info("Service: Users listed successfully with ages calculated")
	return result, nil
}

// Helper function to calculate age from date of birth
func calculateAge(birthDate time.Time) int {
	if birthDate.IsZero() {
		return 0
	}

	today := time.Now()
	age := today.Year() - birthDate.Year()

	// Adjust if birthday hasn't occurred this year
	if today.YearDay() < birthDate.YearDay() {
		age--
	}

	return age
}
