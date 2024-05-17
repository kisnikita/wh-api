package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/kisnikita/wh-api/internal/handler"
	"github.com/kisnikita/wh-api/internal/repository"
	"github.com/kisnikita/wh-api/internal/server"
	"github.com/kisnikita/wh-api/internal/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	srv := new(server.Server)
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	if err := srv.Run("8088", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server: %s", err.Error())
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error ocured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error ocured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
