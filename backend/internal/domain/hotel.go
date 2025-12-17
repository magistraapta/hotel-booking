package domain

import "github.com/google/uuid"

type Hotel struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Rooms       []*Room   `gorm:"foreignKey:HotelId" json:"rooms"`
	Address     string    `json:"address" binding:"required"`
	Rating      float64   `json:"rating" binding:"required"`
}

type Room struct {
	Id          uuid.UUID   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id" binding:"required"`
	Size        int         `json:"size" binding:"required"`
	Price       float64     `json:"price" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Available   bool        `json:"available" binding:"required"`
	Facilities  []*Facility `gorm:"many2many:room_facilities;" json:"facilities"`
	HotelId     uuid.UUID   `gorm:"type:uuid" json:"hotel_id" binding:"required"`
	Hotel       *Hotel      `gorm:"foreignKey:HotelId" json:"-"`
}

type Facility struct {
	Id   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id" binding:"required"`
	Name string    `json:"name" binding:"required"`
}
