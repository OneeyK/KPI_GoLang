package user

import (
	"net/http"
	"strconv"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/service/user"
	"github.com/OneeyK/KPI_GoLang/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	user user.User
}

func NewUserHandler(user user.User) *UserHandler {
	return &UserHandler{
		user: user,
	}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.user.GetByID(userID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to get user")
		return
	}

	if user == nil {
		utils.NewErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":   user.ID,
		"name": user.Name,
	})
}

func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	err = h.user.DeleteByID(userID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to delete user")
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted user")
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	updatedUser, err := h.user.Update(userID, input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to update user")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       updatedUser.ID,
		"name":     updatedUser.Name,
		"password": updatedUser.Password,
	})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.user.GetAll()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "failed to get users")
		return
	}
	c.JSON(http.StatusOK, users)
}
