package handler

import (
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

	c.Set("userID", userID)
}
