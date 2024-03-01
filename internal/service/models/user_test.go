package models

import (
	"errors"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_WhenUserIsCreated(t *testing.T) {
	tests := []struct {
		TestName     string
		Id           string
		FirstName    string
		LastName     string
		Email        string
		Age          int
		ExpectedUser User
		Error        error
	}{
		{
			TestName:  "when user has valid data, then error must be nil",
			Id:        "2814cd53-acde-4e49-9e47-cdc1d5dd48c7",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Age:       18,
			ExpectedUser: User{
				id:        "2814cd53-acde-4e49-9e47-cdc1d5dd48c7",
				firstName: "John",
				lastName:  "Doe",
				email:     "john.doe@example.com",
				age:       18,
			},
			Error: nil,
		},
		{
			TestName:     "When user has invalid uuid, then error must be UserValidationErr with message Invalid id",
			Id:           "12345",
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "john.doe@example.com",
			Age:          18,
			ExpectedUser: User{},
			Error:        &exceptions.UserValidationErr{Err: errors.New("Invalid id")},
		},
		{
			TestName:     "When user has invalid first name, then error must be UserValidationErr with message First name is required",
			Id:           "2814cd53-acde-4e49-9e47-cdc1d5dd48c7",
			FirstName:    "",
			LastName:     "Doe",
			Email:        "john.doe@example.com",
			Age:          18,
			ExpectedUser: User{},
			Error:        &exceptions.UserValidationErr{Err: errors.New("First name is required")},
		},
		{
			TestName:     "When user has invalid last name, then error must be UserValidationErr with message Last name is required",
			Id:           "2814cd53-acde-4e49-9e47-cdc1d5dd48c7",
			FirstName:    "John",
			LastName:     "",
			Email:        "john.doe@example.com",
			Age:          18,
			ExpectedUser: User{},
			Error:        &exceptions.UserValidationErr{Err: errors.New("Last name is required")},
		},
		{
			TestName:     "When user has invalid email, then error must be UserValidationErr with message Email is required",
			Id:           "2814cd53-acde-4e49-9e47-cdc1d5dd48c7",
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "",
			Age:          18,
			ExpectedUser: User{},
			Error:        &exceptions.UserValidationErr{Err: errors.New("Email is required")},
		},
		{
			TestName:     "When user has invalid age, then error must be UserValidationErr with message User must be at least 18 years old",
			Id:           "2814cd53-acde-4e49-9e47-cdc1d5dd48c7",
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "john.doe@example.com",
			Age:          17,
			ExpectedUser: User{},
			Error:        &exceptions.UserValidationErr{Err: errors.New("User must be at least 18 years old")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			user, err := NewUser(tt.Id, tt.FirstName, tt.LastName, tt.Email, tt.Age)
			assert.Equal(t, tt.ExpectedUser, user)
			assert.Equal(t, tt.Error, err)
		})
	}
}

func TestValidateUser_WhenCallValidateUserFunction(t *testing.T) {
	tests := []struct {
		TestName string
		User     User
		Error    error
	}{
		{
			TestName: "When user has valid data, then error must be nil",
			User: User{
				id:        "ab90b9a4-81c8-4199-a6eb-75b16cafc55a",
				firstName: "John",
				lastName:  "Doe",
				email:     "john.doe@example.com",
				age:       25,
			},
			Error: nil,
		},
		{
			TestName: "When user has invalid id, then error must be UserValidationErr with message Invalid id",
			User: User{
				id:        "12345",
				firstName: "John",
				lastName:  "Doe",
				email:     "john.doe@example.com",
				age:       18,
			},
			Error: &exceptions.UserValidationErr{Err: errors.New("Invalid id")},
		},
		{
			TestName: "When user has invalid first name, then error must be UserValidationErr with message First name is required",
			User: User{
				id:        "ab90b9a4-81c8-4199-a6eb-75b16cafc55a",
				firstName: "",
				lastName:  "Doe",
				email:     "john.doe@example.com",
				age:       18,
			},
			Error: &exceptions.UserValidationErr{Err: errors.New("First name is required")},
		},
		{
			TestName: "When user has invalid last name, then error must be UserValidationErr with message Last name is required",
			User: User{
				id:        "ab90b9a4-81c8-4199-a6eb-75b16cafc55a",
				firstName: "John",
				lastName:  "",
				email:     "john.doe@example.com",
				age:       18,
			},
			Error: &exceptions.UserValidationErr{Err: errors.New("Last name is required")},
		},
		{
			TestName: "When user has invalid email, then error must be UserValidationErr with message Email is required",
			User: User{
				id:        "ab90b9a4-81c8-4199-a6eb-75b16cafc55a",
				firstName: "John",
				lastName:  "Doe",
				email:     "",
				age:       18,
			},
			Error: &exceptions.UserValidationErr{Err: errors.New("Email is required")},
		},
		{
			TestName: "When user has invalid age, then error must be UserValidationErr with message User must be at least 18 years old",
			User: User{
				id:        "ab90b9a4-81c8-4199-a6eb-75b16cafc55a",
				firstName: "John",
				lastName:  "Doe",
				email:     "john.doe@example.com",
				age:       17,
			},
			Error: &exceptions.UserValidationErr{Err: errors.New("User must be at least 18 years old")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			err := validateUser(&tt.User)
			assert.Equal(t, tt.Error, err)
		})
	}
}
func TestUser_GetFullName_ShouldReturnFullName(t *testing.T) {
	user := &User{
		id:        "ab90b9a4-81c8-4199-a6eb-75b16cafc55a",
		firstName: "John",
		lastName:  "Doe",
		email:     "john.doe@example.com",
		age:       25,
	}

	assert.Equal(t, "John Doe", user.GetFullName())
}
