package mysql

import (
	"context"
	"testing"
	"time"

	"arritech-user-management/internal/domain/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	require.NoError(t, err)

	return gormDB, mock, func() {
		db.Close()
	}
}

func TestUserRepository_Create(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	user := &entity.User{
		Name:        "Test User",
		Email:       "test@example.com",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:       "1234567890",
		Address:     "Test Address",
	}

	// GORM uses transactions, so we need to expect begin and commit
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").
		WithArgs(user.Name, user.Email, user.DateOfBirth, user.Phone, user.Address, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Create(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_GetByID(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	expectedUser := &entity.User{
		ID:          1,
		Name:        "Test User",
		Email:       "test@example.com",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:       "1234567890",
		Address:     "Test Address",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "date_of_birth", "phone", "address", "created_at", "updated_at"}).
		AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email, expectedUser.DateOfBirth, expectedUser.Phone, expectedUser.Address, expectedUser.CreatedAt, expectedUser.UpdatedAt)

	mock.ExpectQuery("SELECT \\* FROM `users`").
		WithArgs(1).
		WillReturnRows(rows)

	user, err := repo.GetByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Email, user.Email)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_List(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	params := entity.UserSearchParams{
		Page:    1,
		PerPage: 10,
		SortBy:  "name",
		SortDir: "asc",
		Search:  "",
	}

	// Mock count query
	countRows := sqlmock.NewRows([]string{"count"}).AddRow(2)
	mock.ExpectQuery("SELECT count\\(\\*\\) FROM `users`").
		WillReturnRows(countRows)

	// Mock main query
	userRows := sqlmock.NewRows([]string{"id", "name", "email", "date_of_birth", "phone", "address", "created_at", "updated_at"}).
		AddRow(1, "Alice", "alice@example.com", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "1234567890", "Address 1", time.Now(), time.Now()).
		AddRow(2, "Bob", "bob@example.com", time.Date(1991, 1, 1, 0, 0, 0, 0, time.UTC), "0987654321", "Address 2", time.Now(), time.Now())

	mock.ExpectQuery("SELECT \\* FROM `users`").
		WillReturnRows(userRows)

	result, err := repo.List(context.Background(), params)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), result.Total)
	assert.Len(t, result.Users, 2)
	assert.Equal(t, 1, result.Page)
	assert.Equal(t, 10, result.PerPage)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_ListWithSearch(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	params := entity.UserSearchParams{
		Page:    1,
		PerPage: 10,
		SortBy:  "name",
		SortDir: "asc",
		Search:  "alice",
	}

	// Mock count query with search
	countRows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT count\\(\\*\\) FROM `users`").
		WillReturnRows(countRows)

	// Mock main query with search
	userRows := sqlmock.NewRows([]string{"id", "name", "email", "date_of_birth", "phone", "address", "created_at", "updated_at"}).
		AddRow(1, "Alice", "alice@example.com", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "1234567890", "Address 1", time.Now(), time.Now())

	mock.ExpectQuery("SELECT \\* FROM `users`").
		WillReturnRows(userRows)

	result, err := repo.List(context.Background(), params)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result.Total)
	assert.Len(t, result.Users, 1)
	assert.Equal(t, "Alice", result.Users[0].Name)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_ListWithAgeSorting(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	params := entity.UserSearchParams{
		Page:    1,
		PerPage: 10,
		SortBy:  "age",
		SortDir: "desc",
		Search:  "",
	}

	// Mock count query
	countRows := sqlmock.NewRows([]string{"count"}).AddRow(2)
	mock.ExpectQuery("SELECT count\\(\\*\\) FROM `users`").
		WillReturnRows(countRows)

	// Mock main query with age sorting (should order by date_of_birth ASC for desc age)
	userRows := sqlmock.NewRows([]string{"id", "name", "email", "date_of_birth", "phone", "address", "created_at", "updated_at"}).
		AddRow(1, "Younger", "younger@example.com", time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), "1234567890", "Address 1", time.Now(), time.Now()).
		AddRow(2, "Older", "older@example.com", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "0987654321", "Address 2", time.Now(), time.Now())

	mock.ExpectQuery("SELECT \\* FROM `users`").
		WillReturnRows(userRows)

	result, err := repo.List(context.Background(), params)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), result.Total)
	assert.Len(t, result.Users, 2)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_Update(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	user := &entity.User{
		ID:          1,
		Name:        "Updated User",
		Email:       "updated@example.com",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:       "1234567890",
		Address:     "Updated Address",
	}

	// GORM uses transactions, so we need to expect begin and commit
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users`").
		WithArgs(user.Name, user.Email, user.DateOfBirth, user.Phone, user.Address, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.Update(context.Background(), user)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_Delete(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	// GORM uses soft delete by default, so it's actually an UPDATE
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users`").
		WithArgs(sqlmock.AnyArg(), uint(1)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.Delete(context.Background(), 1)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_EmailExists(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	// Test email exists
	countRows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT count\\(\\*\\) FROM `users`").
		WithArgs("test@example.com").
		WillReturnRows(countRows)

	exists, err := repo.EmailExists(context.Background(), "test@example.com", 0)
	assert.NoError(t, err)
	assert.True(t, exists)

	// Test email doesn't exist
	countRows = sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT count\\(\\*\\) FROM `users`").
		WithArgs("nonexistent@example.com").
		WillReturnRows(countRows)

	exists, err = repo.EmailExists(context.Background(), "nonexistent@example.com", 0)
	assert.NoError(t, err)
	assert.False(t, exists)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetSortField(t *testing.T) {
	tests := []struct {
		name     string
		sortBy   string
		expected string
	}{
		{"name field", "name", "name"},
		{"email field", "email", "email"},
		{"age field", "age", "date_of_birth"},
		{"phone field", "phone", "phone"},
		{"created_at field", "created_at", "created_at"},
		{"updated_at field", "updated_at", "updated_at"},
		{"id field", "id", "id"},
		{"invalid field", "invalid", "created_at"},
		{"empty field", "", "created_at"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getSortField(tt.sortBy)
			assert.Equal(t, tt.expected, result)
		})
	}
}
