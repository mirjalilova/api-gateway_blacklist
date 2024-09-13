package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	AUTH_PORT      string
	API_GATEWAY    string
	BLACKLIST_PORT string

	API_KEY string

	LOG_PATH string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.AUTH_PORT = cast.ToString(coalesce("AUTH_PORT", ":8090"))
	config.API_GATEWAY = cast.ToString(coalesce("API_GATEWAY", ":8090"))
	config.BLACKLIST_PORT = cast.ToString(coalesce("BLACKLIST_PORT", ":8090"))

	config.API_KEY = cast.ToString(coalesce("API_KEY", ""))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
