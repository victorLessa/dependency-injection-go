package config

import (
	"github.com/joho/godotenv"
  "log"
  "os"
)

type AppConfig struct {
	AppName string
	AppHost string
	AppPort string
	AppPrefix string
	AppDatabase string
	AppUsername string
	AppPassword string
	AppConnection string
}

func LoadConfig() AppConfig {
	err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	config := AppConfig{
		AppName: os.Getenv("APP_NAME"),
		AppHost: os.Getenv("APP_HOST"),
		AppPort: os.Getenv("APP_PORT"),
		AppPrefix: os.Getenv("APP_PREFIX"),

	}

	return config

}