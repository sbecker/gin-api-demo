package util

import "os"

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}
