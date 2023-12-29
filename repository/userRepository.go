package repository

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/repository/mapper"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models"
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
