package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler/dto/request"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler/dto/response"
	"net/http"
	"time"
)

var (
	validator       = galidator.New()
	customValidator = validator.Validator(request.UserRequestDto{})
)

// Save
// @Summary Create a new user
// @Description Create a new user based on the provided user request data.
// @Tags Users
// @Accept json
// @Produce json
// @Param userRequest body request.UserRequestDto true "User data"
// @Success 201 {object} response.UserResponseDto "User created successfully"
// @Failure 422 {object} response.ErrorDto "Error while binding JSON or validation error"
// @Router /users [post]
func Save(context *gin.Context) {
	var requestBody request.UserRequestDto
	err := context.BindJSON(&requestBody)
	if err != nil {
		logger.Error("Error while binding request body: ", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while binding JSON",
			Timestamp: time.Now(),
			Field:     customValidator.DecryptErrors(err),
		}
		context.IndentedJSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	responseBody := response.UserResponseDto{
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		Email:     requestBody.Email,
		Age:       requestBody.Age,
		ID:        "8ce77f99-5684-4254-b34f-42d496ccab05",
	}
	context.IndentedJSON(http.StatusCreated, responseBody)
}

// FindById
// @Summary Find user by id
// @Description Find user by provided id
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 201 {object} response.UserResponseDto "User Found successfully"
// @Failure 422 {object} response.ErrorDto "Id in the wrong format"
// @Failure 404 {object} response.ErrorDto "User not found"
// @Router /users/{id} [get]
func FindById(context *gin.Context) {
	id := context.Param("id")
	responseBody := response.UserResponseDto{
		FirstName: "Murilo",
		LastName:  "Felpeto",
		Email:     "murilo@wexinc.com",
		Age:       30,
		ID:        id,
	}
	context.IndentedJSON(http.StatusOK, responseBody)
}
