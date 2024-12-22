package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DebugMode = "debug"
	TestMode = "test"
	ReleaseMode = "release"
)

type Config struct {
	Environment string

	ServerHost string
	ServerPort string

	PostgresHost           string
	PostgresUser           string
	PostgresDatabase       string
	PostgresPassword       string
	PostgresPort           string
	PostgresMaxConnections int32

	DefaultOffset int
	DefaultLimit  int

	SecretKey string
}

func Load() Config {

	if err := godotenv.Load("./app.env"); err != nil {
		fmt.Println("No .env file found")
	}

	cfg := Config{}

	cfg.ServerHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	cfg.ServerPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))

	cfg.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "db"))
	cfg.PostgresPort = cast.ToString(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	cfg.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "yourusername"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "yourpassword"))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "yourdatabase"))
	cfg.PostgresMaxConnections = 20

	cfg.DefaultOffset = cast.ToInt(getOrReturnDefaultValue("OFFSET", 0))
	cfg.DefaultLimit = cast.ToInt(getOrReturnDefaultValue("LIMIT", 10))

	cfg.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", "secret"))

	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
