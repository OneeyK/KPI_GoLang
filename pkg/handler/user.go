package handler

import (
	"net/http"
	"strconv"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserByID(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.services.User.GetByID(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get user")
		return
	}

	if user == nil {
		newErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":   user.ID,
		"name": user.Name,
	})
}

func (h *Handler) deleteUserByID(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	err = h.services.User.DeleteByID(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to delete user")
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted user")
}

func (h *Handler) updateUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	updatedUser, err := h.services.User.Update(userID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to update user")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       updatedUser.ID,
		"name":     updatedUser.Name,
		"password": updatedUser.Password,
	})
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.User.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get users")
		return
	}
	c.JSON(http.StatusOK, users)
}
