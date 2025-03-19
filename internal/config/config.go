package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	TMDBApiKey    string
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassword    string
	DbName        string
	Port          string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}

	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))

	return &Config{
		TMDBApiKey:    getEnv("TMDB_API_KEY", ""),
		DbHost:        getEnv("DB_HOST", "localhost"),
		DbPort:        getEnv("DB_PORT", "5432"),
		DbUser:        getEnv("DB_USER", "postgres"),
		DbPassword:    getEnv("DB_PASSWORD", "postgres"),
		DbName:        getEnv("DB_NAME", "tmdb_cli"),
		Port:          getEnv("PORT", "8080"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       redisDB,
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
