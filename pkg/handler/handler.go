package handler

import (
	"github.com/OneeyK/KPI_GoLang/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
	}

	api := router.Group("/api", h.userIndentity)
	{
		users := api.Group("/users")
		{
			users.GET("/:id", h.getUserByID)
			users.DELETE("/:id", h.deleteUserByID)
			users.PUT("/:id", h.updateUser)
			users.GET("/", h.getAllUsers)
		}
		quizzes := api.Group("/quizzes")
		{
			quizzes.POST("/", h.createQuiz)
			quizzes.GET("/:id", h.getQuiz)
			quizzes.PUT("/:id", h.updateQuiz)
			quizzes.DELETE("/:id", h.deleteQuiz)
			quizzes.GET("/", h.getAllQuizzes)
			questions := quizzes.Group(":id/questions")
			{
				questions.POST("/", h.addQuestionToQuiz)
				questions.GET("/:id", h.getQuestionForQuiz)
				questions.GET("/", h.getQuestionsForQuiz)
			}
			quizzes.GET("/:id/leaderboard", h.getLeaderBoard)
			quizzes.POST("/:id/take", h.takeQuiz)
		}
		questions := api.Group("/questions")
		{
			questions.GET("/:id", h.getQuestion)
			questions.PUT("/:id", h.updateQuestion)
			questions.DELETE("/:id", h.deleteQuestion)
			options := questions.Group(":id/options")
			{
				options.POST("/", h.addOptionToQuestion)
				options.GET("/:id", h.getOptionForQuestion)
				options.GET("/", h.getAllOptionsForQuestion)
			}
		}
		options := api.Group("/options")
		{
			options.GET("/:id", h.getOption)
			options.PUT("/:id", h.updateOption)
			options.DELETE("/:id", h.deleteOption)
		}
	}
	return router
}
