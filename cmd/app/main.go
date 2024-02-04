package main

import (
	"log"

	h "github.com/ryodanqqe/flight-bookings/pkg/handler"
)

func main() {

	// Router setup
	router := h.NewHandler()

	go func() {
		if err := router.InitRoutes().Run("0.0.0.0:8080"); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
	}()

	log.Print("running server")
}
