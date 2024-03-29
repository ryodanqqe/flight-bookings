package repository

import (
	"database/sql"
	"time"

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
	BeginTransaction() (*sql.Tx, error)
	CommitTransaction(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
	UpdateUser(id string, req requests.UpdateUserRequest) error
	DeleteUser(id string) error
	CheckAvailableTickets(tx *sql.Tx, req requests.BookTicketRequest, query string) (bool, error)
	BookTicket(tx *sql.Tx, query, userID string, req requests.BookTicketRequest) (string, error)
	UpdateAvailableTickets(tx *sql.Tx, query, flightID string) error
	GetStartTime(flightID string) (time.Time, error)
	GetUserBookings(userID string) ([]models.Ticket, error)
	GetOneUserBooking(ticketID string) (models.Ticket, error)
	UpdateUserBooking(ticketID string, req requests.UpdateUserBookingRequest) error
	DeleteUserBooking(ticketID string) error
	GetFlightIDAndRank(tx *sql.Tx, ticketID string) (string, string, error)
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
		User:          NewUserPostgres(db),
	}
}
