package config

import (
	"log"
	"os"
)

type Config struct {
	ServerPort  string
	DbUrl       string
	LogLevel    string
	AppName     string
	Environment string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort:  GetEnvOrDie("PORT"),
		DbUrl:       GetEnvOrDie("DB_URL"),
		LogLevel:    GetEnvOrDie("LOG_LEVEL"),
		Environment: GetEnvOrDie("ENV"),
		AppName:     GetEnvOrDie("APP_NAME"),
	}
}

func GetEnvOrDie(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("could not get %s from environment variables", key)
	}

	return value
}
