package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestHotelService_CreateHotel(t *testing.T) {
	tests := []struct {
		name          string
		hotel         *domain.Hotel
		shouldError   bool
		expectedError string
	}{
		{
			name: "successful hotel creation",
			hotel: &domain.Hotel{
				Id:          uuid.New(),
				Name:        "Test Hotel",
				Description: "Test Description",
				Address:     "Test Address",
				Rating:      4.5,
			},
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
			assert.NoError(t, err)

			// Create tables manually for SQLite compatibility
			err = db.Exec(`
				CREATE TABLE hotels (
					id TEXT PRIMARY KEY,
					name TEXT,
					description TEXT,
					address TEXT,
					rating REAL
				);
				CREATE TABLE rooms (
					id TEXT PRIMARY KEY,
					size INTEGER,
					price REAL,
					description TEXT,
					available INTEGER DEFAULT 1,
					hotel_id TEXT
				);
				CREATE TABLE facilities (
					id TEXT PRIMARY KEY,
					name TEXT
				);
				CREATE TABLE room_facilities (
					room_id TEXT,
					facility_id TEXT,
					PRIMARY KEY (room_id, facility_id)
				);
			`).Error
			assert.NoError(t, err)

			repo := repository.NewHotelRepository(db)
			service := NewHotelService(repo)
			err = service.CreateHotel(tt.hotel)

			if tt.shouldError {
				assert.Error(t, err)
				if tt.expectedError != "" {
					assert.Contains(t, err.Error(), tt.expectedError)
				}
			} else {
				assert.NoError(t, err)

				// Verify hotel was created
				var createdHotel domain.Hotel
				result := db.First(&createdHotel, "id = ?", tt.hotel.Id)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.hotel.Name, createdHotel.Name)
			}
		})
	}
}

func TestHotelService_GetAllHotels(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Create tables manually for SQLite compatibility
	err = db.Exec(`
		CREATE TABLE hotels (
			id TEXT PRIMARY KEY,
			name TEXT,
			description TEXT,
			address TEXT,
			rating REAL
		);
		CREATE TABLE rooms (
			id TEXT PRIMARY KEY,
			size INTEGER,
			price REAL,
			description TEXT,
			available INTEGER DEFAULT 1,
			hotel_id TEXT
		);
		CREATE TABLE facilities (
			id TEXT PRIMARY KEY,
			name TEXT
		);
		CREATE TABLE room_facilities (
			room_id TEXT,
			facility_id TEXT,
			PRIMARY KEY (room_id, facility_id)
		);
	`).Error
	assert.NoError(t, err)

	repo := repository.NewHotelRepository(db)
	service := NewHotelService(repo)

	// Create test hotels
	hotel1 := &domain.Hotel{
		Id:          uuid.New(),
		Name:        "Hotel 1",
		Description: "First hotel",
		Address:     "Address 1",
		Rating:      4.0,
	}

	hotel2 := &domain.Hotel{
		Id:          uuid.New(),
		Name:        "Hotel 2",
		Description: "Second hotel",
		Address:     "Address 2",
		Rating:      4.5,
	}

	service.CreateHotel(hotel1)
	service.CreateHotel(hotel2)

	hotels, err := service.GetAllHotels()
	assert.NoError(t, err)
	assert.Len(t, hotels, 2)
}

func TestHotelService_GetHotelById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Create tables manually for SQLite compatibility
	err = db.Exec(`
		CREATE TABLE hotels (
			id TEXT PRIMARY KEY,
			name TEXT,
			description TEXT,
			address TEXT,
			rating REAL
		);
		CREATE TABLE rooms (
			id TEXT PRIMARY KEY,
			size INTEGER,
			price REAL,
			description TEXT,
			available INTEGER DEFAULT 1,
			hotel_id TEXT
		);
		CREATE TABLE facilities (
			id TEXT PRIMARY KEY,
			name TEXT
		);
		CREATE TABLE room_facilities (
			room_id TEXT,
			facility_id TEXT,
			PRIMARY KEY (room_id, facility_id)
		);
	`).Error
	assert.NoError(t, err)

	repo := repository.NewHotelRepository(db)
	service := NewHotelService(repo)

	hotel := &domain.Hotel{
		Id:          uuid.New(),
		Name:        "Test Hotel",
		Description: "Test Description",
		Address:     "Test Address",
		Rating:      4.5,
	}

	service.CreateHotel(hotel)

	tests := []struct {
		name          string
		id            string
		expectedError bool
		expectedName  string
	}{
		{
			name:          "existing hotel",
			id:            hotel.Id.String(),
			expectedError: false,
			expectedName:  "Test Hotel",
		},
		{
			name:          "non-existent hotel",
			id:            uuid.New().String(),
			expectedError: true,
			expectedName:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetHotelById(tt.id)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedName, result.Name)
			}
		})
	}
}

func TestNewHotelService(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Create tables manually for SQLite compatibility
	err = db.Exec(`
		CREATE TABLE hotels (
			id TEXT PRIMARY KEY,
			name TEXT,
			description TEXT,
			address TEXT,
			rating REAL
		);
		CREATE TABLE rooms (
			id TEXT PRIMARY KEY,
			size INTEGER,
			price REAL,
			description TEXT,
			available INTEGER DEFAULT 1,
			hotel_id TEXT
		);
		CREATE TABLE facilities (
			id TEXT PRIMARY KEY,
			name TEXT
		);
		CREATE TABLE room_facilities (
			room_id TEXT,
			facility_id TEXT,
			PRIMARY KEY (room_id, facility_id)
		);
	`).Error
	assert.NoError(t, err)

	repo := repository.NewHotelRepository(db)
	service := NewHotelService(repo)

	assert.NotNil(t, service)
	assert.Equal(t, repo, service.(*hotelService).hotelRepository)
}
