package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/gin-gonic/gin"
)

type AnswerRequest struct {
	Answers []models.Answer `json:"answers"`
}

func (h *Handler) createQuiz(c *gin.Context) {
	var input models.Quiz
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Quiz.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}

	quiz, err := h.services.Quiz.GetByID(quizID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get quiz")
		return
	}

	if quiz == nil {
		newErrorResponse(c, http.StatusNotFound, "quiz not found")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":        quiz.ID,
		"name":      quiz.Name,
		"questions": quiz.Questions,
	})
}

func (h *Handler) updateQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}
	var input models.Quiz
	err = c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedQuiz, err := h.services.Quiz.Update(quizID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to update quiz")
		return
	}

	c.JSON(http.StatusOK, updatedQuiz)
}

func (h *Handler) deleteQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}

	err = h.services.Quiz.DeleteByID(quizID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to delete quiz")
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted quiz")
}

func (h *Handler) getAllQuizzes(c *gin.Context) {
	quizzes, err := h.services.Quiz.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get quizzes")
		return
	}
	c.JSON(http.StatusOK, quizzes)
}

func (h *Handler) takeQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}
	token := strings.SplitN(c.GetHeader("Authorization"), " ", 2)
	userID, err := h.services.Authorization.ParseToken(token[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid token")
	}
	if _, err := h.services.Quiz.GetByID(quizID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get quiz")
		return
	}
	if _, err := h.services.User.GetByID(uint64(userID)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get user")
		return
	}
	var answers AnswerRequest
	if err := c.BindJSON(&answers); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid answer format")
		return
	}
	var score int
	for _, answer := range answers.Answers {
		answerQuizID, err := h.services.Question.GetByID(uint64(answer.QuestionID))
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, "wrong quiz id in answer")
		}
		if quizID != uint64(answerQuizID.QuizID) {
			newErrorResponse(c, http.StatusInternalServerError, "quiz id in url doesn't match quiz id in answer")
		}

		correctAnswer, err := h.services.Question.GetCorrectAnswer(uint64(answer.QuestionID))
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, "can't get correct answer")
			return
		}
		if answer.OptionID == correctAnswer.ID {
			score++
		}
	}
	dasboard, err := h.services.Quiz.SaveDashboard(uint64(userID), uint64(quizID), score)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"score":        dasboard.Score,
		"dashboard_id": dasboard.QuizID,
	})
}

func (h *Handler) getLeaderBoard(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}

	_, err = h.services.Quiz.GetByID(quizID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get quiz")
		return
	}
	leaderboard, err := h.services.Quiz.GetLeaderboard(quizID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "no one taken the quiz yet")
		return
	}
	c.JSON(http.StatusOK, leaderboard)
}
