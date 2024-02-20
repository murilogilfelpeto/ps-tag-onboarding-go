package mapper

import (
	"github.com/google/uuid"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/repository/entities"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapUserToUserEntity(t *testing.T) {
	user, _ := models.NewUser("", "John", "Doe", "john.doe@email.com", 36)
	userEntity := UserToUserEntity(user)
	assert.NotNil(t, userEntity.ID)
	assert.Equal(t, user.GetFirstName(), userEntity.FirstName)
	assert.Equal(t, user.GetLastName(), userEntity.LastName)
	assert.Equal(t, user.GetEmail(), userEntity.Email)
	assert.Equal(t, user.GetAge(), userEntity.Age)
}

func TestMapUserEntityToUser(t *testing.T) {
	userEntity := entities.UserEntity{
		ID:        uuid.New(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@email.com",
		Age:       36,
	}
	user, _ := UserEntityToUser(userEntity)
	assert.Equal(t, user.GetID(), userEntity.ID.String())
	assert.Equal(t, user.GetFirstName(), userEntity.FirstName)
	assert.Equal(t, user.GetLastName(), userEntity.LastName)
	assert.Equal(t, user.GetEmail(), userEntity.Email)
	assert.Equal(t, user.GetAge(), userEntity.Age)
}

func TestMapUserEntityToUser_WithInvalidData(t *testing.T) {
	userEntity := entities.UserEntity{
		ID:        uuid.New(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@email.com",
		Age:       17,
	}
	_, err := UserEntityToUser(userEntity)
	assert.Error(t, err)
}
