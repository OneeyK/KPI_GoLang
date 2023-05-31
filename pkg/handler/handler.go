package handler

import (
	"log"

	"github.com/OneeyK/KPI_GoLang/pkg/db"
	authRepository "github.com/OneeyK/KPI_GoLang/pkg/repository/auth"
	optionRepository "github.com/OneeyK/KPI_GoLang/pkg/repository/option"
	questionRepository "github.com/OneeyK/KPI_GoLang/pkg/repository/question"
	quizRepository "github.com/OneeyK/KPI_GoLang/pkg/repository/quiz"
	userRepository "github.com/OneeyK/KPI_GoLang/pkg/repository/user"

	"github.com/OneeyK/KPI_GoLang/pkg/middleware"
	authService "github.com/OneeyK/KPI_GoLang/pkg/service/auth"
	optionService "github.com/OneeyK/KPI_GoLang/pkg/service/option"
	questionService "github.com/OneeyK/KPI_GoLang/pkg/service/question"
	quizService "github.com/OneeyK/KPI_GoLang/pkg/service/quiz"
	userService "github.com/OneeyK/KPI_GoLang/pkg/service/user"

	authHandler "github.com/OneeyK/KPI_GoLang/pkg/handler/auth"
	optionHandler "github.com/OneeyK/KPI_GoLang/pkg/handler/option"
	questionHandler "github.com/OneeyK/KPI_GoLang/pkg/handler/question"
	quizHandler "github.com/OneeyK/KPI_GoLang/pkg/handler/quiz"
	userHandler "github.com/OneeyK/KPI_GoLang/pkg/handler/user"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRoutes() *gin.Engine {
	viper.AddConfigPath("config/")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	cfg := db.Config{
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
	}

	db.ConnectPostgres(cfg)

	authRepository := authRepository.NewAuthDB(db.DB)
	optionRepository := optionRepository.NewOptionDB(db.DB)
	questionRepository := questionRepository.NewQuestionDB(db.DB)
	quizRepository := quizRepository.NewQuizDB(db.DB)
	userRepository := userRepository.NewUserDB(db.DB)

	optionService := optionService.NewOptionService(optionRepository)
	questionService := questionService.NewQuestionService(questionRepository)
	quizService := quizService.NewQuizService(quizRepository)
	userService := userService.NewUserService(userRepository)
	authService := authService.NewAuthService(authRepository, userService)

	useAuthMiddleware := middleware.NewUseAuthMiddleware(authService)
	authHandler := authHandler.NewAuthHandler(authService)
	optionHandler := optionHandler.NewOptionHandler(optionService, questionService)
	questionHandler := questionHandler.NewQuestionHandler(questionService, quizService)
	quizHandler := quizHandler.NewQuizHandler(quizService, authService, userService, questionService)
	userHandler := userHandler.NewUserHandler(userService)

	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/signup", authHandler.SignUp)
		auth.POST("/signin", authHandler.SignIn)
	}

	api := router.Group("/api", useAuthMiddleware.UserIndentity)
	{
		users := api.Group("/users")
		{
			users.GET("/:id", userHandler.GetUserByID)
			users.DELETE("/:id", userHandler.DeleteUserByID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.GET("/", userHandler.GetAllUsers)
		}
		quizzes := api.Group("/quizzes")
		{
			quizzes.POST("/", quizHandler.CreateQuiz)
			quizzes.GET("/:id", quizHandler.GetQuiz)
			quizzes.PUT("/:id", quizHandler.UpdateQuiz)
			quizzes.DELETE("/:id", quizHandler.UpdateQuiz)
			quizzes.GET("/", quizHandler.GetAllQuizzes)
			questions := quizzes.Group(":id/questions")
			{
				questions.POST("/", questionHandler.AddQuestionToQuiz)
				questions.GET("/:id", questionHandler.GetQuestionForQuiz)
				questions.GET("/", questionHandler.GetQuestionsForQuiz)
			}
			quizzes.GET("/:id/leaderboard", quizHandler.GetLeaderBoard)
			quizzes.POST("/:id/take", quizHandler.TakeQuiz)
		}
		questions := api.Group("/questions")
		{
			questions.GET("/:id", questionHandler.GetQuestion)
			questions.PUT("/:id", questionHandler.UpdateQuestion)
			questions.DELETE("/:id", questionHandler.DeleteQuestion)
			options := questions.Group(":id/options")
			{
				options.POST("/", optionHandler.AddOptionToQuestion)
				options.GET("/:id", optionHandler.GetOptionForQuestion)
				options.GET("/", optionHandler.GetAllOptionsForQuestion)
			}
		}
		options := api.Group("/options")
		{
			options.GET("/:id", optionHandler.GetOption)
			options.PUT("/:id", optionHandler.UpdateOption)
			options.DELETE("/:id", optionHandler.DeleteOption)
		}
	}
	return router
}
