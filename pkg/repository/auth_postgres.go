package repository

import (
	"database/sql"

	"github.com/ryodanqqe/flight-bookings/models"
)

const (
	usersTable   = "users"
	flightsTable = "flights"
	ticketsTable = "tickets"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (string, error) {
	var id string

	query := `INSERT INTO Users (phone, email, password) values ($1, $2, $3) RETURNING ID`

	if err := r.db.QueryRow(query, user.Phone, user.Email, user.Password).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
