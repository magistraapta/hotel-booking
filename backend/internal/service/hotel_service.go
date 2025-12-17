package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type HotelService struct {
	hotelRepository *repository.HotelRepository
}

func NewHotelService(hotelRepository *repository.HotelRepository) *HotelService {
	return &HotelService{hotelRepository: hotelRepository}
}

func (s *HotelService) CreateHotel(hotel *domain.Hotel) error {
	return s.hotelRepository.CreateHotel(hotel)
}

func (s *HotelService) GetAllHotels() ([]domain.Hotel, error) {
	return s.hotelRepository.GetAllHotels()
}

func (s *HotelService) GetHotelById(id string) (*domain.Hotel, error) {
	return s.hotelRepository.GetHotelById(id)
}
