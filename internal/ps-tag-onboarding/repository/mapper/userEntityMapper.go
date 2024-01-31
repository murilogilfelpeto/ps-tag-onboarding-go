package mapper

import (
	"github.com/google/uuid"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/ps-tag-onboarding/repository/entities"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/ps-tag-onboarding/service/models"
)

func UserToUserEntity(user models.User) entities.UserEntity {
	return entities.UserEntity{
		ID:        uuid.New(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Email:     user.GetEmail(),
		Age:       user.GetAge(),
	}
}

func UserEntityToUser(userEntity entities.UserEntity) (models.User, error) {
	return models.NewUser(
		userEntity.ID.String(),
		userEntity.FirstName,
		userEntity.LastName,
		userEntity.Email,
		userEntity.Age,
	)
}
