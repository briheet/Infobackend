package config

import (
	"fmt"
	"os"

	"github.com/briheet/basicBackend/internal/models"
)

var envs = initConfig()

func initConfig() models.Config {
	return models.Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "briheet"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "projectmanager"),
		JWTSecret:  getEnv("JWT_SECRET", "randomjwtsecretkey"),
	}
}

func getEnv(name, fallback string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}

	return fallback
}
