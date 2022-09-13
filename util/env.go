package util

import "os"

var (
	Environment = GetEnv("ENVIRONMENT", "production")
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return value
}
