package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"errors"

	"github.com/google/uuid"
)

type BookingService interface {
	CreateBooking(request *domain.CreateBookingRequest) error
	GetAllBookings() ([]domain.Booking, error)
	GetBookingById(id string) (*domain.Booking, error)
	GetBookingsByUserId(userId string) ([]domain.Booking, error)
}
type bookingService struct {
	hotelRepository   repository.HotelRepository
	bookingRepository repository.BookingRepository
}

func NewBookingService(hotelRepository repository.HotelRepository, bookingRepository repository.BookingRepository) BookingService {
	return &bookingService{hotelRepository: hotelRepository, bookingRepository: bookingRepository}
}

/*
CreateBooking
Params: CreateBookingRequest
Returns: error
Description: Create a new booking
*/
func (s *bookingService) CreateBooking(request *domain.CreateBookingRequest) error {
	// get room price
	hotel, err := s.hotelRepository.GetHotelById(request.HotelId.String())
	if err != nil {
		return err
	}
	var targetRoom *domain.Room

	for _, room := range hotel.Rooms {
		if room.Id == request.RoomId {
			targetRoom = room
			break
		}
	}

	if targetRoom == nil {
		return errors.New("room not found")
	}

	if !targetRoom.Available {
		return errors.New("room is not available")
	}

	if request.CheckOutDate.Before(request.CheckInDate) {
		return errors.New("check out date must be after check in date")
	}

	numberOfDays := request.CheckOutDate.Sub(request.CheckInDate).Hours() / 24
	totalPrice := targetRoom.Price * numberOfDays

	booking := &domain.Booking{
		Id:           uuid.New(),
		UserId:       request.UserId,
		HotelId:      request.HotelId,
		RoomId:       request.RoomId,
		CheckInDate:  request.CheckInDate,
		CheckOutDate: request.CheckOutDate,
		TotalPrice:   totalPrice,
		IsCancelled:  false,
	}
	return s.bookingRepository.CreateBooking(booking)
}

func (s *bookingService) GetBookingsByUserId(userId string) ([]domain.Booking, error) {
	return s.bookingRepository.GetBookingsByUserId(userId)
}

func (s *bookingService) GetAllBookings() ([]domain.Booking, error) {
	return s.bookingRepository.GetAllBookings()
}

func (s *bookingService) GetBookingById(id string) (*domain.Booking, error) {
	return s.bookingRepository.GetBookingById(id)
}
