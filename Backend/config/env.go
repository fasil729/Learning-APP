package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	DatabaseURL               string
	Port                      string
	AccessTokenExpiryHour     int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour    int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret         string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret        string `mapstructure:"REFRESH_TOKEN_SECRET"`
	Gemin_Api_key		  string `mapstructure:"GEMINI_API_KEY"`
}

func Load() (*Env, error) {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Env{
		DatabaseURL:            os.Getenv("DATABASE_URL"),
		Port:                   os.Getenv("PORT"),
		AccessTokenExpiryHour:  getEnvAsInt("ACCESS_TOKEN_EXPIRY_HOUR", 1),
		RefreshTokenExpiryHour: getEnvAsInt("REFRESH_TOKEN_EXPIRY_HOUR", 24),
		AccessTokenSecret:      os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:     os.Getenv("REFRESH_TOKEN_SECRET"),
		Gemin_Api_key: 		os.Getenv("GEMINI_API_KEY"),
	}, nil
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Invalid value for %s: %s", key, valueStr)
		return defaultValue
	}
	return valueInt
}

