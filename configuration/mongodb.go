package configuration

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userCollection *mongo.Collection
	ctx            = context.TODO()
	credentials    = options.Credential{
		Username: "root",
		Password: "root",
	}
)

func InitializeDatabase() (*mongo.Collection, error) {
	logger := GetLogger("mongodb")
	logger.Info("Initializing database...")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)
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

	userCollection = client.Database("onboarding").Collection("users")
	return userCollection, nil
}

func GetContext() context.Context {
	return ctx
}
