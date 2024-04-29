package routes

import (
	"Brilliant/api/controllers"
	"Brilliant/api/middlewares"
	"Brilliant/application/services"
	"Brilliant/config"
	repositories "Brilliant/persistence/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewChapterRouter(env *config.Env, getDb func() *gorm.DB, group *gin.RouterGroup) {
	chapterRepository := repositories.NewChapterRepository(getDb)
	chapterService := services.NewChapterService(chapterRepository)
	chapterController := controllers.NewChapterController(chapterService)

	authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

	chapterGroup := group.Group("/chapters", authMiddleware)
	{
		chapterGroup.POST("/chapter", chapterController.CreateChapter)
		chapterGroup.PUT("/:chapterID", chapterController.UpdateChapter)
		chapterGroup.DELETE("/:chapterID", chapterController.DeleteChapter)

		// Add other routes as needed
	}
}
