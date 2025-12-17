package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserService_CreateUser(t *testing.T) {
	tests := []struct {
		name          string
		user          *domain.User
		shouldError   bool
		expectedError string
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
			shouldError: false,
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
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test database
			db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
			assert.NoError(t, err)

			// Create table manually for SQLite compatibility
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
			assert.NoError(t, err)

			repo := repository.NewUserRepository(db)
			service := NewUserService(repo)
			err = service.CreateUser(tt.user)

			if tt.shouldError {
				assert.Error(t, err)
				if tt.expectedError != "" {
					assert.Contains(t, err.Error(), tt.expectedError)
				}
			} else {
				assert.NoError(t, err)

				// Verify user was created
				var createdUser domain.User
				result := db.First(&createdUser, "id = ?", tt.user.Id)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.user.Username, createdUser.Username)
				assert.Equal(t, tt.user.Email, createdUser.Email)
			}
		})
	}
}

func TestUserService_CreateUser_DuplicateEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Create table manually for SQLite compatibility
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
	assert.NoError(t, err)

	repo := repository.NewUserRepository(db)
	service := NewUserService(repo)

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
	err = service.CreateUser(user1)
	assert.NoError(t, err)

	// Try to create second user with same email
	err = service.CreateUser(user2)
	assert.Error(t, err) // Should fail due to unique constraint
}

func TestNewUserService(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Create table manually for SQLite compatibility
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
	assert.NoError(t, err)

	repo := repository.NewUserRepository(db)
	service := NewUserService(repo)

	assert.NotNil(t, service)
	assert.Equal(t, repo, service.(*userService).userRepository)
}
