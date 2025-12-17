package repository

import (
	"backend/internal/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetAllUsers() ([]domain.User, error)
	GetUserById(id string) (*domain.User, error)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *domain.User) error {
	user.Id = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserById(id string) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
