package question

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"gorm.io/gorm"
)

type QuestionDB struct {
	db *gorm.DB
}

func NewQuestionDB(db *gorm.DB) *QuestionDB {
	return &QuestionDB{db: db}
}

func (r *QuestionDB) Create(quizID uint64, question models.Question) (int, error) {
	question.QuizID = int(quizID)
	if err := r.db.Create(&question).Error; err != nil {
		return 0, err
	}
	return question.ID, nil
}

func (r *QuestionDB) FindByID(questionID uint64) (*models.Question, error) {
	var question models.Question
	if err := r.db.Preload("Options").First(&question, questionID).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionDB) FindByIDForQuiz(quizID uint64, questionID uint64) (*models.Question, error) {
	question := models.Question{QuizID: int(quizID), ID: int(questionID)}
	if err := r.db.Preload("Options").First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionDB) Remove(questionID uint64) error {
	var question models.Question
	if err := r.db.Delete(&question, questionID).Error; err != nil {
		return err
	}
	return nil
}

func (r *QuestionDB) Update(questionID uint64, updatedQuestion models.Question) (*models.Question, error) {
	if _, err := r.FindByID(questionID); err != nil {
		return nil, err
	}
	updatedQuestion.ID = int(questionID)
	if err := r.db.Save(&updatedQuestion).Error; err != nil {
		return nil, err
	}
	return &updatedQuestion, nil
}

func (r *QuestionDB) GetAll(quizID uint64) ([]models.Question, error) {
	var questions []models.Question
	if err := r.db.Preload("Options").Find(&questions, "quiz_id = ?", quizID).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *QuestionDB) GetCorrectAnswer(questionID uint64) (*models.Option, error) {
	var option models.Option
	if err := r.db.Where("correct = ?", "t").First(&option).Error; err != nil {
		return nil, err
	}
	return &option, nil
}
