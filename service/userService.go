package service

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/repository"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models/exceptions"
)

func SaveUser(user models.User) (models.User, error) {
	userByFullName, err := repository.GetUserByFullName(user.GetFirstName(), user.GetLastName())
	if err == nil && userByFullName.GetID() != "" {
		logger.Errorf("Error persisting user: %v", err)
		return models.User{}, &exceptions.UserAlreadyExistErr{Message: "User already exists: " + user.GetFullName()}
	}
	createdUser, err := repository.Save(user)
	if err != nil {
		logger.Errorf("Error persisting user: %v", err)
		return models.User{}, &exceptions.UserValidationErr{Message: "Error persisting user: " + user.GetFullName()}
	}
	return createdUser, nil
}

func GetUserById(id string) (models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		logger.Errorf("Error finding user: %v", err)
		return models.User{}, &exceptions.UserNotFoundErr{Message: "User not found: " + id}
	}
	return user, nil
}
