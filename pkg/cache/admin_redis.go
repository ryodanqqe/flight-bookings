package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/ryodanqqe/flight-bookings/models"
	"github.com/ryodanqqe/flight-bookings/models/requests"
)

type AdminCache struct {
	client *redis.Client
}

func NewAdminRedis(client *redis.Client) *AdminCache {
	return &AdminCache{client: client}
}

func (c *AdminCache) SetFlight(flightID string, fl requests.CreateFlightRequest) error {

	cachedFlight := CachedFlight{
		ID:                       flightID,
		StartTime:                fl.StartTime,
		EndTime:                  fl.EndTime,
		DeparturePoint:           fl.DeparturePoint,
		Destination:              fl.Destination,
		EconomyPrice:             fl.EconomyPrice,
		BusinessPrice:            fl.BusinessPrice,
		DeluxePrice:              fl.DeluxePrice,
		TotalEconomyTickets:      fl.TotalEconomyTickets,
		TotalBusinessTickets:     fl.TotalBusinessTickets,
		TotalDeluxeTickets:       fl.TotalDeluxeTickets,
		AvailableEconomyTickets:  fl.AvailableEconomyTickets,
		AvailableBusinessTickets: fl.AvailableBusinessTickets,
		AvailableDeluxeTickets:   fl.AvailableDeluxeTickets,
	}

	data, err := json.Marshal(cachedFlight)
	if err != nil {
		return err
	}

	err = c.client.Set(flightID, data, time.Hour*24).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *AdminCache) GetOneFlight(flightID string) (models.Flight, error) {

	data, err := c.client.Get(flightID).Bytes()
	if err != nil {
		if err == redis.Nil {
			return models.Flight{}, nil
		}
		return models.Flight{}, fmt.Errorf("error occurred while getting flight %s from cache: %w", flightID, err)
	}

	var cachedFlight CachedFlight
	err = json.Unmarshal(data, &cachedFlight)
	if err != nil {
		return models.Flight{}, fmt.Errorf("error occurred while unmarshalling flight data from cache: %w", err)
	}

	flight := models.Flight{
		ID:                       cachedFlight.ID,
		StartTime:                cachedFlight.StartTime,
		EndTime:                  cachedFlight.EndTime,
		DeparturePoint:           cachedFlight.DeparturePoint,
		Destination:              cachedFlight.Destination,
		EconomyPrice:             cachedFlight.EconomyPrice,
		BusinessPrice:            cachedFlight.BusinessPrice,
		DeluxePrice:              cachedFlight.DeluxePrice,
		TotalEconomyTickets:      cachedFlight.TotalEconomyTickets,
		TotalBusinessTickets:     cachedFlight.TotalBusinessTickets,
		TotalDeluxeTickets:       cachedFlight.TotalDeluxeTickets,
		AvailableEconomyTickets:  cachedFlight.AvailableEconomyTickets,
		AvailableBusinessTickets: cachedFlight.AvailableBusinessTickets,
		AvailableDeluxeTickets:   cachedFlight.AvailableDeluxeTickets,
	}

	return flight, nil
}

func (c *AdminCache) GetAllFlights() ([]models.Flight, error) {
	var flights []models.Flight

	keys, err := c.client.Keys("*").Result()
	if err != nil {
		return nil, fmt.Errorf("error occurred while retrieving keys from cache: %w", err)
	}

	for _, key := range keys {

		data, err := c.client.Get(key).Bytes()
		if err != nil {
			continue
		}

		var flight models.Flight
		err = json.Unmarshal(data, &flight)
		if err != nil {
			continue
		}

		flights = append(flights, flight)
	}

	return flights, nil
}

func (c *AdminCache) UpdateFlight(flightID string, fl requests.UpdateFlightRequest) error {
	cachedFlight := CachedFlight{
		ID:                       flightID,
		StartTime:                fl.StartTime,
		EndTime:                  fl.EndTime,
		DeparturePoint:           fl.DeparturePoint,
		Destination:              fl.Destination,
		EconomyPrice:             fl.EconomyPrice,
		BusinessPrice:            fl.BusinessPrice,
		DeluxePrice:              fl.DeluxePrice,
		TotalEconomyTickets:      fl.TotalEconomyTickets,
		TotalBusinessTickets:     fl.TotalBusinessTickets,
		TotalDeluxeTickets:       fl.TotalDeluxeTickets,
		AvailableEconomyTickets:  fl.AvailableEconomyTickets,
		AvailableBusinessTickets: fl.AvailableBusinessTickets,
		AvailableDeluxeTickets:   fl.AvailableDeluxeTickets,
	}

	data, err := json.Marshal(cachedFlight)
	if err != nil {
		return err
	}

	err = c.client.Set(flightID, data, time.Hour*24).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *AdminCache) DeleteFlight(flightID string) error {

	err := c.client.Del(flightID).Err()
	if err != nil {
		return fmt.Errorf("error occurred while deleting flight %s from cache: %w", flightID, err)
	}

	return nil
}
