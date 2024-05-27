package config

import "os"

type Config struct {
	Port string
}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.Port = env("PORT", "3001")

	return cfg
}

func env(key string, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}

	return val
}
