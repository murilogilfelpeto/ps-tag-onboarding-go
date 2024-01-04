package repository

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	logger         *configuration.Logger
	userCollection *mongo.Collection
	context        = configuration.GetContext()
)

func Initialize() {
	logger = configuration.GetLogger("repository")
	logger.Info("Initializing repository...")

	userCollection = configuration.GetUserCollection()
}
