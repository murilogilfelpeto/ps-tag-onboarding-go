package repository

import (
	"context"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var logger = configuration.NewLogger("repository")

type Repository interface {
	Save(ctx context.Context, user models.User) (models.User, error)
	GetUserById(ctx context.Context, id string) (models.User, error)
	GetUserByFullName(ctx context.Context, firstName string, lastName string) (models.User, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client, dbName string, collectionName string) Repository {
	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	return &repository{
		collection: collection,
	}
}
