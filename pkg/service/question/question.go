package question

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
)

type questionRepository interface {
	Create(quizID uint64, question models.Question) (int, error)
	FindByID(questionID uint64) (*models.Question, error)
	FindByIDForQuiz(quizID uint64, questionID uint64) (*models.Question, error)
	Remove(questionID uint64) error
	Update(questionID uint64, updatedQuestion models.Question) (*models.Question, error)
	GetAll(quizID uint64) ([]models.Question, error)
	GetCorrectAnswer(questionID uint64) (*models.Option, error)
}

type QuestionService struct {
	questionRepository questionRepository
}

func NewQuestionService(questionRepository questionRepository) Question {
	return &QuestionService{questionRepository: questionRepository}
}

type Question interface {
	Create(quizID uint64, question models.Question) (int, error)
	GetByID(id uint64) (*models.Question, error)
	GetByIDForQuiz(quizID uint64, questionID uint64) (*models.Question, error)
	DeleteByID(id uint64) error
	Update(questionID uint64, updatedQuestion models.Question) (*models.Question, error)
	GetAll(quizID uint64) ([]models.Question, error)
	GetCorrectAnswer(questionID uint64) (*models.Option, error)
}

func (s *QuestionService) Create(quizID uint64, question models.Question) (int, error) {
	return s.questionRepository.Create(quizID, question)
}

func (s *QuestionService) GetByID(id uint64) (*models.Question, error) {
	return s.questionRepository.FindByID(id)
}

func (s *QuestionService) GetByIDForQuiz(quizID uint64, questionID uint64) (*models.Question, error) {
	return s.questionRepository.FindByIDForQuiz(quizID, questionID)
}

func (s *QuestionService) DeleteByID(id uint64) error {
	return s.questionRepository.Remove(id)
}

func (s *QuestionService) Update(questionID uint64, updatedQuestion models.Question) (*models.Question, error) {
	return s.questionRepository.Update(questionID, updatedQuestion)
}

func (s *QuestionService) GetAll(quizID uint64) ([]models.Question, error) {
	return s.questionRepository.GetAll(quizID)
}

func (s *QuestionService) GetCorrectAnswer(questionID uint64) (*models.Option, error) {
	return s.questionRepository.GetCorrectAnswer(questionID)
}
