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

func setupBookingServiceTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Create all required tables for SQLite compatibility
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
		t.Fatalf("Failed to create tables: %v", err)
	}

	return db
}

func createTestHotelAndRoom(db *gorm.DB, hotelId, roomId uuid.UUID, roomPrice float64) error {
	hotel := &domain.Hotel{
		Id:          hotelId,
		Name:        "Test Hotel",
		Description: "A test hotel",
		Address:     "123 Test St",
		Rating:      4.5,
		Rooms: []*domain.Room{
			{
				Id:          roomId,
				Size:        25,
				Price:       roomPrice,
				Description: "Standard room",
				Available:   true,
				HotelId:     hotelId,
			},
		},
	}
	return db.Create(hotel).Error
}

// func TestBookingService_GetBookingsByUserId(t *testing.T) {
// 	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
// 	assert.NoError(t, err)

// 	// Create table manually for SQLite compatibility
// 	err = db.Exec(`
// 		CREATE TABLE bookings (
// 			id TEXT PRIMARY KEY,
// 			user_id TEXT,
// 			hotel_id TEXT,
// 			room_id TEXT,
// 			check_in_date DATETIME,
// 			check_out_date DATETIME,
// 			total_price REAL,
// 			is_cancelled INTEGER DEFAULT 0
// 		)
// 	`).Error
// 	assert.NoError(t, err)
// 	repo := repository.NewBookingRepository(db)
// 	hotelRepo := repository.NewHotelRepository(db)
// 	service := NewBookingService(hotelRepo, repo)

// 	booking := &domain.Booking{
// 		Id:           uuid.New(),
// 		UserId:       uuid.New(),
// 		HotelId:      uuid.New(),
// 		RoomId:       uuid.New(),
// 		CheckInDate:  time.Now(),
// 		CheckOutDate: time.Now().AddDate(0, 0, 3),
// 		TotalPrice:   300.0,
// 		IsCancelled:  false,
// 	}
// 	service.CreateBooking(&domain.CreateBookingRequest{
// 		HotelId:      booking.HotelId,
// 		UserId:       booking.UserId,
// 		RoomId:       booking.RoomId,
// 		CheckInDate:  booking.CheckInDate,
// 		CheckOutDate: booking.CheckOutDate,
// 	})

// 	bookings, err := service.GetBookingsByUserId(booking.UserId.String())
// 	assert.NoError(t, err)
// 	assert.Len(t, bookings, 1)
// 	assert.Equal(t, booking.Id, bookings[0].Id)
// 	assert.Equal(t, booking.TotalPrice, bookings[0].TotalPrice)
// 	assert.Equal(t, booking.IsCancelled, bookings[0].IsCancelled)
// 	assert.Equal(t, booking.CheckInDate, bookings[0].CheckInDate)
// 	assert.Equal(t, booking.CheckOutDate, bookings[0].CheckOutDate)
// }

func TestBookingService_CreateBooking(t *testing.T) {
	tests := []struct {
		name          string
		hotelId       uuid.UUID
		roomId        uuid.UUID
		userId        uuid.UUID
		roomPrice     float64
		checkIn       time.Time
		checkOut      time.Time
		expectedPrice float64
		shouldError   bool
	}{
		{
			name:          "successful booking creation",
			hotelId:       uuid.New(),
			roomId:        uuid.New(),
			userId:        uuid.New(),
			roomPrice:     100.0,
			checkIn:       time.Now(),
			checkOut:      time.Now().AddDate(0, 0, 3),
			expectedPrice: 300.0, // 100 * 3 days
			shouldError:   false,
		},
		{
			name:          "cancelled booking",
			hotelId:       uuid.New(),
			roomId:        uuid.New(),
			userId:        uuid.New(),
			roomPrice:     100.0,
			checkIn:       time.Now(),
			checkOut:      time.Now().AddDate(0, 0, 5),
			expectedPrice: 500.0, // 100 * 5 days
			shouldError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupBookingServiceTestDB(t)

			// Create hotel and room before creating booking
			err := createTestHotelAndRoom(db, tt.hotelId, tt.roomId, tt.roomPrice)
			assert.NoError(t, err)

			repo := repository.NewBookingRepository(db)
			hotelRepo := repository.NewHotelRepository(db)
			service := NewBookingService(hotelRepo, repo)
			err = service.CreateBooking(&domain.CreateBookingRequest{
				HotelId:      tt.hotelId,
				UserId:       tt.userId,
				RoomId:       tt.roomId,
				CheckInDate:  tt.checkIn,
				CheckOutDate: tt.checkOut,
			})

			if tt.shouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				// Verify booking was created by querying with user_id, hotel_id, and room_id
				var createdBooking domain.Booking
				result := db.Where("user_id = ? AND hotel_id = ? AND room_id = ?",
					tt.userId.String(), tt.hotelId.String(), tt.roomId.String()).
					First(&createdBooking)
				assert.NoError(t, result.Error)
				// Use InDelta for floating point comparison to handle precision issues
				assert.InDelta(t, tt.expectedPrice, createdBooking.TotalPrice, 0.01)
			}
		})
	}
}

