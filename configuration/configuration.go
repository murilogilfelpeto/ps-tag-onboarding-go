package configuration

import "go.mongodb.org/mongo-driver/mongo"

var (
	logger     *Logger
	collection *mongo.Collection
)

func Init() error {
	var err error
	collection, err = InitializeDatabase()
	if err != nil {
		logger.Errorf("Error initializing database: %v", err)
		return err
	}
	return nil
}

func GetLogger(prefix string) *Logger {
	return NewLogger(prefix)
}

func GetUserCollection() *mongo.Collection {
	return collection
}
