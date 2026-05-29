package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct  {
	Port string
	DbPath string
}

func loadConfig() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
		DbPath: getEnv("DB_URI", "./data/orders.db"),
	}
}

func getEnv(key, defval string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defval
}

func loadTemplates(router * gin.Engine) error {
	
}