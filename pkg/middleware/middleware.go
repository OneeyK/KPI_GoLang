package middleware

import (
	"net/http"
	"strings"

	"github.com/OneeyK/KPI_GoLang/pkg/service/auth"
	"github.com/OneeyK/KPI_GoLang/pkg/utils"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

type UseAuthMiddleware struct {
	authService auth.Auth
}

func NewUseAuthMiddleware(authService auth.Auth) *UseAuthMiddleware {
	return &UseAuthMiddleware{authService: authService}
}

func (m *UseAuthMiddleware) UserIndentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userID, err := m.authService.ParseToken(headerParts[1])
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("id", userID)
}
