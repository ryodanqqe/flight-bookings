package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryodanqqe/flight-bookings/cmd/server"
	"github.com/ryodanqqe/flight-bookings/pkg/handler"
	"github.com/ryodanqqe/flight-bookings/pkg/repository"
	"github.com/ryodanqqe/flight-bookings/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Flight Bookings
// @version 1.0
// @description API Server for Flight Bookings

// @host localhost:8080
// @BasePath /

// @CookieParam token string true "JWT token"
// @CookieParam secretKey string true "Secret Key"
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

	service := service.NewService(repo)

	handler := handler.NewHandler(service)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running server")
		}
	}()
	logrus.Printf("Server Started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logrus.Printf("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("error occurred during server shutdown: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on db connection close: %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
