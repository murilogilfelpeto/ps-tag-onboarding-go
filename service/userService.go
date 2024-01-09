package service

import (
	"context"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models/exceptions"
)

func (srv *service) SaveUser(ctx context.Context, user models.User) (models.User, error) {
	logger.Infof("Saving user %s", user.GetFullName())
	userByFullName, err := srv.repository.GetUserByFullName(ctx, user.GetFirstName(), user.GetLastName())
	if err == nil && userByFullName.GetID() != "" {
		logger.Errorf("User already exists: %v", user.GetFullName())
		return models.User{}, &exceptions.UserAlreadyExistErr{Message: "User already exists: " + user.GetFullName()}
	}
	createdUser, err := srv.repository.Save(ctx, user)
	if err != nil {
		logger.Errorf("Error persisting user: %v", err)
		return models.User{}, &exceptions.UserValidationErr{Message: "Error persisting user: " + user.GetFullName()}
	}
	logger.Infof("User %s saved successfully", user.GetFullName())
	return createdUser, nil
}

func (srv *service) GetUserById(ctx context.Context, id string) (models.User, error) {
	logger.Infof("Getting user by id %s", id)
	user, err := srv.repository.GetUserById(ctx, id)
	if err != nil {
		logger.Errorf("Error finding user: %v", err)
		return models.User{}, &exceptions.UserNotFoundErr{Message: "User not found: " + id}
	}
	logger.Infof("User %s found successfully", id)
	return user, nil
}
