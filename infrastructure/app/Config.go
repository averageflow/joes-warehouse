package app

import (
	"os"
)

type ApplicationConfig struct {
	ApplicationMode    string
	DatabaseConnection string
}

// GetConfig initializes the configuration instance to the values described in the config.toml file.
func GetConfig() *ApplicationConfig {
	return &ApplicationConfig{
		ApplicationMode:    os.Getenv("APPLICATION_MODE"),
		DatabaseConnection: os.Getenv("DATABASE_CONNECTION"),
	}
}
