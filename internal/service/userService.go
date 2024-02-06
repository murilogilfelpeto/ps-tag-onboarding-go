package service

import (
	"context"
	"errors"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/repository"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models/exceptions"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
)

type Service interface {
	SaveUser(ctx context.Context, user models.User) (models.User, error)
	GetUserById(ctx context.Context, id string) (models.User, error)
}

type service struct {
	repository repository.Repository
}

func NewUserService(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (srv *service) SaveUser(ctx context.Context, user models.User) (models.User, error) {
	logger.Infof("Saving user %s", user.GetFullName())
	userByFullName, err := srv.repository.GetUserByFullName(ctx, user.GetFirstName(), user.GetLastName())
	if err == nil && userByFullName.GetID() != "" {
		var serverSelectionError topology.ServerSelectionError
		if errors.As(err, &serverSelectionError) {
			return models.User{}, &exceptions.DatabaseConnectionErr{Err: errors.New("Something went wrong: " + err.Error())}
		}
		logger.Errorf("User already exists: %v", user.GetFullName())
		return models.User{}, &exceptions.UserAlreadyExistErr{Err: errors.New("User already exists: " + user.GetFullName())}
	}
	createdUser, err := srv.repository.Save(ctx, user)
	if err != nil {
		var serverSelectionError topology.ServerSelectionError
		if errors.As(err, &serverSelectionError) {
			return models.User{}, &exceptions.DatabaseConnectionErr{Err: errors.New("Something went wrong: " + err.Error())}
		}
		logger.Errorf("Error persisting user: %v", err)
		return models.User{}, &exceptions.UserValidationErr{Err: errors.New("Error persisting user: " + user.GetFullName())}
	}
	logger.Infof("User %s saved successfully", user.GetFullName())
	return createdUser, nil
}

func (srv *service) GetUserById(ctx context.Context, id string) (models.User, error) {
	logger.Infof("Getting user by id %s", id)
	user, err := srv.repository.GetUserById(ctx, id)
	if err != nil {
		var serverSelectionError topology.ServerSelectionError
		if errors.As(err, &serverSelectionError) {
			return models.User{}, &exceptions.DatabaseConnectionErr{Err: errors.New("Something went wrong: " + err.Error())}
		}
		logger.Errorf("Error finding user: %v", err)
		return models.User{}, &exceptions.UserNotFoundErr{Err: errors.New("User not found: " + id)}
	}
	logger.Infof("User %s found successfully", id)
	return user, nil
}
