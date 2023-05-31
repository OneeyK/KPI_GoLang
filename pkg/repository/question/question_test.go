package question

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestQuestionDB_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating SQL mock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM DB: %v", err)
	}

	questionDB := NewQuestionDB(gormDB)

	quizID := uint64(1)
	question := models.Question{
		Description: "Sample question",
		Options:     []models.Option{},
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO questions").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	createdQuestionID, err := questionDB.Create(quizID, question)
	if err != nil {
		t.Fatalf("Error creating question: %v", err)
	}

	expectedQuestionID := 1
	if createdQuestionID != expectedQuestionID {
		t.Errorf("Expected created question ID to be %d, but got %d", expectedQuestionID, createdQuestionID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled SQL mock expectations: %v", err)
	}
}

func TestQuestionDB_FindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating SQL mock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM DB: %v", err)
	}

	questionDB := NewQuestionDB(gormDB)

	questionID := uint64(1)

	expectedQuestion := &models.Question{
		ID:          1,
		Description: "Sample question",
		Options:     []models.Option{},
	}

	mock.ExpectQuery("SELECT").WithArgs(questionID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "description"}).
			AddRow(expectedQuestion.ID, expectedQuestion.Description))

	question, err := questionDB.FindByID(questionID)
	if err != nil {
		t.Fatalf("Error finding question by ID: %v", err)
	}

	if question.ID != expectedQuestion.ID {
		t.Errorf("Expected question ID to be %d, but got %d", expectedQuestion.ID, question.ID)
	}
	if question.Description != expectedQuestion.Description {
		t.Errorf("Expected question description to be %s, but got %s", expectedQuestion.Description, question.Description)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled SQL mock expectations: %v", err)
	}
}

func TestQuestionDB_FindByIDForQuiz(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating SQL mock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM DB: %v", err)
	}

	questionDB := NewQuestionDB(gormDB)

	quizID := uint64(1)
	questionID := uint64(1)

	expectedQuestion := &models.Question{
		ID:     1,
		QuizID: 1,
	}

	mock.ExpectQuery("SELECT").WithArgs(quizID, questionID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "quiz_id"}).
			AddRow(expectedQuestion.ID, expectedQuestion.QuizID))

	question, err := questionDB.FindByIDForQuiz(quizID, questionID)
	if err != nil {
		t.Fatalf("Error finding question by ID for quiz: %v", err)
	}

	if question.ID != expectedQuestion.ID {
		t.Errorf("Expected question ID to be %d, but got %d", expectedQuestion.ID, question.ID)
	}
	if question.QuizID != expectedQuestion.QuizID {
		t.Errorf("Expected question quiz ID to be %d, but got %d", expectedQuestion.QuizID, question.QuizID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled SQL mock expectations: %v", err)
	}
}

func TestQuestionDB_Remove(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating SQL mock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM DB: %v", err)
	}

	questionDB := NewQuestionDB(gormDB)

	questionID := uint64(1)

	mock.ExpectExec("DELETE").WithArgs(questionID).WillReturnResult(sqlmock.NewResult(0, 1))

	err = questionDB.Remove(questionID)
	if err != nil {
		t.Fatalf("Error removing question: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled SQL mock expectations: %v", err)
	}
}
