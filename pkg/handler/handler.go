package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// TODO: Middleware (logging etc)

	// Swagger route

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.GET("/sign-in")
		auth.DELETE("/sign-out")
	}

	api := router.Group("/api")
	{
		admin := api.Group("/admin")
		{
			admin.GET("/flights")
			admin.GET("/flights/id")
			admin.POST("/flights/:id")
			admin.PUT("/flights/:id")
			admin.DELETE("/flights/:id")
		}

		user := api.Group("user")
		{
			user.GET("/flights")
			user.POST("/flights/:flightID")
			user.GET("/bookings")
			user.GET("/bookings/:id")
			user.PUT("/bookings/:id")
			user.DELETE("/bookings/:id")
		}
	}

	return router
}
