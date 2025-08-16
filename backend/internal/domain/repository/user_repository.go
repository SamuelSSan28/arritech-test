package repository

import (
	"arritech-user-management/internal/domain/entity"
	"context"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *entity.User) error

	// GetByID retrieves a user by ID
	GetByID(ctx context.Context, id uint) (*entity.User, error)

	// GetByEmail retrieves a user by email
	GetByEmail(ctx context.Context, email string) (*entity.User, error)

	// Update updates a user
	Update(ctx context.Context, user *entity.User) error

	// Delete soft deletes a user
	Delete(ctx context.Context, id uint) error

	// List retrieves users with pagination and search
	List(ctx context.Context, params entity.UserSearchParams) (*entity.UserListResponse, error)

	// EmailExists checks if an email already exists (for validation)
	EmailExists(ctx context.Context, email string, excludeID uint) (bool, error)
}