func TestBookingService_GetAllBookings(t *testing.T) {
	db := setupBookingServiceTestDB(t)

	// Create test hotels and rooms
	hotel1Id := uuid.New()
	room1Id := uuid.New()
	hotel2Id := uuid.New()
	room2Id := uuid.New()

	err := createTestHotelAndRoom(db, hotel1Id, room1Id, 100.0)
	assert.NoError(t, err)
	err = createTestHotelAndRoom(db, hotel2Id, room2Id, 100.0)
	assert.NoError(t, err)

	repo := repository.NewBookingRepository(db)
	hotelRepo := repository.NewHotelRepository(db)
	service := NewBookingService(hotelRepo, repo)

	// Create test bookings
	userId1 := uuid.New()
	userId2 := uuid.New()
	checkIn1 := time.Now()
	checkOut1 := time.Now().AddDate(0, 0, 2)
	checkIn2 := time.Now()
	checkOut2 := time.Now().AddDate(0, 0, 3)

	err = service.CreateBooking(&domain.CreateBookingRequest{
		HotelId:      hotel1Id,
		UserId:       userId1,
		RoomId:       room1Id,
		CheckInDate:  checkIn1,
		CheckOutDate: checkOut1,
	})
	assert.NoError(t, err)

	err = service.CreateBooking(&domain.CreateBookingRequest{
		HotelId:      hotel2Id,
		UserId:       userId2,
		RoomId:       room2Id,
		CheckInDate:  checkIn2,
		CheckOutDate: checkOut2,
	})
	assert.NoError(t, err)

	bookings, err := service.GetAllBookings()
	assert.NoError(t, err)
	assert.Len(t, bookings, 2)
}

func TestBookingService_GetBookingById(t *testing.T) {
	db := setupBookingServiceTestDB(t)

	// Create test hotel and room
	hotelId := uuid.New()
	roomId := uuid.New()
	userId := uuid.New()

	err := createTestHotelAndRoom(db, hotelId, roomId, 100.0)
	assert.NoError(t, err)

	repo := repository.NewBookingRepository(db)
	hotelRepo := repository.NewHotelRepository(db)
	service := NewBookingService(hotelRepo, repo)

	checkIn := time.Now()
	checkOut := time.Now().AddDate(0, 0, 4)

	err = service.CreateBooking(&domain.CreateBookingRequest{
		HotelId:      hotelId,
		UserId:       userId,
		RoomId:       roomId,
		CheckInDate:  checkIn,
		CheckOutDate: checkOut,
	})
	assert.NoError(t, err)

	// Get the created booking ID by querying by user_id
	bookings, err := service.GetBookingsByUserId(userId.String())
	assert.NoError(t, err)
	assert.Len(t, bookings, 1)
	// Check that the ID is not zero
	assert.NotEqual(t, uuid.Nil, bookings[0].Id)
	bookingId := bookings[0].Id.String()

	tests := []struct {
		name          string
		id            string
		expectedError bool
		expectedPrice float64
	}{
		{
			name:          "existing booking",
			id:            bookingId,
			expectedError: false,
			expectedPrice: 400.0, // 100 * 4 days
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
				// Use InDelta for floating point comparison to handle precision issues
				assert.InDelta(t, tt.expectedPrice, result.TotalPrice, 0.01)
			}
		})
	}
}

func TestNewBookingService(t *testing.T) {
	db := setupBookingServiceTestDB(t)

	repo := repository.NewBookingRepository(db)
	hotelRepo := repository.NewHotelRepository(db)
	service := NewBookingService(hotelRepo, repo)
	assert.NotNil(t, service)
	assert.Equal(t, hotelRepo, service.(*bookingService).hotelRepository)
	assert.Equal(t, repo, service.(*bookingService).bookingRepository)
}
