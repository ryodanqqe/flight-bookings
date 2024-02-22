package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
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
