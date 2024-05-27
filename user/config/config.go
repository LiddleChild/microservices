package config

import "os"

type Config struct {
	Port  string
	DbUri string
}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.Port = env("PORT", "3001")
	cfg.DbUri = env("DB_URI", "")

	return cfg
}

func env(key string, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}

	return val
}
