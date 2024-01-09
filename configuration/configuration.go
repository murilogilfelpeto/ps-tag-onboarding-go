package configuration

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppConfig struct {
	Logger   *Logger
	Database *mongo.Client
	Context  context.Context
}

func Init() (*AppConfig, error) {
	logger := NewLogger("configuration")
	logger.Info("Initializing configuration...")

	appConfig := &AppConfig{
		Logger:  logger,
		Context: context.Background(),
	}

	err := appConfig.initDatabase()
	if err != nil {
		logger.Errorf("Error initializing database: %v", err)
		return nil, err
	}

	return appConfig, nil
}

func (config *AppConfig) initDatabase() error {
	dbConfig := NewDatabaseConfig("localhost", "27017", "root", "root", "onboarding", "users")
	database, err := Connect(config.Context, dbConfig)
	if err != nil {
		return err
	}

	config.Database = database
	return nil
}
