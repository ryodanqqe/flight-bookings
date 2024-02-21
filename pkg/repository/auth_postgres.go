package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/ryodanqqe/flight-bookings/models"
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

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User

	query := `SELECT id FROM users WHERE email = $1 AND password = $2;`

	err := r.db.QueryRow(query, email, password).Scan(&user.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
