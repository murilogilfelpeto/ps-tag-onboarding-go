package configuration

import (
	"context"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
		Host:       viper.GetString("database.host"),
		Port:       viper.GetString("database.port"),
		Username:   viper.GetString("database.user"),
		Password:   viper.GetString("database.password"),
		Database:   viper.GetString("database.database"),
		Collection: viper.GetString("database.collection"),
	}
}

func getCredentials(dbConfig DatabaseConfig) options.Credential {
	return options.Credential{
		Username: dbConfig.Username,
		Password: dbConfig.Password,
	}
}

func getDatabaseUrl(dbConfig DatabaseConfig) string {
	return "mongodb://" + dbConfig.Host + ":" + dbConfig.Port
}
