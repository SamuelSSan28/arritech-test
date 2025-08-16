package entity

import (
	"gorm.io/gorm"
	"time"
)

// User represents a user in the system
type User struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Name        string         `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Email       string         `json:"email" gorm:"uniqueIndex;not null;size:255" validate:"required,email"`
	DateOfBirth time.Time      `json:"date_of_birth" gorm:"not null" validate:"required"`
	Age         int            `json:"age" gorm:"-"` // Computed field, not stored
	Phone       string         `json:"phone,omitempty" gorm:"size:20" validate:"omitempty,min=10,max=20"`
	Address     string         `json:"address,omitempty" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName returns the table name for the User entity
func (User) TableName() string {
	return "users"
}

// CalculateAge calculates the age from date of birth
func (u *User) CalculateAge() int {
	if u.DateOfBirth.IsZero() {
		return 0
	}

	today := time.Now()
	age := today.Year() - u.DateOfBirth.Year()

	// Adjust if birthday hasn't occurred this year
	if today.YearDay() < u.DateOfBirth.YearDay() {
		age--
	}

	return age
}

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
	Phone       string `json:"phone,omitempty" validate:"omitempty,min=10,max=20"`
	Address     string `json:"address,omitempty"`
}

// UpdateUserRequest represents the request payload for updating a user
type UpdateUserRequest struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Email       *string `json:"email,omitempty" validate:"omitempty,email"`
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	Phone       *string `json:"phone,omitempty" validate:"omitempty,min=10,max=20"`
	Address     *string `json:"address,omitempty"`
}

// UserListResponse represents the response for listing users with pagination
type UserListResponse struct {
	Users      []User `json:"users"`
	Total      int64  `json:"total"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	TotalPages int    `json:"total_pages"`
}

// UserSearchParams represents search parameters for users
type UserSearchParams struct {
	Search  string `json:"search" form:"search" query:"search"`
	Page    int    `json:"page" form:"page" query:"page" validate:"min=1"`
	PerPage int    `json:"per_page" form:"per_page" query:"per_page" validate:"min=1,max=100"`
	SortBy  string `json:"sortBy" form:"sortBy" query:"sortBy"`
	SortDir string `json:"sortDir" form:"sortDir" query:"sortDir" validate:"oneof=asc desc"`
}
