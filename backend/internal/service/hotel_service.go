package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type HotelService interface {
	CreateHotel(hotel *domain.Hotel) error
	GetAllHotels() ([]domain.Hotel, error)
	GetHotelById(id string) (*domain.Hotel, error)
}

type hotelService struct {
	hotelRepository repository.HotelRepository
}

func NewHotelService(hotelRepository repository.HotelRepository) HotelService {
	return &hotelService{hotelRepository: hotelRepository}
}

func (s *hotelService) CreateHotel(hotel *domain.Hotel) error {
	return s.hotelRepository.CreateHotel(hotel)
}

func (s *hotelService) GetAllHotels() ([]domain.Hotel, error) {
	return s.hotelRepository.GetAllHotels()
}

func (s *hotelService) GetHotelById(id string) (*domain.Hotel, error) {
	return s.hotelRepository.GetHotelById(id)
}
