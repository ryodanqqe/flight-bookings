package service

import (
	"github.com/ryodanqqe/flight-bookings/models"
	"github.com/ryodanqqe/flight-bookings/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
}

type Admin interface {
}

type User interface {
}

type Service struct {
	Authorization
	Admin
	User
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
