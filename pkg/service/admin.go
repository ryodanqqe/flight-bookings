package service

import (
	"github.com/ryodanqqe/flight-bookings/models"
	"github.com/ryodanqqe/flight-bookings/models/requests"
	"github.com/ryodanqqe/flight-bookings/pkg/repository"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) Create(flight requests.CreateFlightRequest) (string, error) {
	return s.repo.Create(flight)
}

func (s *AdminService) GetOne(id string) (models.Flight, error) {
	return s.repo.GetOne(id)
}

func (s *AdminService) GetAll() ([]models.Flight, error) {
	return s.repo.GetAll()
}

func (s *AdminService) Update(id string, req requests.UpdateFlightRequest) error {
	return s.repo.Update(id, req)
}

func (s *AdminService) Delete(id string) error {
	return s.repo.Delete(id)
}

// TODO: Переопределить структуры реквестов/респонсов в отдельный пакет
