package service

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/repository"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models/exceptions"
)

func SaveUser(user models.User) (models.User, error) {
	createdUser, err := repository.Save(user)
	if err != nil {
		logger.Errorf("Error persisting user: %v", err)
		return models.User{}, &exceptions.UserAlreadyExistErr{Message: "User already exists: " + user.GetFullName()}
	}
	return createdUser, nil
}

func GetUserById(id string) (models.User, error) {
	newUser, err := models.NewUser(id, "Murilo", "Felpeto", "murilo@wexinc.com", 30)
	if err != nil {
		logger.Errorf("Error finding user: %v", err)
		return models.User{}, err
	}

	return newUser, nil
}
