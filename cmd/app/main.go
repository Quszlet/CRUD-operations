package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Quszlet/CRUD-operations/internal/api/handler"
	"github.com/Quszlet/CRUD-operations/internal/repository"
	"github.com/Quszlet/CRUD-operations/internal/service"
	"github.com/spf13/viper"
)

func main() {
	if err := repository.InitConfig(); err != nil {
		log.Fatalf("Failed init config: %s", err.Error())
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

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	routes := handler.InitRoutes()

	srv := &http.Server{
		Addr:           viper.GetString("port"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(srv.ListenAndServe())
}
