package handler

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// TODO: Middleware (logging etc)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.GET("/sign-in", h.signIn)
		auth.DELETE("/sign-out", h.signOut)
	}

	api := router.Group("/api")
	{
		admin := api.Group("/admin")
		{
			admin.GET("/flights", h.getAllFlights)
			admin.GET("/flights/id", h.getFlight)
			admin.POST("/flights/:id", h.createFlight)
			admin.PUT("/flights/:id", h.updateFlight)
			admin.DELETE("/flights/:id", h.deleteFlight)
		}

		user := api.Group("user")
		{
			user.GET("/flights", getAvailableFlights)
			user.POST("/book/:flightID", bookTicket)
			user.GET("/bookings", getUserBookings)
			user.GET("/bookings/:id", getUserBooking)
			user.PUT("/bookings/:id", updateUserBooking)
			user.DELETE("/bookings/:id", deleteUserBooking)
		}
	}

	// Setting routes for Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
