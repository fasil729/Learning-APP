// routes/quiz_router.go

package routes

import (
	"Brilliant/api/controllers"
	"Brilliant/api/middlewares"
	"Brilliant/application/services"
	"Brilliant/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewQuizRouter sets up the routes for the quiz controller.
func NewQuizRouter(env *config.Env, getDb func() *gorm.DB, group *gin.RouterGroup) {
	quizService := services.NewQuizService()
	quizController := controllers.NewQuizController(quizService)

	authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

	quizGroup := group.Group("/quiz", authMiddleware)
	{
		quizGroup.POST("/generate", quizController.GenerateQuiz)
	}
}