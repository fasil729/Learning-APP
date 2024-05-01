// routes/examprep_router.go

package routes

import (
	gemini "Brilliant/Infrastructure/Gemini"
	"Brilliant/api/controllers"
	"Brilliant/api/middlewares"
	"Brilliant/application/services"
	"Brilliant/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewExamPrepRouter sets up the routes for the exam preparation controller.
func NewExamPrepRouter(env *config.Env, getDb func() *gorm.DB, group *gin.RouterGroup) {
	examGeminiHandler := gemini.NewGeminiExamPrepHandler()
	examPrepService := services.NewExamPrepService(examGeminiHandler)
	examPrepController := controllers.NewExamPrepController(examPrepService)

	authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

	examPrepGroup := group.Group("/examprep", authMiddleware)
	{
		examPrepGroup.POST("/generate", examPrepController.GenerateExamPrepMaterial)
	}
}
