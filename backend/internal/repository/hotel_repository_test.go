package repository

import (
	"backend/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupHotelTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Create tables manually for SQLite (no PostgreSQL-specific defaults)
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
	if err != nil {
		t.Fatalf("Failed to create tables: %v", err)
	}

	return db
}

func TestHotelRepository_CreateHotel(t *testing.T) {
	db := setupHotelTestDB(t)
	repo := NewHotelRepository(db)

	tests := []struct {
		name          string
		hotel         *domain.Hotel
		expectedError bool
	}{
		{
			name: "successful hotel creation",
			hotel: &domain.Hotel{
				Id:          uuid.New(),
				Name:        "Test Hotel",
				Description: "A test hotel",
				Address:     "123 Test St",
				Rating:      4.5,
			},
			expectedError: false,
		},
		{
			name: "hotel with rooms",
			hotel: &domain.Hotel{
				Id:          uuid.New(),
				Name:        "Hotel with Rooms",
				Description: "Hotel description",
				Address:     "456 Room Ave",
				Rating:      4.8,
				Rooms: []*domain.Room{
					{
						Id:          uuid.New(),
						Size:        25,
						Price:       100.0,
						Description: "Standard room",
						Available:   true,
						HotelId:     uuid.New(),
					},
				},
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateHotel(tt.hotel)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				// Verify hotel was created
				var createdHotel domain.Hotel
				result := db.First(&createdHotel, "id = ?", tt.hotel.Id)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.hotel.Name, createdHotel.Name)
				assert.Equal(t, tt.hotel.Address, createdHotel.Address)
			}
		})
	}
}

func TestHotelRepository_GetAllHotels(t *testing.T) {
	db := setupHotelTestDB(t)
	repo := NewHotelRepository(db)

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

	repo.CreateHotel(hotel1)
	repo.CreateHotel(hotel2)

	hotels, err := repo.GetAllHotels()
	assert.NoError(t, err)
	assert.Len(t, hotels, 2)
	assert.Equal(t, hotel1.Name, hotels[0].Name)
	assert.Equal(t, hotel2.Name, hotels[1].Name)
}

func TestHotelRepository_GetHotelById(t *testing.T) {
	db := setupHotelTestDB(t)
	repo := NewHotelRepository(db)

	hotel := &domain.Hotel{
		Id:          uuid.New(),
		Name:        "Test Hotel",
		Description: "Test Description",
		Address:     "Test Address",
		Rating:      4.5,
	}

	repo.CreateHotel(hotel)

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
			result, err := repo.GetHotelById(tt.id)
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

