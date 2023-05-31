package option

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/service/option"
	"github.com/OneeyK/KPI_GoLang/pkg/service/question"
)

func TestOptionHandler_AddOptionToQuestion(t *testing.T) {
	optionHandler := NewOptionHandler(&option.OptionService{}, &question.QuestionService{})

	router := gin.Default()
	router.POST("/questions/:id/options", optionHandler.AddOptionToQuestion)

	questionID := uint64(1)

	input := models.Option{
		Content: "Sample Option",
	}

	jsonInput, _ := json.Marshal(input)

	req, _ := http.NewRequest("POST", "/questions/"+strconv.FormatUint(questionID, 10)+"/options", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var responseBody map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &responseBody)

	assert.Contains(t, responseBody, "id")
}

func TestOptionHandler_GetOptionForQuestion(t *testing.T) {
	optionHandler := NewOptionHandler(&option.OptionService{}, &question.QuestionService{})

	router := gin.Default()
	router.GET("/questions/:id/options/:optionID", optionHandler.GetOptionForQuestion)

	questionID := uint64(1)
	optionID := uint64(1)

	req, _ := http.NewRequest("GET", "/questions/"+strconv.FormatUint(questionID, 10)+"/options/"+strconv.FormatUint(optionID, 10), nil)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var responseBody models.Option
	json.Unmarshal(rec.Body.Bytes(), &responseBody)
}
