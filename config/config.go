package config

import (
	"log"
	"os"
)

type Config struct {
	ServerPort string
	DbUrl      string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort: GetEnvOrDie("PORT"),
		DbUrl:      GetEnvOrDie("DB_URL"),
	}
}

func GetEnvOrDie(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("could not get %s from env", key)
	}

	return value
}
