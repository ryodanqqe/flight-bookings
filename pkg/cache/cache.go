package cache

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/ryodanqqe/flight-bookings/models"
	"github.com/ryodanqqe/flight-bookings/models/requests"
)

type Admin interface {
	SetFlight(flightID string, fl requests.CreateFlightRequest) error
	GetOneFlight(flightID string) (models.Flight, error)
	GetAllFlights() ([]models.Flight, error)
	UpdateFlight(flightID string, fl requests.UpdateFlightRequest) error
	DeleteFlight(flightID string) error
}

type Cache struct {
	Admin
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{
		Admin: NewAdminRedis(client),
	}
}

type CachedFlight struct {
	ID                       string    `json:"id"`
	StartTime                time.Time `json:"start_time"`
	EndTime                  time.Time `json:"end_time"`
	DeparturePoint           string    `json:"departure_point"`
	Destination              string    `json:"destination"`
	EconomyPrice             float64   `json:"economy_price"`
	BusinessPrice            float64   `json:"business_price"`
	DeluxePrice              float64   `json:"deluxe_price"`
	TotalEconomyTickets      uint      `json:"total_economy_tickets"`
	TotalBusinessTickets     uint      `json:"total_business_tickets"`
	TotalDeluxeTickets       uint      `json:"total_deluxe_tickets"`
	AvailableEconomyTickets  uint      `json:"available_economy_tickets"`
	AvailableBusinessTickets uint      `json:"available_business_tickets"`
	AvailableDeluxeTickets   uint      `json:"available_deluxe_tickets"`
}
