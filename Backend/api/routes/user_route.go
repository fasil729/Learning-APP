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

func NewUserRouter(env *config.Env, getDb func() *gorm.DB, group *gin.RouterGroup) {
	userRepository := repositories.NewUserRepository(getDb)
	userService := services.NewUserService(userRepository, env.RefreshTokenExpiryHour, env.AccessTokenExpiryHour, env.AccessTokenSecret)
	userController := controllers.NewUserController(userService)

	authMiddleware := middlewares.AuthMiddleware(env.AccessTokenSecret)

	// Auth routes
	authGroup := group.Group("/users")
	{
		authGroup.POST("/register", userController.RegisterUser)
		authGroup.POST("/admin/register", authMiddleware, middlewares.RoleMiddleware("admin"), userController.RegisterAdmin)
		authGroup.POST("/login", userController.SignIn)
		authGroup.POST("/logout", authMiddleware, userController.Logout)
		authGroup.GET("/me", authMiddleware, userController.GetMe)
		authGroup.PUT("/update", authMiddleware, userController.UpdateUser)
		authGroup.DELETE("/delete/:id", authMiddleware, userController.DeleteUser)
	}

}
