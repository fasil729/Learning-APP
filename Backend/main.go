package main

import (
	route "Brilliant/api/routes"
	"Brilliant/config"
	"Brilliant/persistence/migrations"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	// Setup routes
	route.Setup(env, getDb, router)

	router.Run()
}