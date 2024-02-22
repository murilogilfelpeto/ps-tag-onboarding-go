package models

import (
	"errors"
	"github.com/google/uuid"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models/exceptions"
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
		return &exceptions.UserValidationErr{Err: errors.New("Invalid id")}
	}

	if user.firstName == "" {
		return &exceptions.UserValidationErr{Err: errors.New("First name is required")}
	}

	if user.lastName == "" {
		return &exceptions.UserValidationErr{Err: errors.New("Last name is required")}
	}

	if user.email == "" {
		return &exceptions.UserValidationErr{Err: errors.New("Email is required")}
	}

	if user.age < 18 {
		return &exceptions.UserValidationErr{Err: errors.New("User must be at least 18 years old")}
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

func (user User) GetFullName() string {
	return user.firstName + " " + user.lastName
}
