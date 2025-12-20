package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *domain.Booking) error
	GetAllBookings() ([]domain.Booking, error)
	GetBookingById(id string) (*domain.Booking, error)
	GetBookingsByUserId(userId string) ([]domain.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(booking *domain.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetAllBookings() ([]domain.Booking, error) {
	var bookings []domain.Booking
	if err := r.db.Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepository) GetBookingsByUserId(userId string) ([]domain.Booking, error) {
	var bookings []domain.Booking
	if err := r.db.Where("user_id = ?", userId).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepository) GetBookingById(id string) (*domain.Booking, error) {
	var booking domain.Booking
	if err := r.db.First(&booking, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}
