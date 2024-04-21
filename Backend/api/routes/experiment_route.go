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


func NewExperimentRouter(env *config.Env, getDb func() *gorm.DB, group *gin.RouterGroup) {
	experimentRepository := repositories.NewExperimentRepository(getDb)
	experimentService := services.NewExperimentService(experimentRepository)
	experimentController := controllers.NewExperimentController(experimentService)

	authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

	experimentGroup := group.Group("/experiments", authMiddleware)
	{
		experimentGroup.GET("/chapter/:chapterID", experimentController.GetExperimentsPerChapter)
		experimentGroup.GET("/subject/:subjectID", experimentController.GetExperimentsPerSubject)
		experimentGroup.GET("/content/:experimentID", experimentController.GetExperimentContent)
		experimentGroup.POST("/generate", experimentController.GenerateExperimentsPerChapter)
		experimentGroup.POST("", experimentController.AddExperiment)
		experimentGroup.PUT("/:experimentID", experimentController.UpdateExperiment)
		experimentGroup.DELETE("/:experimentID", experimentController.DeleteExperiment)
	}
}

