package main

import (
	"log"

	"github.com/ryodanqqe/flight-bookings/pkg/handler"
	"github.com/ryodanqqe/flight-bookings/pkg/repository"
	"github.com/ryodanqqe/flight-bookings/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"), // перенести в env var
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)

	service := service.NewService(*repo)

	handler := handler.NewHandler(service)

	// go func() {
	// 	if err := handler.InitRoutes().Run("localhost:8080"); err != nil {
	// 		logrus.Fatalf("Failed to run HTTP server: %v", err)
	// 	}
	// }()

	if err := handler.InitRoutes().Run("localhost:8080"); err != nil {
		logrus.Fatalf("Failed to run HTTP server: %v", err)
	}

	log.Print("running server")
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
