package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/OneeyK/KPI_GoLang/pkg/handler"
	"github.com/OneeyK/KPI_GoLang/pkg/server"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("config/")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}
	srv := new(server.Server)
	fmt.Printf("Start server on %s port\n", viper.GetString("port"))
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}
