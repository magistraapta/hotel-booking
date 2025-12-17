package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type HotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) *HotelRepository {
	return &HotelRepository{db: db}
}

func (r *HotelRepository) CreateHotel(hotel *domain.Hotel) error {
	return r.db.Create(hotel).Error
}

func (r *HotelRepository) GetAllHotels() ([]domain.Hotel, error) {
	var hotels []domain.Hotel
	if err := r.db.Preload("Rooms").Preload("Rooms.Facilities").Find(&hotels).Error; err != nil {
		return nil, err
	}
	return hotels, nil
}

func (r *HotelRepository) GetHotelById(id string) (*domain.Hotel, error) {
	var hotel domain.Hotel
	if err := r.db.Preload("Rooms").Preload("Rooms.Facilities").First(&hotel, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &hotel, nil
}
