package repository

import (
	"database/sql"
	"time"

	"github.com/ryodanqqe/flight-bookings/models"
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

func (r *UserPostgres) UpdateAvailableTickets(tx *sql.Tx, query, flightID string) error {

	_, err := tx.Exec(query, flightID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserPostgres) GetStartTime(flightID string) (time.Time, error) {
	var startTime time.Time

	err := r.db.QueryRow("SELECT StartTime FROM flights WHERE id = $1", flightID).Scan(&startTime)
	if err != nil {
		return time.Time{}, err
	}

	return startTime, nil
}

func (r *UserPostgres) GetUserBookings(userID string) ([]models.Ticket, error) {

	var result []models.Ticket

	query := `SELECT * FROM tickets WHERE UserID = $1`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return []models.Ticket{}, err
	}

	for rows.Next() {
		var ticket models.Ticket

		err := rows.Scan(&ticket.ID, &ticket.FlightID, &ticket.UserID, &ticket.Rank, &ticket.Price, &ticket.CreatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, ticket)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserPostgres) GetOneUserBooking(ticketID string) (models.Ticket, error) {

	var ticket models.Ticket

	query := `SELECT * FROM tickets WHERE ID = $1`

	err := r.db.QueryRow(query, ticketID).Scan(&ticket.ID, &ticket.FlightID, &ticket.UserID, &ticket.Rank, &ticket.Price, &ticket.CreatedAt)
	if err != nil {
		return models.Ticket{}, err
	}

	return ticket, nil
}

func (r *UserPostgres) UpdateUserBooking(ticketID string, req requests.UpdateUserBookingRequest) error {

	query := `UPDATE tickets SET UserID = $2 WHERE ID = $1`

	_, err := r.db.Exec(query, ticketID, req.NewUserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserPostgres) DeleteUserBooking(ticketID string) error {

	query := "DELETE FROM tickets WHERE ID = $1"

	_, err := r.db.Exec(query, ticketID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserPostgres) GetFlightIDAndRank(tx *sql.Tx, ticketID string) (string, string, error) {

	var flightID, rank string

	query := "SELECT FlightID, Rank FROM tickets WHERE ID = $1"

	row := tx.QueryRow(query, ticketID)
	if err := row.Scan(&flightID, &rank); err != nil {
		return "", "", err
	}

	return flightID, rank, nil
}
