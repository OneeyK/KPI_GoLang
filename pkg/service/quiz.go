package service

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/repository"
)

type QuizService struct {
	repo repository.Quiz
}

func NewQuizService(repo repository.Quiz) *QuizService {
	return &QuizService{repo: repo}
}

func (s *QuizService) Create(quiz models.Quiz) (int, error) {
	return s.repo.Create(quiz)
}

func (s *QuizService) GetByID(id uint64) (*models.Quiz, error) {
	return s.repo.FindByID(id)
}

func (s *QuizService) DeleteByID(id uint64) error {
	return s.repo.Remove(id)
}

func (s *QuizService) Update(quizID uint64, updatedQuiz models.Quiz) (*models.Quiz, error) {
	return s.repo.Update(quizID, updatedQuiz)
}

func (s *QuizService) GetAll() ([]models.Quiz, error) {
	return s.repo.GetAll()
}

func (s *QuizService) SaveDashboard(userID uint64, quizID uint64, score int) (*models.Dashboard, error) {
	return s.repo.SaveDashboard(userID, quizID, score)
}

func (s *QuizService) GetLeaderboard(quizID uint64) (*models.Dashboard, error) {
	return s.repo.GetLeaderboard(quizID)
}
