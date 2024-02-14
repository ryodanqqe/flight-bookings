package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ryodanqqe/flight-bookings/models"
	"github.com/ryodanqqe/flight-bookings/pkg/repository"
)

const (
	salt       = "qwe743iqewi7dish"
	signingKey = "s3qhwe&ic9!qwe"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.MapClaims
	UserID string `json: "user_id`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) generatePasswordHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) CreateUser(user models.User) (string, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {

	user, err := s.repo.GetUser(email, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.MapClaims{
			"exp": time.Now().Add(tokenTTL).Unix(),
			"iat": time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {

	token, err := jwt.ParseWithClaims(accessToken, tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("failed to parse token claims")
	}

	return claims.UserID, nil
}
