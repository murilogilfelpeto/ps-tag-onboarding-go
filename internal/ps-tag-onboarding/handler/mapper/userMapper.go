package mapper

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/ps-tag-onboarding/handler/dto/request"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/ps-tag-onboarding/handler/dto/response"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/ps-tag-onboarding/service/models"
)

func UserRequestToUser(userRequest request.UserRequestDto) (models.User, error) {
	return models.NewUser("", userRequest.FirstName, userRequest.LastName, userRequest.Email, userRequest.Age)
}

func UserToUserResponseDto(user models.User) response.UserResponseDto {
	return response.UserResponseDto{
		ID:        user.GetID(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Email:     user.GetEmail(),
		Age:       user.GetAge(),
	}
}
