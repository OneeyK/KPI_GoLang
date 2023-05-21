package service

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/repository"
)

type QuestionService struct {
	repo repository.Question
}

func NewQuestionService(repo repository.Question) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) Create(quizID uint64, question models.Question) (int, error) {
	return s.repo.Create(quizID, question)
}

func (s *QuestionService) GetByID(id uint64) (*models.Question, error) {
	return s.repo.FindByID(id)
}

func (s *QuestionService) GetByIDForQuiz(quizID uint64, questionID uint64) (*models.Question, error) {
	return s.repo.FindByIDForQuiz(quizID, questionID)
}

func (s *QuestionService) DeleteByID(id uint64) error {
	return s.repo.Remove(id)
}

func (s *QuestionService) Update(questionID uint64, updatedQuestion models.Question) (*models.Question, error) {
	return s.repo.Update(questionID, updatedQuestion)
}

func (s *QuestionService) GetAll(quizID uint64) ([]models.Question, error) {
	return s.repo.GetAll(quizID)
}

func (s *QuestionService) GetCorrectAnswer(questionID uint64) (*models.Option, error) {
	return s.repo.GetCorrectAnswer(questionID)
}
