package repository

import (
	"backend/internal/domain"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Create table manually for SQLite (no PostgreSQL-specific defaults)
	err = db.Exec(`
		CREATE TABLE users (
			id TEXT PRIMARY KEY,
			username TEXT UNIQUE,
			email TEXT UNIQUE,
			password TEXT,
			created_at DATETIME,
			updated_at DATETIME,
			is_admin INTEGER DEFAULT 0
		)
	`).Error
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	return db
}

func TestUserRepository_CreateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	tests := []struct {
		name          string
		user          *domain.User
		expectedError error
	}{
		{
			name: "successful user creation",
			user: &domain.User{
				Id:        uuid.New(),
				Username:  "testuser",
				Email:     "test@example.com",
				Password:  "password123",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsAdmin:   false,
			},
			expectedError: nil,
		},
		{
			name: "user with admin privileges",
			user: &domain.User{
				Id:        uuid.New(),
				Username:  "admin",
				Email:     "admin@example.com",
				Password:  "admin123",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsAdmin:   true,
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateUser(tt.user)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)

				// Verify user was created
				var createdUser domain.User
				result := db.First(&createdUser, "id = ?", tt.user.Id)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.user.Username, createdUser.Username)
				assert.Equal(t, tt.user.Email, createdUser.Email)
				assert.Equal(t, tt.user.IsAdmin, createdUser.IsAdmin)
			}
		})
	}
}

func TestUserRepository_CreateUser_DuplicateEmail(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	user1 := &domain.User{
		Id:        uuid.New(),
		Username:  "user1",
		Email:     "duplicate@example.com",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsAdmin:   false,
	}

	user2 := &domain.User{
		Id:        uuid.New(),
		Username:  "user2",
		Email:     "duplicate@example.com", // Same email
		Password:  "password456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsAdmin:   false,
	}

	// Create first user
	err := repo.CreateUser(user1)
	assert.NoError(t, err)

	// Try to create second user with same email
	err = repo.CreateUser(user2)
	assert.Error(t, err) // Should fail due to unique constraint
}
