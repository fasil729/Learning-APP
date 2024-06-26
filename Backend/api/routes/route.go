package routes

import (
	"Brilliant/config"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Setup(env *config.Env, getDb func() *gorm.DB, gin *gin.Engine) {
	// Setup routes
	publicRouter := gin.Group("")

	// Setup user routes
	NewUserRouter(env, getDb, publicRouter)

	// setup subject routes
	NewSubjectRouter(env, getDb, publicRouter)

	// setup lesson routes

	NewLessonRouter(env, getDb, publicRouter)

	// setup chapter routes
	NewChapterRouter(env, getDb, publicRouter)

	// setup note routes
	NewNoteRouter(env, getDb, publicRouter)

	// Setup experiment routes
	NewExperimentRouter(env, getDb, publicRouter)

	// setup quiz routes
	NewQuizRouter(env, getDb, publicRouter)

	// setup exam prep routes
	NewExamPrepRouter(env, getDb, publicRouter)

	// setup other routes below when needed
}
