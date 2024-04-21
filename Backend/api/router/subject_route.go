package routes

import (
	"hackathon.com/leariningApp/Backend/config"
	"hackathon.com/leariningApp/api/middlewares"
	"hackathon.com/leariningApp/application/services"
	repositories "hackathon.com/leariningApp/persistence/repository"
	"hackathon.com/learningApp/Backend/api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSubjectRouter(env *config.Env, getDb func() *gorm.DB, group *gin.RouterGroup) {
	subjectRepository := repositories.NewSubjectRepository(getDb)
	subjectService := services.NewSubjectService(subjectRepository)
	subjectController := controllers.NewSubjectController(subjectService)

	authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

	subjectGroup := group.Group("/subjects", authMiddleware)
	{
		subjectGroup.POST("", subjectController.CreateSubject)
		// Add more routes as needed
	}
}
