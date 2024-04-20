package main

import (
	route "Brilliant/api/routes"
	"Brilliant/config"
	"Brilliant/persistence/migrations"
	_ "Brilliant/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Brilliant Learning APP API
// @version 1.0
// @description This is the API for the Brilliant application
// @host localhost:8080
func main() {
	// Load environment variables
	env, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Setup database
	getDb := migrations.GetDb

	router := gin.Default()
    

	// Global CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	router.Use(cors.New(config))

	// Swagger documentation
	// url := ginSwagger.URL("/docs/swagger.json")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	route.Setup(env, getDb, router)


	router.Run()
}
