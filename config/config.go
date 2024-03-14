package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppConfig      AppConfig
	DatabaseConfig DatabaseConfig
}

type AppConfig struct {
	// The port the server will listen on
	Port string `envconfig:"PORT" default:"8080"`
	// The environment the server is running in
	Environment string `envconfig:"ENVIRONMENT" default:"dev"`
}

type DatabaseConfig struct {
	// The database connection string
	ConnectionString string `envconfig:"DATABASE_URL"`
	// The database user
	User string `envconfig:"DATABASE_USER"`
	// The database password
	Password string `envconfig:"DATABASE_PASSWORD"`
	// The database name
	Database string `envconfig:"DATABASE_NAME"`
}

var config Config

func Get() *Config {
	return &config
}

func New() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	envconfig.MustProcess("", &config)
}
