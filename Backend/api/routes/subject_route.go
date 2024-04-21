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
	subjectService := services.NewSubjectService(subjectRepository)
	subjectController := controllers.NewSubjectController(subjectService)

	authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

	subjectGroup := group.Group("/subjects", authMiddleware)
	{
		subjectGroup.POST("", subjectController.CreateSubject)
		// Add more routes as needed
	}
}
