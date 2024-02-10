package repository

import (
	"database/sql"

	"github.com/ryodanqqe/flight-bookings/models"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GetUser(email, password string) (models.User, error)
}

type Admin interface {
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
	}
}
