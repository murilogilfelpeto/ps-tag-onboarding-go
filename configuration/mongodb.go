package configuration

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Collection string
}

func NewDatabaseConfig(host, port, username, password, database, collection string) *DatabaseConfig {
	return &DatabaseConfig{
		Host:       host,
		Port:       port,
		Username:   username,
		Password:   password,
		Database:   database,
		Collection: collection,
	}
}

func Connect(ctx context.Context, config *DatabaseConfig) (*mongo.Client, error) {
	logger := NewLogger("mongodb")
	logger.Info("Initializing database...")

	credentials := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}

	clientOptions := options.Client().ApplyURI("mongodb://" + config.Host + ":" + config.Port).SetAuth(credentials)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		logger.Errorf("Error connecting to database: %v", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Errorf("Error pinging database: %v", err)
		return nil, err
	}
	return client, nil
}
