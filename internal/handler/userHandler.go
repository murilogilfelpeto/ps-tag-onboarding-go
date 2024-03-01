package handler

import (
	"errors"
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
	userService service.UserService
}

func NewUserHandler(service service.UserService) Handler {
	return Handler{
		userService: service,
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
	createdUser, err := h.userService.SaveUser(context, user)

	if errors.Is(err, service.ErrUserAlreadyExists) {
		logger.Errorf("User already exists. %v", err)
		errorResponse := response.ErrorDto{
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusConflict, errorResponse)
		return
	}

	if err != nil {
		logger.Errorf("Error connecting to database. %v", err)
		errorResponse := response.ErrorDto{
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusInternalServerError, errorResponse)
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
	user, err := h.userService.GetUserById(context, id)

	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			logger.Errorf("User not found. %v", err)
			errorResponse := response.ErrorDto{
				Message:   err.Error(),
				Timestamp: time.Now(),
			}
			context.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		}

		logger.Errorf("Something went wrong in database. %v", err)
		errorResponse := response.ErrorDto{
			Message:   "Something went wrong",
			Timestamp: time.Now(),
		}
		context.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	responseBody := mapper.UserToUserResponseDto(*user)
	context.IndentedJSON(http.StatusOK, responseBody)
}
