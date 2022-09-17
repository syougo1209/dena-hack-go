package config

import "github.com/caarlos0/env/v6"

type Config struct {
	MySQLUser     string `env:"MYSQL_USER"`
	MYSQLPassword string `env:"MYSQL_PASSWORD"`
	MYSQLAddr     string `env:"MYSQL_ADDR"`
	MYSQLDbName   string `env:"MYSQL_DATABASE"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
