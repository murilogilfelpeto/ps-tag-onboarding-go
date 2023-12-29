package models

import (
	"github.com/google/uuid"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models/exceptions"
)

type User struct {
	id        string
	firstName string
	lastName  string
	email     string
	age       int
}

func NewUser(id string, firstName string, lastName string, email string, age int) (User, error) {
	newUser := &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		age:       age,
	}
	if err := validateUser(newUser); err != nil {
		return User{}, err
	}
	return *newUser, nil
}

func validateUser(user *User) error {
	if user.id != "" && uuid.Validate(user.id) != nil {
		return &exceptions.UserValidationErr{Message: "Invalid id"}
	}

	if user.firstName == "" {
		return &exceptions.UserValidationErr{Message: "First name is required"}
	}

	if user.lastName == "" {
		return &exceptions.UserValidationErr{Message: "Last name is required"}
	}

	if user.email == "" {
		return &exceptions.UserValidationErr{Message: "email is required"}
	}

	if user.age < 18 {
		return &exceptions.UserValidationErr{Message: "User must be at least 18 years old"}
	}

	return nil
}

func (user User) GetID() string {
	return user.id
}

func (user User) GetFirstName() string {
	return user.firstName
}

func (user User) GetLastName() string {
	return user.lastName
}

func (user User) GetEmail() string {
	return user.email
}

func (user User) GetAge() int {
	return user.age
}
