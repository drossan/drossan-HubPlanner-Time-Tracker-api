package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App    AppConfig
	Server ServerConfig
}

type ServerConfig struct {
	Address   string
	JWTSecret string
}

type AppConfig struct {
	Env string
}

func LoadConfig() *Config {
	// Intentar cargar variables de entorno desde el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("No .env file found, using environment variables")
	}

	config := &Config{
		App: AppConfig{
			Env: GetEnv("ENV", "development"),
		},
		Server: ServerConfig{
			Address:   GetEnv("SERVER_ADDRESS", ":8080"),
			JWTSecret: GetEnv("JWT_SECRET", ""),
		},
	}

	return config
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
