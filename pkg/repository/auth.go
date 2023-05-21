package repository

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"gorm.io/gorm"
)

type AuthDB struct {
	db *gorm.DB
}

func NewAuthDB(db *gorm.DB) *AuthDB {
	return &AuthDB{db: db}
}

func (r *AuthDB) CreateUser(user models.User) (int, error) {
	r.db.Create(&user)
	return user.ID, nil
}

func (r *AuthDB) GetUser(name, password string) (models.User, error) {
	var user models.User
	err := r.db.Where(&models.User{Name: name, Password: password}).First(&user)
	return user, err.Error
}
