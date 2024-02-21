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

	c.JSON(http.StatusOK, ticketID)
}

func (h *Handler) getUserBookings(c *gin.Context) {

}

func (h *Handler) getUserBooking(c *gin.Context) {

}

func (h *Handler) updateUserBooking(c *gin.Context) {

}

func (h *Handler) deleteUserBooking(c *gin.Context) {

}

func (h *Handler) updateUser(c *gin.Context) {
	var input requests.UpdateUserRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")

	if err := h.services.UpdateUser(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := h.services.DeleteUser(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
