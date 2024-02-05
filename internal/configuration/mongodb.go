package configuration

import (
	"context"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConfig struct {
	host       string
	port       string
	username   string
	password   string
	database   string
	collection string
}

func Connect(ctx context.Context) (*mongo.Client, error) {
	logger.Info("Initializing database...")

	dbConfiguration := getDatabaseConfiguration()
	credentials := getCredentials(dbConfiguration)
	dbUrl := getDatabaseUrl(dbConfiguration)
	clientOptions := options.Client().ApplyURI(dbUrl).SetAuth(credentials)
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

func getDatabaseConfiguration() DatabaseConfig {
	return DatabaseConfig{
		host:       viper.GetString("database.host"),
		port:       viper.GetString("database.port"),
		username:   viper.GetString("database.user"),
		password:   viper.GetString("database.password"),
		database:   viper.GetString("database.database"),
		collection: viper.GetString("database.collection"),
	}
}

func getCredentials(dbConfig DatabaseConfig) options.Credential {
	return options.Credential{
		Username: dbConfig.username,
		Password: dbConfig.password,
	}
}

func getDatabaseUrl(dbConfig DatabaseConfig) string {
	return "mongodb://" + dbConfig.host + ":" + dbConfig.port
}
