package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/google/uuid"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler/dto/request"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler/dto/response"
	"net/http"
	"time"
)

var (
	validator       = galidator.New()
	customValidator = validator.Validator(request.UserRequestDto{})
)

func CreateUser(context *gin.Context) {
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
		ID:        uuid.New().String(),
	}
	context.IndentedJSON(http.StatusCreated, responseBody)
}
