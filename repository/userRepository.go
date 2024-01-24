package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/repository/entities"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/repository/mapper"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var logger = configuration.NewLogger("userRepository")

type Repository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, dbName string, collectionName string) *Repository {
	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	return &Repository{
		collection: collection,
	}
}
func (repo *Repository) Save(ctx context.Context, user models.User) (models.User, error) {
	logger.Infof("Saving user %s", user.GetFirstName())
	newUser := mapper.UserToUserEntity(user)
	_, err := repo.collection.InsertOne(ctx, newUser)
	if err != nil {
		logger.Errorf("Error saving user %s: %v", user.GetFirstName(), err)
		return models.User{}, err
	}
	logger.Infof("User %s saved successfully", user.GetFirstName())
	createdUser, err := mapper.UserEntityToUser(newUser)
	return createdUser, nil
}

func (repo *Repository) GetUserById(ctx context.Context, id string) (models.User, error) {
	logger.Infof("Getting user by id %s", id)
	var userEntity entities.UserEntity

	uid, err := uuid.Parse(id)
	if err != nil {
		logger.Errorf("Error parsing user id %s: %v", id, err)
		return models.User{}, err
	}

	err = repo.collection.FindOne(ctx, bson.M{"_id": uid}).Decode(&userEntity)
	if err != nil {
		logger.Errorf("Error getting user by id %s: %v", id, err)
		return models.User{}, err
	}
	logger.Infof("User %s found successfully", id)
	user, err := mapper.UserEntityToUser(userEntity)
	return user, nil
}

func (repo *Repository) GetUserByFullName(ctx context.Context, firstName string, lastName string) (models.User, error) {
	logger.Infof("Getting user by full name %s %s", firstName, lastName)
	var userEntity entities.UserEntity
	err := repo.collection.FindOne(ctx, bson.M{"first_name": firstName, "last_name": lastName}).Decode(&userEntity)
	if err != nil {
		logger.Errorf("Error getting user by full name %s %s: %v", firstName, lastName, err)
		return models.User{}, err
	}
	logger.Infof("User %s %s found successfully", firstName, lastName)
	user, _ := mapper.UserEntityToUser(userEntity)
	return user, nil
}
