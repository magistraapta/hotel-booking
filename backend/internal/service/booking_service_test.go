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

func TestBookingService_CreateBooking(t *testing.T) {
	tests := []struct {
		name        string
		booking     *domain.Booking
		shouldError bool
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
			shouldError: false,
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
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
			assert.NoError(t, err)

			// Create table manually for SQLite compatibility
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
			assert.NoError(t, err)

			repo := repository.NewBookingRepository(db)
			service := NewBookingService(repo)
			err = service.CreateBooking(tt.booking)

			if tt.shouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				// Verify booking was created
				var createdBooking domain.Booking
				result := db.First(&createdBooking, "id = ?", tt.booking.Id)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.booking.TotalPrice, createdBooking.TotalPrice)
			}
		})
	}
}

func TestBookingService_GetAllBookings(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Create table manually for SQLite compatibility
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
	assert.NoError(t, err)

	repo := repository.NewBookingRepository(db)
	service := NewBookingService(repo)

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

	service.CreateBooking(booking1)
	service.CreateBooking(booking2)

	bookings, err := service.GetAllBookings()
	assert.NoError(t, err)
	assert.Len(t, bookings, 2)
}

func TestBookingService_GetBookingById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Create table manually for SQLite compatibility
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
	assert.NoError(t, err)

	repo := repository.NewBookingRepository(db)
	service := NewBookingService(repo)

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

	service.CreateBooking(booking)

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
			result, err := service.GetBookingById(tt.id)
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

func TestNewBookingService(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Create table manually for SQLite compatibility
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
	assert.NoError(t, err)

	repo := repository.NewBookingRepository(db)
	service := NewBookingService(repo)

	assert.NotNil(t, service)
	assert.Equal(t, repo, service.(*bookingService).bookingRepository)
}
