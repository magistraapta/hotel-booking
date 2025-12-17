package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type HotelRepository interface {
	CreateHotel(hotel *domain.Hotel) error
	GetAllHotels() ([]domain.Hotel, error)
	GetHotelById(id string) (*domain.Hotel, error)
}

type hotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) HotelRepository {
	return &hotelRepository{db: db}
}

func (r *hotelRepository) CreateHotel(hotel *domain.Hotel) error {
	return r.db.Create(hotel).Error
}

func (r *hotelRepository) GetAllHotels() ([]domain.Hotel, error) {
	var hotels []domain.Hotel
	if err := r.db.Preload("Rooms").Preload("Rooms.Facilities").Find(&hotels).Error; err != nil {
		return nil, err
	}
	return hotels, nil
}

func (r *hotelRepository) GetHotelById(id string) (*domain.Hotel, error) {
	var hotel domain.Hotel
	if err := r.db.Preload("Rooms").Preload("Rooms.Facilities").First(&hotel, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &hotel, nil
}
