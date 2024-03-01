package service

import (
	"context"
	"errors"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/repository"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
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

var ErrUserAlreadyExists error = errors.New("user with first and last name already exists")
var ErrDatabase error = errors.New("error connecting to database")

var ErrUserNotFound error = errors.New("user not found")

func (srv *service) SaveUser(ctx context.Context, user models.User) (*models.User, error) {
	logger.Infof("Saving user %s", user.GetFullName())
	userByFullName, err := srv.userRepository.GetUserByFullName(ctx, user.GetFirstName(), user.GetLastName())

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Errorf("Something went wrong: %v", err)
		return nil, ErrDatabase
	}

	if userByFullName != nil {
		logger.Errorf("User already exists: %v", user.GetFullName())
		return nil, ErrUserAlreadyExists
	}

	createdUser, err := srv.userRepository.Save(ctx, user)
	if err != nil {
		return nil, ErrDatabase
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
			return nil, ErrUserNotFound

		}
		logger.Errorf("Something went wrong: %v", err)
		return nil, ErrDatabase
	}
	logger.Infof("User %s found successfully", id)
	return user, nil
}
