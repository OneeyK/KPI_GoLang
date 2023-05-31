package quiz

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/service/auth"
	"github.com/OneeyK/KPI_GoLang/pkg/service/question"
	"github.com/OneeyK/KPI_GoLang/pkg/service/quiz"
	"github.com/OneeyK/KPI_GoLang/pkg/service/user"
	"github.com/OneeyK/KPI_GoLang/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AnswerRequest struct {
	Answers []models.Answer `json:"answers"`
}

type QuizHandler struct {
	quiz     quiz.Quiz
	auth     auth.Auth
	user     user.User
	question question.Question
}

func NewQuizHandler(quiz quiz.Quiz, auth auth.Auth, user user.User, question question.Question) *QuizHandler {
	return &QuizHandler{
		quiz:     quiz,
		auth:     auth,
		user:     user,
		question: question,
	}
}

func (h *QuizHandler) CreateQuiz(c *gin.Context) {
	var input models.Quiz
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.quiz.Create(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *QuizHandler) GetQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}

	quiz, err := h.quiz.GetByID(quizID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to get quiz")
		return
	}

	if quiz == nil {
		utils.NewErrorResponse(c, http.StatusNotFound, "quiz not found")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":        quiz.ID,
		"name":      quiz.Name,
		"questions": quiz.Questions,
	})
}

func (h *QuizHandler) UpdateQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}
	var input models.Quiz
	err = c.BindJSON(&input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedQuiz, err := h.quiz.Update(quizID, input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to update quiz")
		return
	}

	c.JSON(http.StatusOK, updatedQuiz)
}

func (h *QuizHandler) DeleteQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}

	err = h.quiz.DeleteByID(quizID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to delete quiz")
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted quiz")
}

func (h *QuizHandler) GetAllQuizzes(c *gin.Context) {
	quizzes, err := h.quiz.GetAll()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to get quizzes")
		return
	}
	c.JSON(http.StatusOK, quizzes)
}

func (h *QuizHandler) TakeQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}
	token := strings.SplitN(c.GetHeader("Authorization"), " ", 2)
	userID, err := h.auth.ParseToken(token[1])
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "invalid token")
	}
	if _, err := h.quiz.GetByID(quizID); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to get quiz")
		return
	}
	if _, err := h.user.GetByID(uint64(userID)); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to get user")
		return
	}
	var answers AnswerRequest
	if err := c.BindJSON(&answers); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid answer format")
		return
	}
	var score int
	for _, answer := range answers.Answers {
		answerQuizID, err := h.question.GetByID(uint64(answer.QuestionID))
		if err != nil {
			utils.NewErrorResponse(c, http.StatusInternalServerError, "wrong quiz id in answer")
		}
		if quizID != uint64(answerQuizID.QuizID) {
			utils.NewErrorResponse(c, http.StatusInternalServerError, "quiz id in url doesn't match quiz id in answer")
		}

		correctAnswer, err := h.question.GetCorrectAnswer(uint64(answer.QuestionID))
		if err != nil {
			utils.NewErrorResponse(c, http.StatusInternalServerError, "can't get correct answer")
			return
		}
		if answer.OptionID == correctAnswer.ID {
			score++
		}
	}
	dasboard, err := h.quiz.SaveDashboard(uint64(userID), uint64(quizID), score)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"score":        dasboard.Score,
		"dashboard_id": dasboard.QuizID,
	})
}

func (h *QuizHandler) GetLeaderBoard(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}

	_, err = h.quiz.GetByID(quizID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to get quiz")
		return
	}
	leaderboard, err := h.quiz.GetLeaderboard(quizID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "no one taken the quiz yet")
		return
	}
	c.JSON(http.StatusOK, leaderboard)
}
