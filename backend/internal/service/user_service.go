package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.userRepository.CreateUser(user)
}
