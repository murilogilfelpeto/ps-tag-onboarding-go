package models

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser_WithValidData_ShouldSucceed(t *testing.T) {
	id := "2814cd53-acde-4e49-9e47-cdc1d5dd48c7"
	firstName := "John"
	lastName := "Doe"
	email := "john.doe@example.com"
	age := 18

	user, err := NewUser(id, firstName, lastName, email, age)

	assert.NoError(t, err)
	assert.Equal(t, id, user.GetID())
	assert.Equal(t, firstName, user.GetFirstName())
	assert.Equal(t, lastName, user.GetLastName())
	assert.Equal(t, email, user.GetEmail())
	assert.Equal(t, age, user.GetAge())
	assert.Equal(t, "John Doe", user.GetFullName())
}

func TestNewUser_WithInvalidData_ShouldFail(t *testing.T) {
	t.Run("User with invalid uuid", func(t *testing.T) {
		id := "12345"
		firstName := "John"
		lastName := "Doe"
		email := "john.doe@example.com"
		age := 25

		errMsg := &exceptions.UserValidationErr{Message: "Invalid id"}
		user, err := NewUser(id, firstName, lastName, email, age)

		assert.Error(t, err)
		assert.Equal(t, errMsg, err)
		assert.Equal(t, User{}, user)
	})
	t.Run("User with invalid first name", func(t *testing.T) {
		id := "20dc141d-3108-4529-8d65-ff3b046954be"
		firstName := ""
		lastName := "Doe"
		email := "john.doe@example.com"
		age := 25

		errMsg := &exceptions.UserValidationErr{Message: "First name is required"}
		user, err := NewUser(id, firstName, lastName, email, age)

		assert.Error(t, err)
		assert.Equal(t, errMsg, err)
		assert.Equal(t, User{}, user)
	})
	t.Run("User with invalid last name", func(t *testing.T) {
		id := "20dc141d-3108-4529-8d65-ff3b046954be"
		firstName := "John"
		lastName := ""
		email := "john.doe@example.com"
		age := 25

		errMsg := &exceptions.UserValidationErr{Message: "Last name is required"}
		user, err := NewUser(id, firstName, lastName, email, age)

		assert.Error(t, err)
		assert.Equal(t, errMsg, err)
		assert.Equal(t, User{}, user)
	})
	t.Run("User with invalid email", func(t *testing.T) {
		id := "20dc141d-3108-4529-8d65-ff3b046954be"
		firstName := "John"
		lastName := "Doe"
		email := ""
		age := 25

		errMsg := &exceptions.UserValidationErr{Message: "Email is required"}
		user, err := NewUser(id, firstName, lastName, email, age)

		assert.Error(t, err)
		assert.Equal(t, errMsg, err)
		assert.Equal(t, User{}, user)
	})
	t.Run("User with invalid age", func(t *testing.T) {
		id := "20dc141d-3108-4529-8d65-ff3b046954be"
		firstName := "John"
		lastName := "Doe"
		email := "johhn.doe@email.com"
		age := 17

		errMsg := &exceptions.UserValidationErr{Message: "User must be at least 18 years old"}
		user, err := NewUser(id, firstName, lastName, email, age)

		assert.Error(t, err)
		assert.Equal(t, errMsg, err)
		assert.Equal(t, User{}, user)
	})
}

func TestValidateUser_WithValidUser_ShouldSucceed(t *testing.T) {
	user := &User{
		id:        "ab90b9a4-81c8-4199-a6eb-75b16cafc55a",
		firstName: "John",
		lastName:  "Doe",
		email:     "john.doe@example.com",
		age:       25,
	}

	err := validateUser(user)

	assert.NoError(t, err)
}

func TestValidateUser_WithInvalidUser_ShouldFail(t *testing.T) {
	t.Run("User with invalid uuid", func(t *testing.T) {
		user := &User{
			id:        "12345",
			firstName: "John",
			lastName:  "Doe",
			email:     "john.doe@example.com",
			age:       25,
		}

		errMsg := &exceptions.UserValidationErr{Message: "Invalid id"}
		err := validateUser(user)

		assert.Error(t, err)
		assert.Equal(t, errMsg, err)
	})
	t.Run("User with invalid first name", func(t *testing.T) {
		user := &User{
			id:        "74c9e052-9fe9-4005-935a-e346d393f5ed",
			firstName: "",
			lastName:  "Doe",
			email:     "john.doe@example.com",
			age:       25,
		}

		errMsg := &exceptions.UserValidationErr{Message: "First name is required"}
		err := validateUser(user)

		assert.Error(t, err)
		assert.Equal(t, errMsg, err)
	})
	t.Run("User with invalid last name", func(t *testing.T) {
		user := &User{
			id:        "e59209a0-0cc9-46ff-80dc-ce75733b2bfd",
			firstName: "John",
			lastName:  "",
			email:     "john.doe@example.com",
			age:       25,
		}

		errMsg := &exceptions.UserValidationErr{Message: "Last name is required"}
		err := validateUser(user)

		assert.Error(t, err)
		assert.Equal(t, errMsg, err)
	})
	t.Run("User with invalid email", func(t *testing.T) {
		user := &User{
			id:        "1729f44d-1492-4279-80d5-8b05c42a5411",
			firstName: "John",
			lastName:  "Doe",
			email:     "",
			age:       25,
		}

		errMsg := &exceptions.UserValidationErr{Message: "Email is required"}
		err := validateUser(user)

		assert.Error(t, err)
		assert.Equal(t, errMsg, err)
	})
	t.Run("User with invalid age", func(t *testing.T) {
		user := &User{
			id:        "73f7bd4b-1997-4627-a3e3-6f592bd72d8c",
			firstName: "John",
			lastName:  "Doe",
			email:     "john.doe@example.com",
			age:       17,
		}

		errMsg := &exceptions.UserValidationErr{Message: "User must be at least 18 years old"}
		err := validateUser(user)

		assert.Error(t, err)
		assert.Equal(t, errMsg, err)
	})
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
