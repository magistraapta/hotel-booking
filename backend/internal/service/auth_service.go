package service

import (
	"backend/internal/auth"
	"backend/internal/domain"
	"backend/internal/repository"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repository.UserRepository
}

type IAuthService interface {
	Login(email string, password string) (domain.LoginResponse, error)
	Register(registerRequest *domain.RegisterRequest) error
}

func NewAuthService(userRepo repository.UserRepository) IAuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(email string, password string) (domain.LoginResponse, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.LoginResponse{}, errors.New("invalid password")
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		return domain.LoginResponse{}, err
	}
	return domain.LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (s *AuthService) Register(registerRequest *domain.RegisterRequest) error {
	user := &domain.User{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return s.userRepo.CreateUser(user)
}
