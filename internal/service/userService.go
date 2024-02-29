package service

import (
	"context"
	"errors"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/repository"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models/exceptions"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
)

type UserService interface {
	SaveUser(ctx context.Context, user models.User) (*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
}

type service struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &service{
		userRepository: repository,
	}
}

const errorMessage = "something went wrong"

func (srv *service) SaveUser(ctx context.Context, user models.User) (*models.User, error) {
	logger.Infof("Saving user %s", user.GetFullName())
	userByFullName, err := srv.userRepository.GetUserByFullName(ctx, user.GetFirstName(), user.GetLastName())

	if userByFullName != nil {
		if err == nil {
			logger.Errorf("User already exists: %v", user.GetFullName())
			return nil, &exceptions.UserAlreadyExistsErr{Err: errors.New("User already exists: " + user.GetFullName())}
		}
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Errorf("Something went wrong: %v", err)
		return nil, &exceptions.DatabaseError{Err: errors.New(errorMessage)}
	}

	createdUser, err := srv.userRepository.Save(ctx, user)
	if err != nil {
		var serverSelectionError topology.ServerSelectionError
		if errors.As(err, &serverSelectionError) {
			logger.Errorf("Error connecting to database: %v", err)
			return nil, &exceptions.DatabaseError{Err: errors.New("Error connecting to database: ")}
		}
		logger.Errorf("Error persisting user: %v", err)
		return nil, &exceptions.UserValidationErr{Err: errors.New("Error persisting user: " + user.GetFullName())}
	}
	logger.Infof("User %s saved successfully", user.GetFullName())
	return createdUser, nil
}

func (srv *service) GetUserById(ctx context.Context, id string) (*models.User, error) {
	logger.Infof("Getting user by id %s", id)
	user, err := srv.userRepository.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Errorf("User not found: %v", err)
			return nil, &exceptions.UserNotFoundErr{Err: errors.New("User not found: " + id)}

		}
		logger.Errorf("Something went wrong: %v", err)
		return nil, &exceptions.DatabaseError{Err: errors.New(errorMessage)}
	}
	logger.Infof("User %s found successfully", id)
	return user, nil
}
