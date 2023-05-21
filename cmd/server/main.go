package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/OneeyK/KPI_GoLang/pkg/db"
	"github.com/OneeyK/KPI_GoLang/pkg/handler"
	"github.com/OneeyK/KPI_GoLang/pkg/repository"
	"github.com/OneeyK/KPI_GoLang/pkg/server"
	"github.com/OneeyK/KPI_GoLang/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("config/")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	cfg := db.Config{
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
	}

	db.ConnectPostgres(cfg)
	repos := repository.NewRepository(db.DB)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	fmt.Printf("Start server on %s port\n", viper.GetString("port"))
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}
