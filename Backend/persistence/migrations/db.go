package migrations

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hackathon.com/leariningApp/domain"
)

var db *gorm.DB
var once sync.Once

func GetDB() *gorm.DB {
	once.Do(func() {
		// Load environment variables from .env file
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}

		// Get the database URL from the environment
		dbURL := os.Getenv("DATABASE_URL")

		// Connect to the database
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database!")
		}

		// Auto-migrate the model
		db.AutoMigrate(&domain.User{}, &domain.Subject{}, &domain.Experiment{}, &domain.Quiz{})
	})
	return db
}
