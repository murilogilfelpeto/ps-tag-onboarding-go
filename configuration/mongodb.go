package configuration

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userCollection *mongo.Collection
	ctx            = context.TODO()
)

type DatabaseConfiguration struct {
	Host       string
	Port       string
	User       string
	Password   string
	Database   string
	Collection string
}

func InitializeDatabase() (*mongo.Collection, error) {
	logger := GetLogger("mongodb")
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

	userCollection = client.Database(dbConfiguration.Database).Collection(dbConfiguration.Collection)
	return userCollection, nil
}

func GetContext() context.Context {
	return ctx
}

func getDatabaseConfiguration() DatabaseConfiguration {
	return DatabaseConfiguration{
		Host:       viper.GetString("database.host"),
		Port:       viper.GetString("database.port"),
		User:       viper.GetString("database.user"),
		Password:   viper.GetString("database.password"),
		Database:   viper.GetString("database.database"),
		Collection: viper.GetString("database.collection"),
	}
}

func getCredentials(dbConfig DatabaseConfiguration) options.Credential {
	return options.Credential{
		Username: dbConfig.User,
		Password: dbConfig.Password,
	}
}

func getDatabaseUrl(dbConfig DatabaseConfiguration) string {
	return "mongodb://" + dbConfig.Host + ":" + dbConfig.Port
}
