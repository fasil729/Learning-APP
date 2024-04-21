package routes

import (
	"Brilliant/Backend/config"
	"Brilliant/api/middlewares"
	"Brilliant/application/services"
	repositories "Brilliant/persistence/repository"
	"Brilliant/Backend/api/controllers"

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
