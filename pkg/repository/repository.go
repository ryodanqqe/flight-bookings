package repository

import (
	"database/sql"

	"github.com/ryodanqqe/flight-bookings/models"
	"github.com/ryodanqqe/flight-bookings/models/requests"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GetUser(email, password string) (models.User, error)
}

type Admin interface {
	Create(flight requests.CreateFlightRequest) (string, error)
	GetOne(id string) (models.Flight, error)
	GetAll() ([]models.Flight, error)
	Update(id string, input requests.UpdateFlightRequest) error
	Delete(id string) error
}

type User interface {
}

type Repository struct {
	Authorization
	Admin
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Admin:         NewAdminPostgres(db),
	}
}
