package models

import "time"

type Flight struct {
	ID                       string    `pg:"type:uuid default:uuid_generate_v4()" json:"id"`
	StartTime                time.Time `json:"start_time" binding:"required"`
	EndTime                  time.Time `json:"end_time" binding:"required"`
	DeparturePoint           string    `json:"departure_point" binding:"required"`
	Destination              string    `json:"destination" binding:"required"`
	EconomyPrice             float64   `json:"economy_price" binding:"required"`
	BusinessPrice            float64   `json:"business_price" binding:"required"`
	DeluxePrice              float64   `json:"deluxe_price" binding:"required"`
	TotalEconomyTickets      uint      `json:"total_economy_tickets" binding:"required"`
	TotalBusinessTickets     uint      `json:"total_business_tickets" binding:"required"`
	TotalDeluxeTickets       uint      `json:"total_deluxe_tickets" binding:"required"`
	AvailableEconomyTickets  uint      `json:"available_economy_tickets" binding:"required"`
	AvailableBusinessTickets uint      `json:"available_business_tickets" binding:"required"`
	AvailableDeluxeTickets   uint      `json:"available_deluxe_tickets" binding:"required"`
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
