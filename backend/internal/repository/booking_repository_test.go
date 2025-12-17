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

func setupBookingTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Create table manually for SQLite (no PostgreSQL-specific defaults)
	err = db.Exec(`
		CREATE TABLE bookings (
			id TEXT PRIMARY KEY,
			user_id TEXT,
			hotel_id TEXT,
			room_id TEXT,
			check_in_date DATETIME,
			check_out_date DATETIME,
			total_price REAL,
			is_cancelled INTEGER DEFAULT 0
		)
	`).Error
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	return db
}

func TestBookingRepository_CreateBooking(t *testing.T) {
	db := setupBookingTestDB(t)
	repo := NewBookingRepository(db)

	tests := []struct {
		name          string
		booking       *domain.Booking
		expectedError bool
	}{
		{
			name: "successful booking creation",
			booking: &domain.Booking{
				Id:           uuid.New(),
				UserId:       uuid.New(),
				HotelId:      uuid.New(),
				RoomId:       uuid.New(),
				CheckInDate:  time.Now(),
				CheckOutDate: time.Now().AddDate(0, 0, 3),
				TotalPrice:   300.0,
				IsCancelled:  false,
			},
			expectedError: false,
		},
		{
			name: "cancelled booking",
			booking: &domain.Booking{
				Id:           uuid.New(),
				UserId:       uuid.New(),
				HotelId:      uuid.New(),
				RoomId:       uuid.New(),
				CheckInDate:  time.Now(),
				CheckOutDate: time.Now().AddDate(0, 0, 5),
				TotalPrice:   500.0,
				IsCancelled:  true,
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateBooking(tt.booking)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				// Verify booking was created
				var createdBooking domain.Booking
				result := db.First(&createdBooking, "id = ?", tt.booking.Id)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.booking.TotalPrice, createdBooking.TotalPrice)
				assert.Equal(t, tt.booking.IsCancelled, createdBooking.IsCancelled)
			}
		})
	}
}

func TestBookingRepository_GetAllBookings(t *testing.T) {
	db := setupBookingTestDB(t)
	repo := NewBookingRepository(db)

	// Create test bookings
	booking1 := &domain.Booking{
		Id:           uuid.New(),
		UserId:       uuid.New(),
		HotelId:      uuid.New(),
		RoomId:       uuid.New(),
		CheckInDate:  time.Now(),
		CheckOutDate: time.Now().AddDate(0, 0, 2),
		TotalPrice:   200.0,
		IsCancelled:  false,
	}

	booking2 := &domain.Booking{
		Id:           uuid.New(),
		UserId:       uuid.New(),
		HotelId:      uuid.New(),
		RoomId:       uuid.New(),
		CheckInDate:  time.Now(),
		CheckOutDate: time.Now().AddDate(0, 0, 3),
		TotalPrice:   300.0,
		IsCancelled:  false,
	}

	repo.CreateBooking(booking1)
	repo.CreateBooking(booking2)

	bookings, err := repo.GetAllBookings()
	assert.NoError(t, err)
	assert.Len(t, bookings, 2)
}

func TestBookingRepository_GetBookingById(t *testing.T) {
	db := setupBookingTestDB(t)
	repo := NewBookingRepository(db)

	booking := &domain.Booking{
		Id:           uuid.New(),
		UserId:       uuid.New(),
		HotelId:      uuid.New(),
		RoomId:       uuid.New(),
		CheckInDate:  time.Now(),
		CheckOutDate: time.Now().AddDate(0, 0, 4),
		TotalPrice:   400.0,
		IsCancelled:  false,
	}

	repo.CreateBooking(booking)

	tests := []struct {
		name          string
		id            string
		expectedError bool
		expectedPrice float64
	}{
		{
			name:          "existing booking",
			id:            booking.Id.String(),
			expectedError: false,
			expectedPrice: 400.0,
		},
		{
			name:          "non-existent booking",
			id:            uuid.New().String(),
			expectedError: true,
			expectedPrice: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := repo.GetBookingById(tt.id)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedPrice, result.TotalPrice)
			}
		})
	}
}
