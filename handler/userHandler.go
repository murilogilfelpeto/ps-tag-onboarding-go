package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler/dto/request"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler/dto/response"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler/mapper"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service"
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
	logger.Infof("Saving user...")

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
	user, err := mapper.UserRequestToUser(requestBody)
	if err != nil {
		logger.Error("Error while creating user. ", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while mapping request to user: " + err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}
	createdUser, err := service.SaveUser(user)
	if err != nil {
		logger.Error("Error while persisting user. ", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while creating user. " + err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	responseBody := mapper.UserToUserResponseDto(createdUser)
	context.IndentedJSON(http.StatusCreated, responseBody)
}

// FindById
// @Summary Find user by id
// @Description Find user by provided id
// @Tags Users
// @Produce json
// @Param id path int true "User id"
// @Success 201 {object} response.UserResponseDto "User Found successfully"
// @Failure 422 {object} response.ErrorDto "Id in the wrong format"
// @Failure 404 {object} response.ErrorDto "User not found"
// @Router /users/{id} [get]
func FindById(context *gin.Context) {
	id := context.Param("id")
	logger.Infof("Finding user by id %s", id)
	user, err := service.GetUserById(id)
	if err != nil {
		logger.Errorf("Error while finding user by id %s. %v", id, err)
		errorResponse := response.ErrorDto{
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusNotFound, errorResponse)
		return
	}
	responseBody := mapper.UserToUserResponseDto(user)
	context.IndentedJSON(http.StatusOK, responseBody)
}
