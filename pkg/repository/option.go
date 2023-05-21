package repository

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"gorm.io/gorm"
)

type OptionDB struct {
	db *gorm.DB
}

func NewOptionDB(db *gorm.DB) *OptionDB {
	return &OptionDB{db: db}
}

func (r *OptionDB) Create(questionID uint64, option models.Option) (int, error) {
	option.QuestionID = int(questionID)
	if err := r.db.Create(&option).Error; err != nil {
		return 0, err
	}
	return option.ID, nil
}

func (r *OptionDB) FindByID(optionID uint64) (*models.Option, error) {
	var option models.Option
	if err := r.db.First(&option, optionID).Error; err != nil {
		return nil, err
	}
	return &option, nil
}

func (r *OptionDB) FindByIDForQuestion(questionID uint64, optionID uint64) (*models.Option, error) {
	option := models.Option{QuestionID: int(questionID), ID: int(optionID)}
	if err := r.db.First(&option).Error; err != nil {
		return nil, err
	}
	return &option, nil
}

func (r *OptionDB) Remove(optionID uint64) error {
	var option models.Option
	if err := r.db.Delete(&option, optionID).Error; err != nil {
		return err
	}
	return nil
}

func (r *OptionDB) Update(optionID uint64, updatedOption models.Option) (*models.Option, error) {
	if _, err := r.FindByID(optionID); err != nil {
		return nil, err
	}
	updatedOption.ID = int(optionID)
	if err := r.db.Save(&updatedOption).Error; err != nil {
		return nil, err
	}
	return &updatedOption, nil
}

func (r *OptionDB) GetAll(questionID uint64) ([]models.Option, error) {
	var options []models.Option
	if err := r.db.Find(&options, "question_id = ?", questionID).Error; err != nil {
		return nil, err
	}
	return options, nil
}
