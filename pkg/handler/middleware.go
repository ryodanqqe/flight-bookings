package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {

	cookie, err := c.Cookie("token")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "token cookie not found")
		return
	}

	userID, err := h.services.ParseToken(cookie)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("user_id", userID)
}

func (h *Handler) adminIdentity(c *gin.Context) {

	cookie, err := c.Cookie("token")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "token cookie not found")
		return
	}

	userID, err := h.services.ParseToken(cookie)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	secretKey, err := c.Cookie("secretKey")
	if err != nil || secretKey != "extremelysecurekey" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid secret key")
		return
	}

	c.Set("user_id", userID)
}

func (h *Handler) getUserID(c *gin.Context) (string, error) {

	id, ok := c.Get("user_id")
	if !ok {
		return "", errors.New("user id not found")
	}

	idStr, ok := id.(string)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}

	return idStr, nil
}
