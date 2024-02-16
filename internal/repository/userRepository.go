package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/repository/entities"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/repository/mapper"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Save(ctx context.Context, user models.User) (*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByFullName(ctx context.Context, firstName string, lastName string) (*models.User, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, dbName string, collectionName string) UserRepository {
	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	return &repository{
		collection: collection,
	}
}
func (repo *repository) Save(ctx context.Context, user models.User) (*models.User, error) {
	logger.Infof("Saving user %s", user.GetFirstName())
	newUser := mapper.UserToUserEntity(user)
	_, err := repo.collection.InsertOne(ctx, newUser)
	if err != nil {
		logger.Errorf("Error saving user %s: %v", user.GetFirstName(), err)
		return nil, err
	}
	logger.Infof("User %s saved successfully", user.GetFirstName())
	createdUser, err := mapper.UserEntityToUser(newUser)
	return &createdUser, nil
}

func (repo *repository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	logger.Infof("Getting user by id %s", id)
	var userEntity entities.UserEntity

	uid, err := uuid.Parse(id)
	if err != nil {
		logger.Errorf("Error parsing user id %s: %v", id, err)
		return nil, err
	}

	err = repo.collection.FindOne(ctx, bson.M{"_id": uid}).Decode(&userEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Infof("User %s not found", id)
			return nil, nil
		}
		logger.Errorf("Error getting user by id %s: %v", id, err)
		return nil, err
	}
	logger.Infof("User %s found successfully", id)
	user, err := mapper.UserEntityToUser(userEntity)
	return &user, nil
}

func (repo *repository) GetUserByFullName(ctx context.Context, firstName string, lastName string) (*models.User, error) {
	logger.Infof("Getting user by full name %s %s", firstName, lastName)
	var userEntity entities.UserEntity
	err := repo.collection.FindOne(ctx, bson.M{"first_name": firstName, "last_name": lastName}).Decode(&userEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Infof("User %s %s not found", firstName, lastName)
			return nil, nil
		}
		logger.Errorf("Error getting user by full name %s %s: %v", firstName, lastName, err)
		return nil, err
	}
	logger.Infof("User %s %s found successfully", firstName, lastName)
	user, _ := mapper.UserEntityToUser(userEntity)
	return &user, nil
}
