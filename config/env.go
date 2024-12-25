package config

import "os"

const (
	SecretKey = "SECRET_KEY"
	Port      = "PORT"
)

func GetEnv(name, defaultValue string) string {
	env := os.Getenv(name)

	if env == "" {
		return defaultValue
	}

	return env
}
