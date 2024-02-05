package configuration

import (
	"context"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppConfig struct {
	Database *mongo.Client
	Context  context.Context
}

func Init() (*AppConfig, error) {
	logger.Info("Initializing configuration...")

	err := LoadConfiguration()
	if err != nil {
		logger.Errorf("Error loading configuration: %v", err)
		return nil, err
	}
	appConfig := &AppConfig{
		Context: context.Background(),
	}

	err = appConfig.initDatabase()
	if err != nil {
		logger.Errorf("Error initializing database: %v", err)
		return nil, err
	}

	return appConfig, nil
}

func (config *AppConfig) initDatabase() error {
	database, err := Connect(config.Context)
	if err != nil {
		return err
	}

	config.Database = database
	return nil
}
