package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	// header := c.GetHeader(authorizationHeader)
	// if header == "" {
	// 	newErrorResponse(c, http.StatusUnauthorized, "emty auth header")
	// 	return
	// }

	// headerParts := strings.Split(header, " ")
	// if len(headerParts) != 2 {
	// 	newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	// }

	// cookie
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "token cookie not found")
		return
	}

	accessToken := cookie.Value

	userID, err := h.services.ParseToken(accessToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userID", userID)
}
