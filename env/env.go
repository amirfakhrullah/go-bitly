package env

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var DB_URL_STRING string

func getEnvSafely(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic("environment variable not found: " + k)
	}
	return v
}

func init() {
	DB_URL_STRING = getEnvSafely("DB_URL_STRING")
}
