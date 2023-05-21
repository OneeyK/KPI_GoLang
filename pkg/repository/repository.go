package repository

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(name, password string) (models.User, error)
}

type User interface {
	FindByID(userID uint64) (*models.User, error)
	Remove(userID uint64) error
	Update(userID uint64, updatedUser models.User) (*models.User, error)
	GetAll() ([]models.User, error)
}

type Quiz interface {
	Create(quiz models.Quiz) (int, error)
	FindByID(quizID uint64) (*models.Quiz, error)
	Remove(quizID uint64) error
	Update(quizID uint64, updatedQuiz models.Quiz) (*models.Quiz, error)
	GetAll() ([]models.Quiz, error)
	SaveDashboard(userID uint64, quizID uint64, score int) (*models.Dashboard, error)
	GetLeaderboard(quizID uint64) (*models.Dashboard, error)
}

type Question interface {
	Create(quizID uint64, question models.Question) (int, error)
	FindByID(questionID uint64) (*models.Question, error)
	FindByIDForQuiz(quizID uint64, questionID uint64) (*models.Question, error)
	Remove(questionID uint64) error
	Update(questionID uint64, updatedQuestion models.Question) (*models.Question, error)
	GetAll(quizID uint64) ([]models.Question, error)
	GetCorrectAnswer(questionID uint64) (*models.Option, error)
}

type Option interface {
	Create(questionID uint64, option models.Option) (int, error)
	FindByID(optionID uint64) (*models.Option, error)
	FindByIDForQuestion(questionID uint64, optionID uint64) (*models.Option, error)
	Remove(optionID uint64) error
	Update(optionID uint64, updateOption models.Option) (*models.Option, error)
	GetAll(questionID uint64) ([]models.Option, error)
}

type Repository struct {
	Authorization
	User
	Quiz
	Question
	Option
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		User:          NewUserDB(db),
		Quiz:          NewQuizDB(db),
		Question:      NewQuestionDB(db),
		Option:        NewOptionDB(db),
	}
}
