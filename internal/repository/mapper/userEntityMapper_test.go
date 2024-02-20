package mapper

import (
	"github.com/google/uuid"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/repository/entities"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserMapping(t *testing.T) {
	testCases := []struct {
		TestName string
		Entity   entities.UserEntity
		Expected func(entity entities.UserEntity, t *testing.T)
	}{
		{
			TestName: "When map user to user entity, then should return valid user entity",
			Entity:   entities.UserEntity{},
			Expected: func(entity entities.UserEntity, t *testing.T) {
				user, _ := models.NewUser("", "John", "Doe", "john.doe@email.com", 36)
				userEntity := UserToUserEntity(user)
				assert.NotNil(t, userEntity.ID)
				assert.Equal(t, user.GetFirstName(), userEntity.FirstName)
				assert.Equal(t, user.GetLastName(), userEntity.LastName)
				assert.Equal(t, user.GetEmail(), userEntity.Email)
				assert.Equal(t, user.GetAge(), userEntity.Age)
			},
		},
		{
			TestName: "When map user entity to user, then should return valid user",
			Entity: entities.UserEntity{
				ID:        uuid.New(),
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@email.com",
				Age:       36,
			},
			Expected: func(entity entities.UserEntity, t *testing.T) {
				user, _ := UserEntityToUser(entity)
				assert.Equal(t, user.GetID(), entity.ID.String())
				assert.Equal(t, user.GetFirstName(), entity.FirstName)
				assert.Equal(t, user.GetLastName(), entity.LastName)
				assert.Equal(t, user.GetEmail(), entity.Email)
				assert.Equal(t, user.GetAge(), entity.Age)
			},
		},
		{
			TestName: "When map user entity to user with invalid age, then should return error",
			Entity: entities.UserEntity{
				ID:        uuid.New(),
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@email.com",
				Age:       17,
			},
			Expected: func(entity entities.UserEntity, t *testing.T) {
				_, err := UserEntityToUser(entity)
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			tc.Expected(tc.Entity, t)
		})
	}
}
