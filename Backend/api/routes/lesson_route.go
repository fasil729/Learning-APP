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

func NewLessonRouter(env *config.Env, getDb func() *gorm.DB, group *gin.RouterGroup) {
    lessonRepository := repositories.NewLessonRepository(getDb)
    lessonService := services.NewLessonService(lessonRepository)
    lessonController := controllers.NewLessonController(lessonService)

    authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

    lessonGroup := group.Group("/lessons", authMiddleware)
    {
        lessonGroup.GET("/content/:lessonID", lessonController.GetLessonContent)
		lessonGroup.POST("/addLesson", lessonController.AddLesson)
		lessonGroup.PUT("/updateLesson/:lessonID", lessonController.UpdateLesson)
		lessonGroup.DELETE("/deleteLesson/:lessonID", lessonController.DeleteLesson)
	
        // Add other routes as needed
    }
}