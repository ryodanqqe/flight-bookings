package repository

import (
	"database/sql"

	"github.com/ryodanqqe/flight-bookings/models/requests"
	"github.com/ryodanqqe/flight-bookings/pkg/repository/query"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) UpdateUser(id string, req requests.UpdateUserRequest) error {

	_, err := r.db.Exec(query.UpdateUserQuery, id, req.Email, req.Password, req.Phone)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserPostgres) DeleteUser(id string) error {

	_, err := r.db.Exec(query.DeleteUserQuery, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserPostgres) CheckAvailableTickets(tx *sql.Tx, req requests.BookTicketRequest, query string) (bool, error) {
	var exists int

	err := tx.QueryRow(query, req.FlightID).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return exists == 1, nil
}

func (r *UserPostgres) BookTicket(tx *sql.Tx, query, userID string, req requests.BookTicketRequest) (string, error) {
	var ticketID string

	err := tx.QueryRow(query, req.FlightID, userID).Scan(&ticketID)
	if err != nil {
		return "", err
	}

	return ticketID, nil
}

func (r *UserPostgres) UpdateAvailableTickets(tx *sql.Tx, query string, req requests.BookTicketRequest) error {

	_, err := tx.Exec(query, req.FlightID)
	if err != nil {
		return err
	}

	return nil
}
