package handler

import (
	"net/http"
	"strconv"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addOptionToQuestion(c *gin.Context) {
	optionIDStr := c.Param("id")
	optionID, err := strconv.ParseUint(optionIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid option id")
		return
	}
	var input models.Option
	err = c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.services.Question.GetByID(optionID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Option.Create(optionID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getOptionForQuestion(c *gin.Context) {
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
		newErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}
	optionIDStr := ids[1]
	optionID, err := strconv.ParseUint(optionIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid option id")
		return
	}
	option, err := h.services.Option.GetByIDForQuestion(questionID, optionID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, option)
}

func (h *Handler) getAllOptionsForQuestion(c *gin.Context) {
	optionIDStr := c.Param("id")
	optionID, err := strconv.ParseUint(optionIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid question id")
		return
	}

	options, err := h.services.Option.GetAll(optionID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(options) == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get quiz")
		return
	}
	c.JSON(http.StatusOK, options)
}

func (h *Handler) getOption(c *gin.Context) {
	optionIDStr := c.Param("id")
	optionID, err := strconv.ParseUint(optionIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid option id")
		return
	}

	option, err := h.services.Option.GetByID(optionID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get option")
		return
	}

	if option == nil {
		newErrorResponse(c, http.StatusNotFound, "option not found")
		return
	}

	c.JSON(http.StatusOK, option)
}

func (h *Handler) updateOption(c *gin.Context) {
	optionIDStr := c.Param("id")
	optionID, err := strconv.ParseUint(optionIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid option id")
		return
	}
	var input models.Option
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedOption, err := h.services.Option.Update(optionID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to update option")
		return
	}

	c.JSON(http.StatusOK, updatedOption)
}

func (h *Handler) deleteOption(c *gin.Context) {
	optionIDStr := c.Param("id")
	optionID, err := strconv.ParseUint(optionIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid option id")
		return
	}

	err = h.services.Option.DeleteByID(optionID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to delete option")
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted option")
}
