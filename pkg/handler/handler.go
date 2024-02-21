package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ryodanqqe/flight-bookings/pkg/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// TODO: Middleware (logging etc)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-out", h.signOut)
	}

	api := router.Group("/api", h.userIdentity)
	{
		admin := api.Group("/admin")
		{
			admin.POST("/flights", h.createFlight)
			admin.GET("/flights", h.getAllFlights)
			admin.GET("/flights/:id", h.getFlight)
			admin.PUT("/flights/:id", h.updateFlight)
			admin.DELETE("/flights/:id", h.deleteFlight)
		}

		user := api.Group("user")
		{
			user.PUT("/account/:id", h.updateUser)
			user.DELETE("/account/:id", h.deleteUser)
			user.GET("/flights", h.getAvailableFlights)
			user.POST("/book/:flightID", h.bookTicket)
			user.GET("/bookings", h.getUserBookings)
			user.GET("/bookings/:id", h.getUserBooking)
			user.PUT("/bookings/:id", h.updateUserBooking)
			user.DELETE("/bookings/:id", h.deleteUserBooking)
		}
	}

	// Setting routes for Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
