package quiz

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestQuizDB_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)
	repo := NewQuizDB(gormDB)
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "quizzes" \("name"\) VALUES \(\$1\) RETURNING "id"`).
		WithArgs("Quiz 1").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	quiz := models.Quiz{Name: "Quiz 1"}
	quizID, err := repo.Create(quiz)
	assert.NoError(t, err)
	assert.Equal(t, 1, quizID)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestQuizDB_FindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)
	repo := NewQuizDB(gormDB)

	quizID := uint64(1)
	expectedQuiz := &models.Quiz{ID: 1, Name: "Quiz 1", Questions: []models.Question{}}

	mock.ExpectQuery(`SELECT \* FROM "quizzes" WHERE "quizzes"."id" = \$1 ORDER BY "quizzes"."id" LIMIT 1`).
		WithArgs(quizID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(expectedQuiz.ID, expectedQuiz.Name))

	mock.ExpectQuery(`SELECT \* FROM "questions" WHERE "questions"."quiz_id" = \$1`).
		WithArgs(quizID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "text", "quiz_id"})) // Provide mock result for questions query

	quiz, err := repo.FindByID(quizID)
	assert.NoError(t, err)

	assert.Equal(t, expectedQuiz, quiz)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestQuizDB_Remove(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)
	repo := NewQuizDB(gormDB)

	quizID := uint64(1)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "quizzes"`).
		WithArgs(quizID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = repo.Remove(quizID)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestQuizDB_SaveDashboard(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)
	repo := NewQuizDB(gormDB)

	userID := uint64(1)
	quizID := uint64(1)
	score := 10

	dashboard := models.Dashboard{UserID: 1, QuizID: 1, Score: 10}

	mock.ExpectQuery(`SELECT \* FROM "dashboards" WHERE "dashboards"."user_id" = \? AND "dashboards"."quiz_id" = \? ORDER BY "dashboards"."id" LIMIT 1`).
		WithArgs(userID, quizID).
		WillReturnError(gorm.ErrRecordNotFound)

	mock.ExpectQuery(`SELECT score FROM "dashboards" WHERE "dashboards"."user_id" = \? AND "dashboards"."quiz_id" = \? ORDER BY "dashboards"."id" LIMIT 1`).
		WithArgs(userID, quizID).
		WillReturnRows(sqlmock.NewRows([]string{"score"}).AddRow(10))

	mock.ExpectExec(`UPDATE "dashboards" SET "score" = \? WHERE "dashboards"."user_id" = \? AND "dashboards"."quiz_id" = \?`).
		WithArgs(20, userID, quizID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	result, err := repo.SaveDashboard(userID, quizID, score)

	assert.NoError(t, err)
	assert.Equal(t, &dashboard, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestQuizDB_GetLeaderboard(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)
	repo := NewQuizDB(gormDB)

	quizID := uint64(1)
	expectedDashboard := models.Dashboard{UserID: 1, QuizID: 1, Score: 10}

	mock.ExpectQuery(`SELECT \* FROM "dashboards" WHERE "dashboards"."quiz_id" = \? ORDER BY "dashboards"."id" LIMIT 1`).
		WithArgs(quizID).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "quiz_id", "score"}).AddRow(1, 1, 10))

	result, err := repo.GetLeaderboard(quizID)

	assert.NoError(t, err)
	assert.Equal(t, &expectedDashboard, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestQuizDB_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)
	repo := NewQuizDB(gormDB)

	expectedQuizzes := []models.Quiz{
		{ID: 1, Name: "Quiz 1", Questions: []models.Question{{ID: 1, QuizID: 1, Description: "Question 1"}, {ID: 2, QuizID: 1, Description: "Question 2"}}},
		{ID: 2, Name: "Quiz 2", Questions: []models.Question{{ID: 3, QuizID: 2, Description: "Question 3"}, {ID: 4, QuizID: 2, Description: "Question 4"}}},
	}

	mock.ExpectQuery(`SELECT \* FROM "quizzes"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Quiz 1").
			AddRow(2, "Quiz 2"))

	mock.ExpectQuery(`SELECT \* FROM "questions" WHERE "questions"."quiz_id" IN \(1, 2\)`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "quiz_id", "description"}).
			AddRow(1, 1, "Question 1").
			AddRow(2, 1, "Question 2").
			AddRow(3, 2, "Question 3").
			AddRow(4, 2, "Question 4"))

	// Call the GetAll method and get the result
	result, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedQuizzes, result)

	// Verify that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
