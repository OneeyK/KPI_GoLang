package option

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
)

type optionRepository interface {
	Create(questionID uint64, option models.Option) (int, error)
	FindByID(optionID uint64) (*models.Option, error)
	FindByIDForQuestion(questionID uint64, optionID uint64) (*models.Option, error)
	Remove(optionID uint64) error
	Update(optionID uint64, updateOption models.Option) (*models.Option, error)
	GetAll(questionID uint64) ([]models.Option, error)
}

type OptionService struct {
	optionRepository optionRepository
}

func NewOptionService(optionRepository optionRepository) Option {
	return &OptionService{optionRepository: optionRepository}
}

type Option interface {
	Create(questionID uint64, option models.Option) (int, error)
	GetByID(id uint64) (*models.Option, error)
	GetByIDForQuestion(questionID uint64, optionID uint64) (*models.Option, error)
	DeleteByID(id uint64) error
	Update(optionID uint64, updatedOption models.Option) (*models.Option, error)
	GetAll(questionID uint64) ([]models.Option, error)
}

func (s *OptionService) Create(questionID uint64, option models.Option) (int, error) {
	return s.optionRepository.Create(questionID, option)
}

func (s *OptionService) GetByID(id uint64) (*models.Option, error) {
	return s.optionRepository.FindByID(id)
}

func (s *OptionService) GetByIDForQuestion(questionID uint64, optionID uint64) (*models.Option, error) {
	return s.optionRepository.FindByIDForQuestion(questionID, optionID)
}

func (s *OptionService) DeleteByID(id uint64) error {
	return s.optionRepository.Remove(id)
}

func (s *OptionService) Update(optionID uint64, updatedOption models.Option) (*models.Option, error) {
	return s.optionRepository.Update(optionID, updatedOption)
}

func (s *OptionService) GetAll(questionID uint64) ([]models.Option, error) {
	return s.optionRepository.GetAll(questionID)
}
