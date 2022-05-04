package config

import (
	"log"
	"os"

	_env "github.com/joho/godotenv"
)

func init() {
	err := _env.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func get(key string) string {
	if os.Getenv(key) == "" {
		log.Fatal("Environment variable " + key + " is not set")
	}

	return os.Getenv(key)
}
