package service

import (
	"github.com/ryodanqqe/flight-bookings/models"
	"github.com/ryodanqqe/flight-bookings/models/requests"
	"github.com/ryodanqqe/flight-bookings/pkg/cache"
	"github.com/ryodanqqe/flight-bookings/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Admin interface {
	Create(flight requests.CreateFlightRequest) (string, error)
	GetOne(id string) (models.Flight, error)
	GetAll() ([]models.Flight, error)
	Update(id string, req requests.UpdateFlightRequest) error
	Delete(id string) error
}

type User interface {
	UpdateUser(id string, req requests.UpdateUserRequest) error
	DeleteUser(id string) error
	BookTicket(userID string, req requests.BookTicketRequest) (string, error)
	GetUserBookings(userID string) ([]models.Ticket, error)
	GetOneUserBooking(ticketID string) (models.Ticket, error)
	UpdateUserBooking(ticketID string, req requests.UpdateUserBookingRequest) error
	DeleteUserBooking(ticketID string) error
}

type Service struct {
	Authorization
	Admin
	User
}

func NewService(repos *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Admin:         NewAdminService(repos.Admin, cache.Admin),
		User:          NewUserService(repos.User),
	}
}
