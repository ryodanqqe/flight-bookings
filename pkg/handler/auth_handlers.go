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

// @Summary SignUp
// @Tags Auth
// @Description Create Account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "userID"
// @Success 200 {string} string "userID"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
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
type signInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags Auth
// @Description Login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var signInInput signInput

	if err := c.ShouldBindJSON(&signInInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(signInInput.Email, signInInput.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie("token", token, int(tokenTTL.Seconds()), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

// @Summary SignOut
// @Tags Auth
// @Description Signout
// @ID signout
// @Accept  json
// @Produce  json
// @Success 200 {string} string "User successfully signed out"
// @Router /auth/sign-out [post]
func (h *Handler) signOut(c *gin.Context) {

	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusOK, gin.H{"message": "User successfully signed out"})
}
