package config

import (
	"github.com/caarlos0/env/v7"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Host string `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	Port string `env:"SERVER_PORT" envDefault:"8000"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"3306"`
	User     string `env:"DB_USER" envDefault:"ims"`
	Password string `env:"DB_PASSWORD" envDefault:"passw0rd"`
	Name     string `env:"DB_DATABASE" envDefault:"ims"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
