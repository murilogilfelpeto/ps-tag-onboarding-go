package service

import "github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models"

func SaveUser(user models.User) (models.User, error) {
	newUser, err := models.NewUser("59849d54-4ff1-468c-adad-6e9d94f37311", user.GetFirstName(), user.GetLastName(), user.GetEmail(), user.GetAge())
	if err != nil {
		logger.Errorf("Error creating user: %v", err)
		return models.User{}, err
	}
	return newUser, nil
}
