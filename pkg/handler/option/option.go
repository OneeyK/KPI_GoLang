package option

import (
	"net/http"
	"strconv"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/service/option"
	"github.com/OneeyK/KPI_GoLang/pkg/service/question"
	"github.com/OneeyK/KPI_GoLang/pkg/utils"
	"github.com/gin-gonic/gin"
)

type OptionHandler struct {
	option   option.Option
	question question.Question
}

func NewOptionHandler(option option.Option, question question.Question) *OptionHandler {
	return &OptionHandler{option: option, question: question}
}

func (h *OptionHandler) AddOptionToQuestion(c *gin.Context) {
	questionIDStr := c.Param("id")
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid option id")
		return
	}
	var input models.Option
	err = c.BindJSON(&input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.question.GetByID(questionID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.option.Create(questionID, input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *OptionHandler) GetOptionForQuestion(c *gin.Context) {
	params := c.Params
	var ids []string

	for _, param := range params {
		if param.Key == "id" {
			ids = append(ids, param.Value)
		}
	}
	questionIDStr := ids[0]
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}
	optionIDStr := ids[1]
	optionID, err := strconv.ParseUint(optionIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid option id")
		return
	}
	option, err := h.option.GetByIDForQuestion(questionID, optionID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, option)
}

func (h *OptionHandler) GetAllOptionsForQuestion(c *gin.Context) {
	optionIDStr := c.Param("id")
	optionID, err := strconv.ParseUint(optionIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}

	options, err := h.option.GetAll(optionID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(options) == 0 {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to get quiz")
		return
	}
	c.JSON(http.StatusOK, options)
}

func (h *OptionHandler) GetOption(c *gin.Context) {
	optionIDStr := c.Param("id")
	optionID, err := strconv.ParseUint(optionIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid option id")
		return
	}

	option, err := h.option.GetByID(optionID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to get option")
		return
	}

	if option == nil {
		utils.NewErrorResponse(c, http.StatusNotFound, "option not found")
		return
	}

	c.JSON(http.StatusOK, option)
}

func (h *OptionHandler) UpdateOption(c *gin.Context) {
	optionIDStr := c.Param("id")
	optionID, err := strconv.ParseUint(optionIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid option id")
		return
	}
	var input models.Option
	if err = c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedOption, err := h.option.Update(optionID, input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to update option")
		return
	}

	c.JSON(http.StatusOK, updatedOption)
}

func (h *OptionHandler) DeleteOption(c *gin.Context) {
	optionIDStr := c.Param("id")
	optionID, err := strconv.ParseUint(optionIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid option id")
		return
	}

	err = h.option.DeleteByID(optionID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to delete option")
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted option")
}
