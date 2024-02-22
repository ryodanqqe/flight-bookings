package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryodanqqe/flight-bookings/models/requests"
)

func (h *Handler) getAvailableFlights(c *gin.Context) {
	flights, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, flights)
}

func (h *Handler) bookTicket(c *gin.Context) {

	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input requests.BookTicketRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ticketID, err := h.services.BookTicket(userID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"ticket_id": ticketID})
}

func (h *Handler) getUserBookings(c *gin.Context) {

	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tickets, err := h.services.GetUserBookings(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func (h *Handler) getUserBooking(c *gin.Context) {

	ticketID := c.Param("id")

	if ticketID == "" {
        newErrorResponse(c, http.StatusBadRequest, "Ticket ID is required")
        return
    }

	ticket, err := h.services.GetOneUserBooking(ticketID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// Для передачи билета другому пользователю
func (h *Handler) updateUserBooking(c *gin.Context) {
	var input requests.UpdateUserBookingRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ticketID := c.Param("id")
	if ticketID == "" {
        newErrorResponse(c, http.StatusBadRequest, "Ticket ID is required")
        return
    }

	if err := h.services.UpdateUserBooking(ticketID, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// Flights (AvailableTickets +1)
func (h *Handler) deleteUserBooking(c *gin.Context) {

	ticketID := c.Param("id")
	if ticketID == "" {
        newErrorResponse(c, http.StatusBadRequest, "Ticket ID is required")
        return
    }

	if err := h.services.DeleteUserBooking(ticketID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})

}

func (h *Handler) updateUser(c *gin.Context) {
	var input requests.UpdateUserRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	if id == "" {
        newErrorResponse(c, http.StatusBadRequest, "ID is required")
        return
    }

	if err := h.services.UpdateUser(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
        newErrorResponse(c, http.StatusBadRequest, "ID is required")
        return
    }

	if err := h.services.DeleteUser(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
