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

func NewNoteRouter(env *config.Env, getDb func() *gorm.DB, group *gin.RouterGroup) {
	noteRepository := repositories.NewNoteRepository(getDb)
	chapterRepository := repositories.NewChapterRepository(getDb)
	noteService := services.NewNoteService(noteRepository, chapterRepository)
	noteController := controllers.NewNoteController(noteService)

	authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

	noteGroup := group.Group("", authMiddleware)
	{
		noteGroup.POST("/chapters/:chapterID/notes", noteController.AddNoteWithImage)
		noteGroup.GET("/chapters/:chapterID/notes", noteController.GetNoteByChapterID)

		// Add other routes as needed
	}
}
