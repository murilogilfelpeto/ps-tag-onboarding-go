package configuration

import (
	"errors"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func LoadConfiguration() error {
	loadEnvironmentVariables()
	err := viper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			logger.Error("Config file not found")
			return err
		}
		logger.Errorf("Error reading config file: %v", err)
		return err
	}
	return nil
}

func loadEnvironmentVariables() {
	environment := os.Getenv("ENVIRONMENT")
	viper.AddConfigPath("configuration")
	switch strings.ToUpper(environment) {
	case "DOCKER-MACOS":
		viper.SetConfigName("docker-macos")
		viper.SetConfigType("toml")

	default:
		viper.SetConfigName("local")
		viper.SetConfigType("toml")
	}
}
