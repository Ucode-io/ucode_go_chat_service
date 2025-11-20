package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	LogLevel string

	HTTPPort string

	PostgresHost           string
	PostgresPort           int
	PostgresUser           string
	PostgresPassword       string
	PostgresDatabase       string
	PostgresMaxConnections int32
}

func Load() Config {
	if err := godotenv.Load("/app/.env"); err != nil {
		if err := godotenv.Load(".env"); err != nil {
			log.Println("No .env file found")
		}
		log.Println("No /app/.env file found")
	}

	cfg := Config{}

	cfg.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "DEBUG"))

	cfg.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ""))

	cfg.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", ""))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 0))
	cfg.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", ""))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", ""))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", ""))
	cfg.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 200))

	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue any) any {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
