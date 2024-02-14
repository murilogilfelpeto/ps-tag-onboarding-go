package service

import (
	"context"
	"errors"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/mocks"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestSaveUser(t *testing.T) {
	t.Run("Persist User", func(t *testing.T) {
		mockRepo := &mocks.Repository{}

		user, _ := models.NewUser("f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1", "John", "Doe", "john.doe@email.com", 36)

		mongo.ErrNoDocuments = errors.New("no documents")
		mockRepo.EXPECT().GetUserByFullName(context.Background(), user.GetFirstName(), user.GetLastName()).Return(nil, mongo.ErrNoDocuments).Once()
		mockRepo.EXPECT().Save(context.Background(), user).Return(&user, nil).Once()

		srv := &service{
			repository: mockRepo,
		}

		createdUser, err := srv.SaveUser(context.Background(), user)
		assert.NoError(t, err)
		assert.Equal(t, user, *createdUser)
		mockRepo.AssertExpectations(t)
	})

	// Test when user already exists in the repository
	t.Run("User Already Exists", func(t *testing.T) {
		mockRepo := &mocks.Repository{}

		user, _ := models.NewUser("f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1", "John", "Doe", "john.doe@email.com", 36)

		mockRepo.EXPECT().GetUserByFullName(context.Background(), user.GetFirstName(), user.GetLastName()).Return(&user, nil).Once()

		srv := &service{
			repository: mockRepo,
		}

		createdUser, err := srv.SaveUser(context.Background(), user)
		assert.Nil(t, err)
		assert.Nil(t, createdUser)
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNumberOfCalls(t, "Save", 0)
	})

	// Test when there is an error persisting the user
	t.Run("Error Persisting User", func(t *testing.T) {
		mockRepo := &mocks.Repository{}

		user, _ := models.NewUser("f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1", "John", "Doe", "john.doe@email.com", 36)

		mongo.ErrNoDocuments = errors.New("no documents")
		mockRepo.EXPECT().GetUserByFullName(context.Background(), user.GetFirstName(), user.GetLastName()).Return(nil, mongo.ErrNoDocuments).Once()
		mockRepo.EXPECT().Save(context.Background(), user).Return(nil, errors.New("some error")).Once()

		srv := &service{
			repository: mockRepo,
		}

		createdUser, err := srv.SaveUser(context.Background(), user)
		assert.Error(t, err)
		assert.Nil(t, createdUser)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserById(t *testing.T) {
	t.Run("User found", func(t *testing.T) {
		mockRepo := &mocks.Repository{}

		id := "f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1"
		mockUser, _ := models.NewUser("f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1", "John", "Doe", "john.doe@email.com", 36)

		mockRepo.EXPECT().GetUserById(context.Background(), id).Return(&mockUser, nil).Once()

		srv := &service{
			repository: mockRepo,
		}

		user, err := srv.GetUserById(context.Background(), id)
		assert.NoError(t, err)
		assert.Equal(t, mockUser, *user)
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNumberOfCalls(t, "Save", 0)
		mockRepo.AssertNumberOfCalls(t, "GetUserByFullName", 0)
	})

	t.Run("User not found", func(t *testing.T) {
		mockRepo := &mocks.Repository{}

		id := "59ddb747-9767-4c1f-81b4-054877caf06d"

		mockRepo.EXPECT().GetUserById(context.Background(), id).Return(nil, nil).Once()

		srv := &service{
			repository: mockRepo,
		}

		user, err := srv.GetUserById(context.Background(), id)
		assert.Nil(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNumberOfCalls(t, "Save", 0)
		mockRepo.AssertNumberOfCalls(t, "GetUserByFullName", 0)
	})

	t.Run("Database error", func(t *testing.T) {
		mockRepo := &mocks.Repository{}

		id := "59ddb747-9767-4c1f-81b4-054877caf06d"

		mockRepo.EXPECT().GetUserById(context.Background(), id).Return(nil, errors.New("some error")).Once()

		srv := &service{
			repository: mockRepo,
		}

		user, err := srv.GetUserById(context.Background(), id)
		assert.Error(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNumberOfCalls(t, "Save", 0)
		mockRepo.AssertNumberOfCalls(t, "GetUserByFullName", 0)
	})
}
