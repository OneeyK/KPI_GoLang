package auth

import (
	"net/http"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/service/auth"
	"github.com/OneeyK/KPI_GoLang/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth auth.Auth
}

func NewAuthHandler(auth auth.Auth) *AuthHandler {
	return &AuthHandler{auth: auth}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.auth.CreateUser(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Name     string `json:"name,omitempty"     binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.auth.GenerateToken(input.Name, input.Password)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
