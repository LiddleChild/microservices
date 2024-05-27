package config

type Config struct {
	Port int
}

func NewConfig() *Config {
	return &Config{
		Port: 3001,
	}
}
