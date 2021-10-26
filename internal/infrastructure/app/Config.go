package app

import (
	"os"
)

// ApplicationConfig represents the necessary configuration for the application.
type ApplicationConfig struct {
	ApplicationMode    string
	DatabaseConnection string
	WebAssetLocation   string
}

// GetConfig will return a populated ApplicationConfig object from variables in the OS environment.
func GetConfig() *ApplicationConfig {
	return &ApplicationConfig{
		ApplicationMode:    os.Getenv("APPLICATION_MODE"),
		DatabaseConnection: os.Getenv("DATABASE_CONNECTION"),
		WebAssetLocation:   os.Getenv("WEB_ASSET_LOCATION"),
	}
}
