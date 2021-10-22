package app

import (
	"log"

	"github.com/spf13/viper"
)

type ApplicationConfig struct {
}

// ConfigSetup will prepare and setup the viper instance to the correct config file.
func ConfigSetup(configName, configPath string) {
	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)
}

// ConfigUpdate initializes the configuration instance to the values described in the config.toml file.
func GetConfig() ApplicationConfig {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("fatal error config file: %s", err)
	}

	validateVariablesAreSet([]string{})

	return ApplicationConfig{}
}

// ValidateVariablesAreSet will assert the existence of each variable,
// and kill the application when a wanted variable does not exist in the config.
func validateVariablesAreSet(variables []string) {
	for i := range variables {
		if !viper.IsSet(variables[i]) {
			log.Fatalf("%s variable was not set!\nAborting application start!", variables[i])
		}
	}
}
