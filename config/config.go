package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App  app
		Log  log
		Http http
		Pg   pg
		Jwt  jwt
	}

	app struct {
		Name    string `env:"APP_NAME" envDefault:"tenant-gate"`
		Version string `env:"VERSION" envDefault:"1.0.0"`
	}

	http struct {
		Port       string `env:"HTTP_PORT" envDefault:"8080"`
		UsePrefork bool   `env:"HTTP_USE_PREFORK" envDefault:"false"`
	}

	log struct {
		Level string `env:"LOG_LEVEL" envDefault:"info"`
	}

	pg struct {
		Host string `env:"PG_HOST" envDefault:"localhost"`
		Port string `env:"PG_PORT" envDefault:"5432"`
	}

	jwt struct {
		Secret      string        `env:"JWT_SECRET" envDefault:"JWT_SECRET"`
		TokenExpiry time.Duration `env:"JWT_TOKEN_EXPIRY" envDefault:"24h"`
	}

	// Metrics
	// Swagger
	// Tracing
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return cfg, nil
}
