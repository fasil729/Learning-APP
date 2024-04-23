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

func NewSubjectRouter(env *config.Env, getDb func() *gorm.DB, group *gin.RouterGroup) {
	subjectRepository := repositories.NewSubjectRepository(getDb)
	chapterRepository := repositories.NewChapterRepository(getDb)
	lessonRepository := repositories.NewLessonRepository(getDb)
	subjectService := services.NewSubjectService(subjectRepository, chapterRepository, lessonRepository)
	subjectController := controllers.NewSubjectController(subjectService)

	authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

	subjectGroup := group.Group("/subjects", authMiddleware)
	{
		subjectGroup.POST("/Create", subjectController.CreateSubject)
		subjectGroup.POST("/GenerateSubject", subjectController.GenerateRoadMap)

		// Add other routes as needed
	}
}
