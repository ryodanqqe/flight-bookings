package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ryodanqqe/flight-bookings/models"
)

const (
	tokenTTL = 12 * time.Hour
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userID": id,
	})
}

// Можно сделать email/phone
type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var signInInput signInInput

	if err := c.ShouldBindJSON(&signInInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(signInInput.Email, signInInput.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	expiration := time.Now().Add(tokenTTL) // tokenTTL
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (h *Handler) signOut(c *gin.Context) {
	// Кэшировать в Redis токены, и проверять через него | Сделать отдельной структурой с функциями Add, Contains

	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{"message": "User signed out"})
}
