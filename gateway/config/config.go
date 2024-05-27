package config

type Config struct {
	Port            int
	UserServiceHost string
}

func NewConfig() *Config {
	return &Config{
		Port:            3000,
		UserServiceHost: ":3001",
	}
}
