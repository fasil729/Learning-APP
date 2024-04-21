// routes/examprep_router.go

package routes

import (
	"Brilliant/api/controllers"
	"Brilliant/api/middlewares"
	"Brilliant/application/services"
	"Brilliant/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewExamPrepRouter sets up the routes for the exam preparation controller.
func NewExamPrepRouter(env *config.Env, getDb func() *gorm.DB, group *gin.RouterGroup) {
	examPrepService := services.NewExamPrepService()
	examPrepController := controllers.NewExamPrepController(examPrepService)

	authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

	examPrepGroup := group.Group("/examprep", authMiddleware)
	{
		examPrepGroup.POST("/generate", examPrepController.GenerateExamPrepMaterial)
	}
}