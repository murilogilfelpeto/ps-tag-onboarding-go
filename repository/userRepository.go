package repository

import (
	"github.com/google/uuid"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/repository/entities"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/repository/mapper"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Save(user models.User) (models.User, error) {
	logger.Infof("Saving user %s", user.GetFirstName())
	newUser := mapper.UserToUserEntity(user)
	_, err := userCollection.InsertOne(context, newUser)
	if err != nil {
		logger.Errorf("Error saving user %s: %v", user.GetFirstName(), err)
		return models.User{}, err
	}
	logger.Infof("User %s saved successfully", user.GetFirstName())
	createdUser, err := mapper.UserEntityToUser(newUser)
	return createdUser, nil
}

func GetUserById(id string) (models.User, error) {
	logger.Infof("Getting user by id %s", id)
	var userEntity entities.UserEntity

	uid, err := uuid.Parse(id)
	if err != nil {
		logger.Errorf("Error parsing user id %s: %v", id, err)
		return models.User{}, err
	}

	err = userCollection.FindOne(context, bson.M{"_id": uid}).Decode(&userEntity)
	if err != nil {
		logger.Errorf("Error getting user by id %s: %v", id, err)
		return models.User{}, err
	}
	logger.Infof("User %s found successfully", id)
	user, err := mapper.UserEntityToUser(userEntity)
	return user, nil
}

func GetUserByFullName(firstName string, lastName string) (models.User, error) {
	logger.Infof("Getting user by full name %s %s", firstName, lastName)
	var userEntity entities.UserEntity
	err := userCollection.FindOne(context, bson.M{"first_name": firstName, "last_name": lastName}).Decode(&userEntity)
	if err != nil {
		logger.Errorf("Error getting user by full name %s %s: %v", firstName, lastName, err)
		return models.User{}, err
	}
	logger.Infof("User %s %s found successfully", firstName, lastName)
	user, _ := mapper.UserEntityToUser(userEntity)
	return user, nil
}
