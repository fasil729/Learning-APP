package routes

import (
	gemini "Brilliant/Infrastructure/Gemini"
	"Brilliant/api/controllers"
	"Brilliant/api/middlewares"
	contracts "Brilliant/application/contracts/gemini"
	"Brilliant/application/services"
	"Brilliant/config"
	repositories "Brilliant/persistence/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
var subjectGeminiHandler contracts.IGeminiSubjectHandler

func NewSubjectRouter(env *config.Env, getDb func() *gorm.DB, group *gin.RouterGroup) {
	subjectRepository := repositories.NewSubjectRepository(getDb)
	chapterRepository := repositories.NewChapterRepository(getDb)
	lessonRepository := repositories.NewLessonRepository(getDb)
	subjectGeminiHandler = gemini.NewGeminiSubjectHandler()
	subjectService := services.NewSubjectService(subjectRepository, chapterRepository, lessonRepository, subjectGeminiHandler)
	subjectController := controllers.NewSubjectController(subjectService)

	authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

	subjectGroup := group.Group("/subjects", authMiddleware)
	{
		subjectGroup.POST("/create", subjectController.CreateSubject)
		subjectGroup.GET("/search", subjectController.SearchSubjectsByName)

		// Add other routes as needed
	}
}
