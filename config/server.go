package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Version     string
	Env         string
	Port        string
	CorsMethods string
	CorsOrigins string
}

var Server ServerConfig = ServerConfig{
	Version:     "1.0.0",
	Env:         "default",
	Port:        "8080",
	CorsMethods: "OPTIONS,GET,POST,PUT,PATCH,DELETE",
	CorsOrigins: "*",
}

func init() {
	env := getEnv("ENV", "default")
	var envFilePath string
	if env == "default" {
		return
	} else {
		envFilePath = ".env." + env
	}

	err := godotenv.Load(envFilePath)
	if err != nil {
		return
	}

	Server.Env = env
	Server.Port = getEnv("PORT", Server.Port)
	Server.CorsMethods = getEnv("ALLOW_CORS_METHODS", Server.CorsMethods)
	Server.CorsOrigins = getEnv("ALLOW_CORS_ORIGINS", Server.CorsOrigins)
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
