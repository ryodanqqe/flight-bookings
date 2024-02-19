package repository

import (
	"database/sql"

	"github.com/ryodanqqe/flight-bookings/models"
	"github.com/ryodanqqe/flight-bookings/models/requests"
	"github.com/ryodanqqe/flight-bookings/pkg/repository/query"
)

type AdminPostgres struct {
	db *sql.DB
}

func NewAdminPostgres(db *sql.DB) *AdminPostgres {
	return &AdminPostgres{db: db}
}

func (r *AdminPostgres) Create(fl requests.CreateFlightRequest) (string, error) {
	var id string

	row := r.db.QueryRow(query.CreateFlightQuery, fl.StartTime, fl.EndTime, fl.DeparturePoint,
		fl.Destination, fl.EconomyPrice, fl.BusinessPrice, fl.DeluxePrice, fl.TotalEconomyTickets,
		fl.TotalBusinessTickets, fl.TotalDeluxeTickets, fl.AvailableEconomyTickets, fl.AvailableBusinessTickets,
		fl.AvailableDeluxeTickets)

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *AdminPostgres) GetOne(id string) (models.Flight, error) {
	var result models.Flight

	row := r.db.QueryRow(query.GetOneFlightQuery, id)

	if err := row.Scan(&result.ID, &result.StartTime, &result.EndTime, &result.DeparturePoint, &result.Destination,
		&result.EconomyPrice, &result.BusinessPrice, &result.DeluxePrice, &result.TotalEconomyTickets,
		&result.TotalBusinessTickets, &result.TotalDeluxeTickets, &result.AvailableEconomyTickets,
		&result.AvailableBusinessTickets, &result.AvailableDeluxeTickets, &result.CreatedAt); err != nil {
		return models.Flight{}, err
	}

	return result, nil
}

func (r *AdminPostgres) GetAll() ([]models.Flight, error) {
	var result []models.Flight

	rows, err := r.db.Query(query.GetAllFlightQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var flight models.Flight
		if err := rows.Scan(&flight.ID, &flight.StartTime, &flight.EndTime, &flight.DeparturePoint, &flight.Destination,
			&flight.EconomyPrice, &flight.BusinessPrice, &flight.DeluxePrice, &flight.TotalEconomyTickets,
			&flight.TotalBusinessTickets, &flight.TotalDeluxeTickets, &flight.AvailableEconomyTickets,
			&flight.AvailableBusinessTickets, &flight.AvailableDeluxeTickets, &flight.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, flight)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *AdminPostgres) Update(id string, inp requests.UpdateFlightRequest) error {

	_, err := r.db.Exec(query.UpdateFlightQuery, id, inp.StartTime, inp.EndTime, inp.DeparturePoint,
		inp.Destination, inp.EconomyPrice, inp.BusinessPrice, inp.DeluxePrice, inp.TotalEconomyTickets,
		inp.TotalBusinessTickets, inp.TotalDeluxeTickets, inp.AvailableEconomyTickets, inp.AvailableBusinessTickets,
		inp.AvailableDeluxeTickets)

	if err != nil {
		return err
	}

	return nil
}

func (r *AdminPostgres) Delete(id string) error {

	_, err := r.db.Exec(query.DeleteFlightQuery, id)
	if err != nil {
		return err
	}

	return nil
}
