package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type BookingService interface {
	CreateBooking(booking *domain.Booking) error
	GetAllBookings() ([]domain.Booking, error)
	GetBookingById(id string) (*domain.Booking, error)
}
type bookingService struct {
	bookingRepository repository.BookingRepository
}

func NewBookingService(bookingRepository repository.BookingRepository) BookingService {
	return &bookingService{bookingRepository: bookingRepository}
}

func (s *bookingService) CreateBooking(booking *domain.Booking) error {
	return s.bookingRepository.CreateBooking(booking)
}

func (s *bookingService) GetAllBookings() ([]domain.Booking, error) {
	return s.bookingRepository.GetAllBookings()
}

func (s *bookingService) GetBookingById(id string) (*domain.Booking, error) {
	return s.bookingRepository.GetBookingById(id)
}
