package service

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(name, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetByID(id uint64) (*models.User, error)
	DeleteByID(id uint64) error
	Update(userID uint64, updatedUser models.User) (*models.User, error)
	GetAll() ([]models.User, error)
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

type Question interface {
	Create(quizID uint64, question models.Question) (int, error)
	GetByID(id uint64) (*models.Question, error)
	GetByIDForQuiz(quizID uint64, questionID uint64) (*models.Question, error)
	DeleteByID(id uint64) error
	Update(questionID uint64, updatedQuestion models.Question) (*models.Question, error)
	GetAll(quizID uint64) ([]models.Question, error)
	GetCorrectAnswer(questionID uint64) (*models.Option, error)
}

type Option interface {
	Create(optionID uint64, option models.Option) (int, error)
	GetByID(id uint64) (*models.Option, error)
	GetByIDForQuestion(questionID uint64, optionID uint64) (*models.Option, error)
	DeleteByID(id uint64) error
	Update(optionID uint64, updatedOption models.Option) (*models.Option, error)
	GetAll(questionID uint64) ([]models.Option, error)
}

type Service struct {
	Authorization
	User
	Quiz
	Question
	Option
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Quiz:          NewQuizService(repos.Quiz),
		Question:      NewQuestionService(repos.Question),
		Option:        NewOptionService(repos.Option),
	}
}
