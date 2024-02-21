package service

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/ryodanqqe/flight-bookings/models/requests"
	"github.com/ryodanqqe/flight-bookings/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UpdateUser(id string, req requests.UpdateUserRequest) error {
	return s.repo.UpdateUser(id, req)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) BookTicket(userID string, req requests.BookTicketRequest) (string, error) {

	ok := validateRank(req.Rank)
	if !ok {
		return "", errors.New("invalid ticket rank")
	}

	tx, err := s.repo.BeginTransaction()
	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			if rollbackErr := s.repo.Rollback(tx); rollbackErr != nil {
				log.Printf("error rolling back transaction: %v", rollbackErr)
			}
		}
	}()

	// Проверка доступности билетов
	checkQuery := getQueryForRank(req.Rank, "check")
	available, err := s.repo.CheckAvailableTickets(tx, req, checkQuery)

	if err != nil {
		return "", err
	}

	if !available {
		return "", errors.New("tickets are not available")
	}

	// Резервирование билета
	reserveQuery := getQueryForRank(req.Rank, "reserve")
	ticketID, err := s.repo.BookTicket(tx, reserveQuery, userID, req)
	if err != nil {
		return "", err
	}

	// Обновление количества доступных билетов (Available -1)
	updateQuery := getQueryForRank(req.Rank, "update")
	if err := s.repo.UpdateAvailableTickets(tx, updateQuery, req); err != nil {
		return "", err
	}

	if err := s.repo.CommitTransaction(tx); err != nil {
		return "", err
	}

	return ticketID, nil
}

// return check, reserve, update
func getQueryForRank(rank string, operation string) string {
	switch operation {
	case "check":
		return fmt.Sprintf("SELECT 1 FROM flights WHERE ID = $1 AND Available%sTickets > 0", rank)
	case "reserve":
		return fmt.Sprintf("INSERT INTO tickets (FlightID, UserID, Rank, Price) VALUES ($1, $2, '%s', (SELECT %sPrice FROM flights WHERE ID = $1)) RETURNING ID", rank, rank)
	case "update":
		return fmt.Sprintf("UPDATE flights SET Available%sTickets = Available%sTickets - 1 WHERE ID = $1 AND Available%sTickets >= 0", rank, rank, rank)
	default:
		return ""
	}
}

func validateRank(rank string) bool {
	rank = strings.ToLower(rank)
	switch rank {
	case "economy", "business", "deluxe":
		return true
	default:
		return false
	}
}
