package configuration

import (
	"errors"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func LoadConfiguration() error {
	logger := NewLogger("mongodb")
	loadEnvironmentVariables()
	err := viper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			logger.Error("Config file not found")
			return err
		}
	}
	return nil
}

func loadEnvironmentVariables() {
	environment := os.Getenv("ENVIRONMENT")
	switch strings.ToUpper(environment) {
	case "DOCKER-MACOS":
		viper.SetConfigName("docker-macos")
		viper.SetConfigType("toml")
		viper.AddConfigPath("configuration/files")
	default:
		viper.SetConfigName("local")
		viper.SetConfigType("toml")
		viper.AddConfigPath("configuration/files")
	}
}
