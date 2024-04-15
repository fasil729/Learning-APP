package routes

import (
	"gorm.io/gorm"
    "Brilliant/config"

    "github.com/gin-gonic/gin"
)

func Setup(env *config.Env, getDb func() *gorm.DB, gin *gin.Engine) {
    // Setup routes
    publicRouter := gin.Group("")

	// Setup user routes
    NewUserRouter(env, getDb, publicRouter)


	// setup other routes below when needed
}