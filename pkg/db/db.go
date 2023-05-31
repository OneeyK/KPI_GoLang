package db

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Username string
	Password string
	Port     string
	DBName   string
}

var DB *gorm.DB

func ConnectPostgres(cfg Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error during connection to database: %s", err.Error())
	}

	err = DB.AutoMigrate(
		&models.Quiz{},
		&models.Dashboard{},
		&models.User{},
		&models.Option{},
		&models.Question{},
	)
	if err != nil {
		log.Fatalf("Error during migrations: %s", err.Error())
	}
}
