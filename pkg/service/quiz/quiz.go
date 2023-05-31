package quiz

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
)

type quizRepository interface {
	Create(quiz models.Quiz) (int, error)
	FindByID(quizID uint64) (*models.Quiz, error)
	Remove(quizID uint64) error
	Update(quizID uint64, updatedQuiz models.Quiz) (*models.Quiz, error)
	GetAll() ([]models.Quiz, error)
	SaveDashboard(userID uint64, quizID uint64, score int) (*models.Dashboard, error)
	GetLeaderboard(quizID uint64) (*models.Dashboard, error)
}

type QuizService struct {
	quizRepository quizRepository
}

func NewQuizService(quizRepository quizRepository) Quiz {
	return &QuizService{quizRepository: quizRepository}
}

type Quiz interface {
	Create(quiz models.Quiz) (int, error)
	GetByID(id uint64) (*models.Quiz, error)
	DeleteByID(id uint64) error
	Update(quizID uint64, updatedQuiz models.Quiz) (*models.Quiz, error)
	GetAll() ([]models.Quiz, error)
	SaveDashboard(userID uint64, quizID uint64, score int) (*models.Dashboard, error)
	GetLeaderboard(quizID uint64) (*models.Dashboard, error)
}

func (s *QuizService) Create(quiz models.Quiz) (int, error) {
	return s.quizRepository.Create(quiz)
}

func (s *QuizService) GetByID(id uint64) (*models.Quiz, error) {
	return s.quizRepository.FindByID(id)
}

func (s *QuizService) DeleteByID(id uint64) error {
	return s.quizRepository.Remove(id)
}

func (s *QuizService) Update(quizID uint64, updatedQuiz models.Quiz) (*models.Quiz, error) {
	return s.quizRepository.Update(quizID, updatedQuiz)
}

func (s *QuizService) GetAll() ([]models.Quiz, error) {
	return s.quizRepository.GetAll()
}

func (s *QuizService) SaveDashboard(userID uint64, quizID uint64, score int) (*models.Dashboard, error) {
	return s.quizRepository.SaveDashboard(userID, quizID, score)
}

func (s *QuizService) GetLeaderboard(quizID uint64) (*models.Dashboard, error) {
	return s.quizRepository.GetLeaderboard(quizID)
}
