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
	// Cargar variables de entorno desde el archivo .env si está en local
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("No .env file found")
	}

	config := &Config{
		App: AppConfig{
			Env: os.Getenv("ENV"),
		},
		Server: ServerConfig{
			Address:   os.Getenv("SERVER_ADDRESS"),
			JWTSecret: os.Getenv("JWT_SECRET"),
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
