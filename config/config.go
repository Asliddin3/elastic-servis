package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment       string
	LogLevel          string
	MONGOHost      string
	MONGOPort      string
	MONGODatabase  string
	MONGOUser      string
	MONGOPassword  string
	POLLServiceHost string
	POLLServicePort string
}

func LoadConfig() *Config {
	c := &Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.MONGOHost = cast.ToString(getOrReturnDefault("MONGO_HOST", "localhost"))
	c.MONGOPort = cast.ToString(getOrReturnDefault("MONGO_PORT", 27017))
	c.MONGODatabase = cast.ToString(getOrReturnDefault("MONGO_DATABASE", "polldb"))
	c.MONGOUser = cast.ToString(getOrReturnDefault("MONGO_USER", "asliddin"))
	c.MONGOPassword = cast.ToString(getOrReturnDefault("MONGO_PASSWORD", "compos1995"))
	c.POLLServiceHost = cast.ToString(getOrReturnDefault("POLL_SERVICE_HOST", "localhost"))
	c.POLLServicePort = cast.ToString(getOrReturnDefault("POLL_SERVICE_PORT", "8000"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
