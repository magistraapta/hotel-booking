package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type UserService interface {
	CreateUser(user *domain.User) error
	GetAllUsers() ([]domain.User, error)
	GetUserById(id string) (*domain.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(user *domain.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]domain.User, error) {
	return s.userRepository.GetAllUsers()
}

func (s *userService) GetUserById(id string) (*domain.User, error) {
	return s.userRepository.GetUserById(id)
}
