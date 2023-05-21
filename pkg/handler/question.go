package handler

import (
	"net/http"
	"strconv"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addQuestionToQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}
	var input models.Question
	err = c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.services.Quiz.GetByID(quizID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Question.Create(quizID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getQuestionForQuiz(c *gin.Context) {
	params := c.Params
	var ids []string

	for _, param := range params {
		if param.Key == "id" {
			ids = append(ids, param.Value)
		}
	}
	quizIDStr := ids[0]
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}
	questionIDStr := ids[1]
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}
	question, err := h.services.Question.GetByIDForQuiz(quizID, questionID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, question)
}

func (h *Handler) getQuestion(c *gin.Context) {
	questionIDStr := c.Param("id")
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}

	question, err := h.services.Question.GetByID(questionID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get question")
		return
	}

	if question == nil {
		newErrorResponse(c, http.StatusNotFound, "question not found")
		return
	}

	c.JSON(http.StatusOK, question)
}

func (h *Handler) updateQuestion(c *gin.Context) {
	questionIDStr := c.Param("id")
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}
	var input models.Question
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedQuestion, err := h.services.Question.Update(questionID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to update question")
		return
	}

	c.JSON(http.StatusOK, updatedQuestion)
}

func (h *Handler) deleteQuestion(c *gin.Context) {
	questionIDStr := c.Param("id")
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}

	err = h.services.Question.DeleteByID(questionID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to delete question")
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted question")
}

func (h *Handler) getQuestionsForQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}

	questions, err := h.services.Question.GetAll(quizID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(questions) == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get quiz")
		return
	}
	c.JSON(http.StatusOK, questions)
}
