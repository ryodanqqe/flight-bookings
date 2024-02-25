package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryodanqqe/flight-bookings/models/requests"
)

// @Summary Get available flights
// @Tags User
// @Description Retrieve a list of available flights
// @Produce json
// @Param token header string true "JWT token" in:header
// @Success 200 {array} models.Flight
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/flights [get]
func (h *Handler) getAvailableFlights(c *gin.Context) {
	flights, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, flights)
}

// @Summary Book a ticket
// @Tags User
// @Description Book a ticket for a flight
// @Accept json
// @Param token header string true "JWT token" in:header
// @Param input body requests.BookTicketRequest true "Ticket booking request data"
// @Produce json
// @Success 200 {string} string "ticket_id"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/bookings/book [post]
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

// @Summary Get user bookings
// @Tags User
// @Description Retrieve a list of bookings made by the user
// @Param token header string true "JWT token" in:header
// @Produce json
// @Success 200 {array} models.Ticket
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/bookings [get]
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

// @Summary Get a user booking by ID
// @Tags User
// @Description Retrieve information about a specific booking made by the user
// @Param token header string true "JWT token" in:header
// @Param id path string true "Booking ID"
// @Produce json
// @Success 200 {object} models.Ticket
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/bookings/{id} [get]
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

// @Summary Update a user booking
// @Tags User
// @Description Update information about a booking made by the user
// @Accept json
// @Param token header string true "JWT token" in:header
// @Param id path string true "Booking ID"
// @Param input body requests.UpdateUserBookingRequest true "Updated booking data"
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/bookings/{id} [put]
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

// @Summary Delete a user booking
// @Tags User
// @Description Delete a booking made by the user
// @Param token header string true "JWT token" in:header
// @Param id path string true "Booking ID"
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/bookings/{id} [delete]
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

// @Summary Update a user
// @Tags User
// @Description Update information about the user
// @Accept json
// @Param token header string true "JWT token" in:header
// @Param id path string true "User ID"
// @Param input body requests.UpdateUserRequest true "Updated data"
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/account/{id} [put]
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

// @Summary Delete a user
// @Tags User
// @Description Delete the user account
// @Accept json
// @Param token header string true "JWT token" in:header
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/account/{id} [delete]
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
