package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id" binding:"required"`
	Username  string    `gorm:"uniqueIndex" json:"username" binding:"required"`
	Email     string    `gorm:"uniqueIndex" json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	IsAdmin   bool      `json:"is_admin" binding:"required"`
}
