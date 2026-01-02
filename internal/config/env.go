package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	_ = godotenv.Load()
}

func Get(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
