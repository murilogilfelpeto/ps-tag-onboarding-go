package service

import (
	"context"
	"errors"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/mocks"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models/exceptions"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
	"testing"
)

func TestSaveUser(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	user, _ := models.NewUser("f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1", "John", "Doe", "john.doe@email.com", 36)

	mongo.ErrNoDocuments = errors.New("no documents")
	mockRepo.EXPECT().GetUserByFullName(context.Background(), user.GetFirstName(), user.GetLastName()).Return(nil, mongo.ErrNoDocuments).Once()
	mockRepo.EXPECT().Save(context.Background(), user).Return(&user, nil).Once()

	srv := &service{
		userRepository: mockRepo,
	}

	createdUser, err := srv.SaveUser(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, user, *createdUser)
}

func TestUserAlreadyExists(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	user, _ := models.NewUser("f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1", "John", "Doe", "john.doe@email.com", 36)

	mockRepo.EXPECT().GetUserByFullName(context.Background(), user.GetFirstName(), user.GetLastName()).Return(&user, nil).Once()

	srv := &service{
		userRepository: mockRepo,
	}

	createdUser, err := srv.SaveUser(context.Background(), user)
	userAlreadyExistsErr := &exceptions.UserAlreadyExistsErr{Err: errors.New("User already exists: " + user.GetFullName())}
	assert.Equal(t, userAlreadyExistsErr, err)
	assert.Nil(t, createdUser)
}

func TestDatabaseErrorGettingUserByFullName(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	user, _ := models.NewUser("f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1", "John", "Doe", "john.doe@email.com", 36)

	var expectedError topology.ServerSelectionError
	mockRepo.EXPECT().GetUserByFullName(context.Background(), user.GetFirstName(), user.GetLastName()).Return(nil, expectedError).Once()

	srv := &service{
		userRepository: mockRepo,
	}

	createdUser, err := srv.SaveUser(context.Background(), user)
	databaseError := &exceptions.DatabaseError{Err: errors.New("something went wrong")}
	assert.Equal(t, databaseError, err)
	assert.Nil(t, createdUser)
}
func TestErrorPersistingUser(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	user, _ := models.NewUser("f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1", "John", "Doe", "john.doe@email.com", 36)

	mongo.ErrNoDocuments = errors.New("no documents")
	mockRepo.EXPECT().GetUserByFullName(context.Background(), user.GetFirstName(), user.GetLastName()).Return(nil, mongo.ErrNoDocuments).Once()
	mockRepo.EXPECT().Save(context.Background(), user).Return(nil, errors.New("some error")).Once()

	srv := &service{
		userRepository: mockRepo,
	}

	createdUser, err := srv.SaveUser(context.Background(), user)
	assert.Error(t, err)
	assert.Nil(t, createdUser)
}

func TestServerSelectionErrorPersistingUser(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	user, _ := models.NewUser("f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1", "John", "Doe", "john.doe@email.com", 36)

	var expectedError topology.ServerSelectionError
	mockRepo.EXPECT().GetUserByFullName(context.Background(), user.GetFirstName(), user.GetLastName()).Return(nil, mongo.ErrNoDocuments).Once()
	mockRepo.EXPECT().Save(context.Background(), user).Return(nil, expectedError).Once()

	srv := &service{
		userRepository: mockRepo,
	}

	createdUser, err := srv.SaveUser(context.Background(), user)
	databaseError := &exceptions.DatabaseError{Err: errors.New("Error connecting to database: ")}
	assert.Equal(t, databaseError, err)
	assert.Nil(t, createdUser)
}
func TestGetUserById(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	id := "f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1"
	mockUser, _ := models.NewUser("f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1", "John", "Doe", "john.doe@email.com", 36)

	mockRepo.EXPECT().GetUserById(context.Background(), id).Return(&mockUser, nil).Once()

	srv := &service{
		userRepository: mockRepo,
	}

	user, err := srv.GetUserById(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, mockUser, *user)
}

func TestUserNotFound(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	id := "59ddb747-9767-4c1f-81b4-054877caf06d"

	mongo.ErrNoDocuments = errors.New("no documents")
	mockRepo.EXPECT().GetUserById(context.Background(), id).Return(nil, mongo.ErrNoDocuments).Once()

	srv := &service{
		userRepository: mockRepo,
	}

	user, err := srv.GetUserById(context.Background(), id)
	userNotFoundErr := &exceptions.UserNotFoundErr{Err: errors.New("User not found: " + id)}
	assert.Equal(t, userNotFoundErr, err)
	assert.Nil(t, user)
}

func TestDatabaseError(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	id := "59ddb747-9767-4c1f-81b4-054877caf06d"

	mockRepo.EXPECT().GetUserById(context.Background(), id).Return(nil, errors.New("some error")).Once()

	srv := &service{
		userRepository: mockRepo,
	}

	user, err := srv.GetUserById(context.Background(), id)
	databaseError := &exceptions.DatabaseError{Err: errors.New("something went wrong")}
	assert.Equal(t, databaseError, err)
	assert.Nil(t, user)
}
