package repository

import "gorm.io/gorm"

type AuthRepository struct {
	db *gorm.DB
}

type IAuthRepository interface {
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db: db}
}
