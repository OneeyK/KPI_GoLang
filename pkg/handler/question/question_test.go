package question

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/service/question"
	"github.com/OneeyK/KPI_GoLang/pkg/service/quiz"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddQuestionToQuiz(t *testing.T) {
	questionService := &question.QuestionService{} // Replace with your question service implementation
	quizService := &quiz.QuizService{}             // Replace with your quiz service implementation
	questionHandler := NewQuestionHandler(questionService, quizService)

	router := gin.Default()
	router.POST("/quizzes/:id/questions", questionHandler.AddQuestionToQuiz)

	quizID := uint64(1) // Replace with the desired quiz ID
	input := models.Question{
		Description: "Sample question description",
		Options: []models.Option{
			{
				Content: "Option A",
				Correct: true,
			},
			{
				Content: "Option B",
				Correct: false,
			},
			{
				Content: "Option C",
				Correct: false,
			},
			{
				Content: "Option D",
				Correct: false,
			},
		},
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/quizzes/"+strconv.FormatUint(quizID, 10)+"/questions", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Body = ioutil.NopCloser(strings.NewReader(inputJSON))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetQuestionForQuiz(t *testing.T) {
	questionService := &question.QuestionService{}
	quizService := &quiz.QuizService{}
	questionHandler := NewQuestionHandler(questionService, quizService)

	router := gin.Default()
	router.GET("/quizzes/:id/questions/:qid", questionHandler.GetQuestionForQuiz)

	quizID := uint64(1)     // Replace with the desired quiz ID
	questionID := uint64(1) // Replace with the desired question ID
	input := models.Question{
		Description: "Sample question description",
		Options: []models.Option{
			{
				Content: "Option A",
				Correct: true,
			},
			{
				Content: "Option B",
				Correct: false,
			},
			{
				Content: "Option C",
				Correct: false,
			},
			{
				Content: "Option D",
				Correct: false,
			},
		},
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/quizzes/"+strconv.FormatUint(quizID, 10)+"/questions/"+strconv.FormatUint(questionID, 10), nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
