package models

import "time"

type Flight struct {
	ID                       string    `pg:"type:uuid default:uuid_generate_v4()" json:"id"`
	StartTime                time.Time `json:"start_time"`
	EndTime                  time.Time `json:"end_time"`
	DeparturePoint           string    `json:"departure_point"`
	Destination              string    `json:"destination"`
	EconomyPrice             float64   `json:"economy_price"`
	BusinessPrice            float64   `json:"business_price"`
	DeluxePrice              float64   `json:"deluxe_price"`
	TotalEconomyTickets      int       `json:"total_economy_tickets"`
	TotalBusinessTickets     int       `json:"total_business_tickets"`
	TotalDeluxeTickets       int       `json:"total_deluxe_tickets"`
	AvailableEconomyTickets  int       `json:"available_economy_tickets"`
	AvailableBusinessTickets int       `json:"available_business_tickets"`
	AvailableDeluxeTickets   int       `json:"available_deluxe_tickets"`
	CreatedAt                time.Time `json:"created_at"`
}

type Ticket struct {
	ID        string    `pg:"type:uuid default:uuid_generate_v4()" json:"id"`
	FlightID  string    `json:"flight_id"`
	UserID    string    `json:"user_id"`
	Rank      string    `json:"rank"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
