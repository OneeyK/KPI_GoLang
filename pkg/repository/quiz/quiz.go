package quiz

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"gorm.io/gorm"
)

type Quiz interface {
	Create(quiz models.Quiz) (int, error)
	FindByID(quizID uint64) (*models.Quiz, error)
	Remove(quizID uint64) error
	Update(quizID uint64, updatedQuiz models.Quiz) (*models.Quiz, error)
	GetAll() ([]models.Quiz, error)
	SaveDashboard(userID uint64, quizID uint64, score int) (*models.Dashboard, error)
	GetLeaderboard(quizID uint64) (*models.Dashboard, error)
}

type QuizDB struct {
	db *gorm.DB
}

func NewQuizDB(db *gorm.DB) *QuizDB {
	return &QuizDB{db: db}
}

func (r *QuizDB) Create(quiz models.Quiz) (int, error) {
	err := r.db.Create(&quiz).Error
	if err != nil {
		return 0, err
	}
	return quiz.ID, nil
}

func (r *QuizDB) FindByID(quizID uint64) (*models.Quiz, error) {
	var quiz models.Quiz
	if err := r.db.Preload("Questions").First(&quiz, quizID).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *QuizDB) Remove(quizID uint64) error {
	var quiz models.Quiz
	err := r.db.Delete(&quiz, quizID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *QuizDB) Update(quizID uint64, updatedQuiz models.Quiz) (*models.Quiz, error) {
	if _, err := r.FindByID(quizID); err != nil {
		return nil, err
	}

	updatedQuiz.ID = int(quizID)
	if err := r.db.Save(&updatedQuiz).Error; err != nil {
		return nil, err
	}

	return &updatedQuiz, nil
}

func (r *QuizDB) GetAll() ([]models.Quiz, error) {
	var quizzes []models.Quiz
	if err := r.db.Preload("Questions").Find(&quizzes).Error; err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (r *QuizDB) SaveDashboard(userID uint64, quizID uint64, score int) (*models.Dashboard, error) {
	dashboard := models.Dashboard{UserID: int(userID), QuizID: int(quizID)}

	if err := r.db.First(&dashboard).Error; err != nil {
		dashboard.Score = 0
	} else {
		oldScore := 0
		r.db.Model(&dashboard).Select("score").Scan(&oldScore)
		dashboard.Score = oldScore
	}

	dashboard.Score += score

	if err := r.db.Save(&dashboard).Error; err != nil {
		return nil, err
	}

	return &dashboard, nil
}

func (r *QuizDB) GetLeaderboard(quizID uint64) (*models.Dashboard, error) {
	var dashboard models.Dashboard
	if err := r.db.First(&dashboard, quizID).Error; err != nil {
		return nil, err
	}
	return &dashboard, nil
}
