package service

import (
	"github.com/ryodanqqe/flight-bookings/models"
	"github.com/ryodanqqe/flight-bookings/models/requests"
	"github.com/ryodanqqe/flight-bookings/pkg/cache"
	"github.com/ryodanqqe/flight-bookings/pkg/repository"
	"github.com/sirupsen/logrus"
)

type AdminService struct {
	repo  repository.Admin
	cache cache.Admin
}

func NewAdminService(repo repository.Admin, cache cache.Admin) *AdminService {
	return &AdminService{
		repo:  repo,
		cache: cache,
	}
}

func (s *AdminService) Create(flight requests.CreateFlightRequest) (string, error) {

	flightID, err := s.repo.Create(flight)
	if err != nil {
		return "", err
	}

	_ = s.cache.SetFlight(flightID, flight)

	return flightID, nil
}

func (s *AdminService) GetOne(id string) (models.Flight, error) {

	cachedFlight, err := s.cache.GetOneFlight(id)
	if err != nil {
		return models.Flight{}, err
	}

	if cachedFlight.ID != "" {
		return cachedFlight, nil
	}

	flight, err := s.repo.GetOne(id)
	if err != nil {
		return models.Flight{}, err
	}

	if err = s.cache.SetFlight(id, convertFlightToCreateFlightRequest(flight)); err != nil {
		logrus.Println("Error caching flight:", err)
	}

	return flight, nil
}

func (s *AdminService) GetAll() ([]models.Flight, error) {

	flights, err := s.cache.GetAllFlights()
	if err != nil {
		return nil, err
	}

	if len(flights) > 0 {
		return flights, nil
	}

	flights, err = s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, flight := range flights {
		err := s.cache.SetFlight(flight.ID, convertFlightToCreateFlightRequest(flight))
		if err != nil {
			logrus.Println("Error caching flight:", err)
		}
	}

	return flights, nil
}

func (s *AdminService) Update(id string, req requests.UpdateFlightRequest) error {
	if err := s.repo.Update(id, req); err != nil {
		return err
	}

	if err := s.cache.UpdateFlight(id, req); err != nil {
		return err
	}

	return nil

}

func (s *AdminService) Delete(id string) error {

	if err := s.repo.Delete(id); err != nil {
		return err
	}

	if err := s.cache.DeleteFlight(id); err != nil {
		return err
	}

	return nil
}

func convertFlightToCreateFlightRequest(flight models.Flight) requests.CreateFlightRequest {
	return requests.CreateFlightRequest{
		StartTime:                flight.StartTime,
		EndTime:                  flight.EndTime,
		DeparturePoint:           flight.DeparturePoint,
		Destination:              flight.Destination,
		EconomyPrice:             flight.EconomyPrice,
		BusinessPrice:            flight.BusinessPrice,
		DeluxePrice:              flight.DeluxePrice,
		TotalEconomyTickets:      flight.TotalEconomyTickets,
		TotalBusinessTickets:     flight.TotalBusinessTickets,
		TotalDeluxeTickets:       flight.TotalDeluxeTickets,
		AvailableEconomyTickets:  flight.AvailableEconomyTickets,
		AvailableBusinessTickets: flight.AvailableBusinessTickets,
		AvailableDeluxeTickets:   flight.AvailableDeluxeTickets,
	}
}
