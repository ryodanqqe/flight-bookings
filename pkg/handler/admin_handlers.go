package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryodanqqe/flight-bookings/models/requests"
)

// @Summary Get all flights
// @Tags Admin
// @Description Retrieve a list of all flights
// @Produce json
// @Param token header string true "JWT token" in:header
// @Param secretKey header string true "Secret key" in:header
// @Success 200 {array} models.Flight
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/admin/flights [get]
func (h *Handler) getAllFlights(c *gin.Context) {
	flights, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, flights)

}

// @Summary Get a flight by ID
// @Tags Admin
// @Description Retrieve flight by its ID
// @Param token header string true "JWT token" in:header
// @Param secretKey header string true "Secret key" in:header
// @Param id path string true "Flight ID"
// @Produce json
// @Success 200 {object} models.Flight
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/admin/flights/{id} [get]
func (h *Handler) getFlight(c *gin.Context) {
	id := c.Param("id")

	flight, err := h.services.GetOne(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, flight)
}

// @Summary Create a new flight
// @Tags Admin
// @Description Create a new flight based on the provided data
// @Accept json
// @Param token header string true "JWT token" in:header
// @Param secretKey header string true "Secret key" in:header
// @Param input body requests.CreateFlightRequest true "Flight data"
// @Produce json
// @Success 200 {string} string "id"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/admin/flights [post]
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

// @Summary Update a flight
// @Tags Admin
// @Description Update information about a flight based on the provided data
// @Param id path string true "Flight ID"
// @Accept json
// @Param token header string true "JWT token" in:header
// @Param secretKey header string true "Secret key" in:header
// @Param input body requests.UpdateFlightRequest true "Updated flight data"
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/admin/flights/{id} [put]
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

// @Summary Delete a flight
// @Tags Admin
// @Description Delete a flight by its ID
// @Param token header string true "JWT token" in:header
// @Param secretKey header string true "Secret key" in:header
// @Param id path string true "Flight ID"
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/admin/flights/{id} [delete]
func (h *Handler) deleteFlight(c *gin.Context) {
	id := c.Param("id")

	if err := h.services.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
