package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/handler/dto/request"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/handler/dto/response"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/handler/mapper"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	validator       = galidator.New()
	customValidator = validator.Validator(request.UserRequestDto{})
)

type Handler struct {
	service service.Service
}

func NewUserHandler(service service.Service) Handler {
	return Handler{
		service: service,
	}
}

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
func (h *Handler) Save(context *gin.Context) {
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
		context.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}
	user, err := mapper.UserRequestToUser(requestBody)
	if err != nil {
		logger.Error("Error while mapping request to user. ", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while mapping request to user",
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}
	createdUser, err := h.service.SaveUser(context, user)

	if err != nil {
		logger.Errorf("Something went wrong while persisting user in database. %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Something went wrong",
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if createdUser == nil {
		logger.Error("Error while persisting user. ", err)
		errorResponse := response.ErrorDto{
			Message:   "Error while creating user.",
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	responseBody := mapper.UserToUserResponseDto(*createdUser)
	context.IndentedJSON(http.StatusCreated, responseBody)
}

// FindById
// @Summary Find user by id
// @Description Find user by provided id
// @Tags Users
// @Produce json
// @Param id path string true "User id"
// @Failure 422 {object} response.ErrorDto "ID in the wrong format"
// @Failure 404 {object} response.ErrorDto "User not found"
// @Failure 500 {object} response.ErrorDto "Failed to connect to database"
// @Router /users/{id} [get]
func (h *Handler) FindById(context *gin.Context) {
	id := context.Param("id")
	logger.Infof("Finding user by id %s", id)
	user, err := h.service.GetUserById(context, id)

	if err != nil {
		logger.Errorf("Something went wrong in database. %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Something went wrong",
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if user == nil {
		logger.Errorf("Error while finding user by id %s.", id)
		errorResponse := response.ErrorDto{
			Message:   "No user found with id " + id,
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusNotFound, errorResponse)
		return
	}

	responseBody := mapper.UserToUserResponseDto(*user)
	context.IndentedJSON(http.StatusOK, responseBody)
}
