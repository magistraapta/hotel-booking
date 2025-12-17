package domain

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	Id           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id" binding:"required"`
	UserId       uuid.UUID `gorm:"type:uuid" json:"user_id" binding:"required"`
	HotelId      uuid.UUID `gorm:"type:uuid" json:"hotel_id" binding:"required"`
	RoomId       uuid.UUID `gorm:"type:uuid" json:"room_id" binding:"required"`
	CheckInDate  time.Time `json:"check_in_date" binding:"required"`
	CheckOutDate time.Time `json:"check_out_date" binding:"required"`
	TotalPrice   float64   `json:"total_price" binding:"required"`
	IsCancelled  bool      `gorm:"default:false" json:"is_cancelled" binding:"required"`
}
