package service

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/repository"
)

type OptionService struct {
	repo repository.Option
}

func NewOptionService(repo repository.Option) *OptionService {
	return &OptionService{repo: repo}
}

func (s *OptionService) Create(questionID uint64, option models.Option) (int, error) {
	return s.repo.Create(questionID, option)
}

func (s *OptionService) GetByID(id uint64) (*models.Option, error) {
	return s.repo.FindByID(id)
}

func (s *OptionService) GetByIDForQuestion(questionID uint64, optionID uint64) (*models.Option, error) {
	return s.repo.FindByIDForQuestion(questionID, optionID)
}

func (s *OptionService) DeleteByID(id uint64) error {
	return s.repo.Remove(id)
}

func (s *OptionService) Update(optionID uint64, updatedOption models.Option) (*models.Option, error) {
	return s.repo.Update(optionID, updatedOption)
}

func (s *OptionService) GetAll(questionID uint64) ([]models.Option, error) {
	return s.repo.GetAll(questionID)
}
