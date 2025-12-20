package auth

import (
	"backend/internal/domain"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateToken(user *domain.User) (domain.LoginResponse, error) {

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return domain.LoginResponse{}, errors.New("JWT_SECRET is not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"isAdmin": user.IsAdmin,
		"sub":     user.Id.String(),
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
		"type":    "access",
	})

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return domain.LoginResponse{}, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"isAdmin": user.IsAdmin,
		"sub":     user.Id.String(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"type":    "refresh",
	})

	rt, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		AccessToken:  t,
		RefreshToken: rt,
	}, nil
}

func ValidateToken(token string) (domain.User, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return domain.User{}, errors.New("JWT_SECRET is not set")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return domain.User{}, err
	}

	tokenType, ok := claims["type"].(string)

	if !ok {
		return domain.User{}, errors.New("invalid token")
	}

	if tokenType != "access" {
		return domain.User{}, errors.New("invalid token type: access token required")
	}

	userId, ok := claims["sub"].(string)
	isAdmin, ok := claims["isAdmin"].(bool)
	if !ok {
		return domain.User{}, errors.New("invalid token: user id required")
	}

	return domain.User{Id: uuid.MustParse(userId), IsAdmin: isAdmin}, nil
}

// ValidateRefreshToken validates a refresh token and returns user information
func ValidateRefreshToken(token string) (domain.User, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return domain.User{}, errors.New("JWT_SECRET is not set")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return domain.User{}, err
	}

	tokenType, ok := claims["type"].(string)

	if !ok {
		return domain.User{}, errors.New("invalid token")
	}

	if tokenType != "refresh" {
		return domain.User{}, errors.New("invalid token type: refresh token required")
	}

	userId, ok := claims["sub"].(string)
	if !ok {
		return domain.User{}, errors.New("invalid token: user id required")
	}

	isAdmin, ok := claims["isAdmin"].(bool)
	if !ok {
		return domain.User{}, errors.New("invalid token: admin status required")
	}

	return domain.User{Id: uuid.MustParse(userId), IsAdmin: isAdmin}, nil
}
