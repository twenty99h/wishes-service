package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/twenty99h/wishes-service/pkg/httpserver"
	"github.com/twenty99h/wishes-service/pkg/postgres"
)

type App struct {
	Name    string `envconfig:"APP_NAME" required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App      App
	Postgres postgres.Config
	HTTP     httpserver.Config
}

func New() (Config, error) {
	var config Config

	if err := godotenv.Load(".env"); err != nil {
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	if err := envconfig.Process("", &config); err != nil {
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}

	return config, nil
}
