package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryodanqqe/flight-bookings/models/requests"
)

func (h *Handler) getAllFlights(c *gin.Context) {
	flights, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, flights)

}

func (h *Handler) getFlight(c *gin.Context) {
	id := c.Param("id")

	flight, err := h.services.GetOne(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, flight)
}

func (h *Handler) createFlight(c *gin.Context) {
	var input requests.CreateFlightRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) updateFlight(c *gin.Context) {
	var input requests.UpdateFlightRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")

	err := h.services.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteFlight(c *gin.Context) {
	id := c.Param("id")

	if err := h.services.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
