package config

import "os"

type Config struct {
	Port            string
	UserServiceHost string
}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.Port = env("PORT", "3000")
	cfg.UserServiceHost = env("USER_SERVICE_HOST", ":3001")

	return cfg
}

func env(key string, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}

	return val
}
