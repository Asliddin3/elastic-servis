package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cast"
)

type Config struct {
	Environment     string
	LogLevel        string
	MONGOHost       string
	MONGOPort       string
	MONGODatabase   string
	MONGOUser       string
	MONGOPassword   string
	POLLServiceHost string
	POLLServicePort string
	Logger          zerolog.Logger
	POSTGRES_HOST   string
	POSTGRES_PORT   int
	POSTGRES_DB     string
	POSTGRES_USER   string
	POSTGRES_PASS   string
}

func LoadConfig() *Config {
	c := &Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.MONGOHost = cast.ToString(getOrReturnDefault("MONGO_HOST", "127.0.0.1"))
	c.MONGOPort = cast.ToString(getOrReturnDefault("MONGO_PORT", 27017))
	c.MONGODatabase = cast.ToString(getOrReturnDefault("MONGO_DATABASE", "polldb"))
	c.MONGOUser = cast.ToString(getOrReturnDefault("MONGO_USER", "asliddin"))
	c.MONGOPassword = cast.ToString(getOrReturnDefault("MONGO_PASSWORD", "compos1995"))
	c.POLLServiceHost = cast.ToString(getOrReturnDefault("POLL_SERVICE_HOST", "localhost"))
	c.POLLServicePort = cast.ToString(getOrReturnDefault("POLL_SERVICE_PORT", "8070"))
	c.POSTGRES_HOST = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.POSTGRES_PORT = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.POSTGRES_USER = cast.ToString(getOrReturnDefault("POSTGRES_USER", "asliddin"))
	c.POSTGRES_PASS = cast.ToString(getOrReturnDefault("POSTGRES_PASS", "compos1995"))
	c.POSTGRES_DB = cast.ToString(getOrReturnDefault("POSTGRES_DB", "postdb"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
