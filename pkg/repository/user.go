package repository

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (r *UserDB) FindByID(userID uint64) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserDB) Remove(userID uint64) error {
	var user models.User
	if err := r.db.Delete(&user, userID).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserDB) Update(userID uint64, updatedUser models.User) (*models.User, error) {
	if _, err := r.FindByID(userID); err != nil {
		return nil, err
	}
	updatedUser.ID = int(userID)
	if err := r.db.Save(&updatedUser).Error; err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (r *UserDB) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
