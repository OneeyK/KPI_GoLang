package question

import (
	"net/http"
	"strconv"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/service/question"
	"github.com/OneeyK/KPI_GoLang/pkg/service/quiz"
	"github.com/OneeyK/KPI_GoLang/pkg/utils"
	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	question question.Question
	quiz     quiz.Quiz
}

func NewQuestionHandler(question question.Question, quiz quiz.Quiz) *QuestionHandler {
	return &QuestionHandler{question: question, quiz: quiz}
}

func (h *QuestionHandler) AddQuestionToQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}
	var input models.Question
	err = c.BindJSON(&input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.quiz.GetByID(quizID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.question.Create(quizID, input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *QuestionHandler) GetQuestionForQuiz(c *gin.Context) {
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
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}
	questionIDStr := ids[1]
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}
	question, err := h.question.GetByIDForQuiz(quizID, questionID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) GetQuestion(c *gin.Context) {
	questionIDStr := c.Param("id")
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}

	question, err := h.question.GetByID(questionID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to get question")
		return
	}

	if question == nil {
		utils.NewErrorResponse(c, http.StatusNotFound, "question not found")
		return
	}

	c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) UpdateQuestion(c *gin.Context) {
	questionIDStr := c.Param("id")
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}
	var input models.Question
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedQuestion, err := h.question.Update(questionID, input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to update question")
		return
	}

	c.JSON(http.StatusOK, updatedQuestion)
}

func (h *QuestionHandler) DeleteQuestion(c *gin.Context) {
	questionIDStr := c.Param("id")
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}

	err = h.question.DeleteByID(questionID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to delete question")
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted question")
}

func (h *QuestionHandler) GetQuestionsForQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := strconv.ParseUint(quizIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid quiz id")
		return
	}

	questions, err := h.question.GetAll(quizID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(questions) == 0 {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to get quiz")
		return
	}
	c.JSON(http.StatusOK, questions)
}
