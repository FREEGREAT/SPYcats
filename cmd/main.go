package main

import (
	"context"
	"os"
	"spy-cats/internal/handler"
	middleware "spy-cats/internal/middlewae"

	"spy-cats/internal/server"
	service "spy-cats/internal/services"
	"spy-cats/pkg/logger"
	db_connection "spy-cats/pkg/pg_connection"
	"spy-cats/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {

	log := logger.New("debug")

	if err := utils.InitConfig(); err != nil {
		log.Fatal("Error initializing configs", "error", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env variables", "error", err)
	}

	db, err := db_connection.NewClient(context.Background(), viper.GetInt("postgre.connectTries"), db_connection.StorageConfig{
		Host:     viper.GetString("postgre.host"),
		Port:     viper.GetString("postgre.port"),
		Username: viper.GetString("postgre.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: viper.GetString("postgre.db"),
		SSLMode:  viper.GetString("postgre.ssl"),
	})
	if err != nil {
		log.Fatal("Failed to initialize database", "error", err)
	}

	services := service.NewService(db, log)
	handlers := handler.NewHandler(services, log)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.RequestLogger(log))
	router.Use(middleware.Recovery(log))

	handlers.InitRoutes(router)
	srv := new(server.Server)
	port := viper.GetString("srv.port")
	log.Info("Server is running", "port", port)

	if err := srv.Run(port, router); err != nil {
		log.Fatal("Error occurred while running HTTP server", "error", err)
	}
}
