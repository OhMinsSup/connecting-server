package app

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func NewEnv() {
	env := os.Getenv("APP_ENV")
	//allowOrigins := []string{""}
	switch env {
	case "production":
		loadEnv(".env")
		break
	case "development":
		//allowOrigins = append(allowOrigins, "")
		loadEnv(".env.dev")
		break
	default:
		loadEnv(".env.dev")
		break
	}
}
