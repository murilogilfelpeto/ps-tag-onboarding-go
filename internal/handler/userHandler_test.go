package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/handler/dto/request"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/handler/dto/response"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/handler/mapper"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/mocks"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveUserSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService := mocks.NewService(t)

	requestBody := request.UserRequestDto{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@email.com",
		Age:       18,
	}

	jsonBody, _ := json.Marshal(requestBody)
	ctx.Request = &http.Request{
		Body: io.NopCloser(bytes.NewBuffer(jsonBody)),
	}

	user, _ := mapper.UserRequestToUser(requestBody)
	createdUser, _ := models.NewUser("f7d2ea4b-a4d0-4103-9c63-55ec7977e4d1", "John", "Doe", "john.doe@email.com", 36)
	mockService.EXPECT().SaveUser(ctx, user).Return(&createdUser, nil).Once()

	handler := &Handler{
		userService: mockService,
	}

	handler.Save(ctx)
	var responseBody response.UserResponseDto
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, createdUser.GetID(), responseBody.ID)
	assert.Equal(t, createdUser.GetFirstName(), responseBody.FirstName)
	assert.Equal(t, createdUser.GetLastName(), responseBody.LastName)
	assert.Equal(t, createdUser.GetEmail(), responseBody.Email)
	assert.Equal(t, createdUser.GetAge(), responseBody.Age)
}

func TestErrorPersistingUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService := mocks.NewService(t)

	requestBody := request.UserRequestDto{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@email.com",
		Age:       18,
	}

	jsonBody, _ := json.Marshal(requestBody)
	ctx.Request = &http.Request{
		Body: io.NopCloser(bytes.NewBuffer(jsonBody)),
	}

	user, _ := mapper.UserRequestToUser(requestBody)
	err := service.ErrDatabase
	mockService.EXPECT().SaveUser(ctx, user).Return(nil, err).Once()

	handler := &Handler{
		userService: mockService,
	}

	handler.Save(ctx)
	var responseBody response.ErrorDto
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, err.Error(), responseBody.Message)
	assert.NotEmpty(t, responseBody.Timestamp)
	assert.Nil(t, responseBody.Field)
}

func TestErrorUserAlreadyExists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService := mocks.NewService(t)

	requestBody := request.UserRequestDto{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@email.com",
		Age:       18,
	}

	jsonBody, _ := json.Marshal(requestBody)
	ctx.Request = &http.Request{
		Body: io.NopCloser(bytes.NewBuffer(jsonBody)),
	}

	user, _ := mapper.UserRequestToUser(requestBody)
	err := service.ErrUserAlreadyExists
	mockService.EXPECT().SaveUser(ctx, user).Return(nil, err).Once()

	handler := &Handler{
		userService: mockService,
	}

	handler.Save(ctx)
	var responseBody response.ErrorDto
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	assert.Equal(t, http.StatusConflict, recorder.Code)
	assert.Equal(t, err.Error(), responseBody.Message)
	assert.NotEmpty(t, responseBody.Timestamp)
	assert.Nil(t, responseBody.Field)
}

func TestErrorConnectingToDatabase(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService := mocks.NewService(t)

	requestBody := request.UserRequestDto{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@email.com",
		Age:       18,
	}

	jsonBody, _ := json.Marshal(requestBody)
	ctx.Request = &http.Request{
		Body: io.NopCloser(bytes.NewBuffer(jsonBody)),
	}

	user, _ := mapper.UserRequestToUser(requestBody)
	err := service.ErrDatabase
	mockService.EXPECT().SaveUser(ctx, user).Return(nil, err).Once()

	handler := &Handler{
		userService: mockService,
	}

	handler.Save(ctx)
	var responseBody response.ErrorDto
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, err.Error(), responseBody.Message)
	assert.NotEmpty(t, responseBody.Timestamp)
	assert.Nil(t, responseBody.Field)
}

func TestBindJsonFailure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService := mocks.NewService(t)

	requestBody := request.UserRequestDto{
		FirstName: "",
		LastName:  "",
		Email:     "",
		Age:       17,
	}

	jsonBody, _ := json.Marshal(requestBody)
	ctx.Request = &http.Request{
		Body: io.NopCloser(bytes.NewBuffer(jsonBody)),
	}

	handler := &Handler{
		userService: mockService,
	}

	handler.Save(ctx)
	var responseBody response.ErrorDto
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "Error while binding JSON", responseBody.Message)
	assert.NotEmpty(t, responseBody.Timestamp)
	assert.NotEmpty(t, responseBody.Field)

	fields, _ := convertToMap(responseBody.Field)
	assert.Equal(t, "required", fields["firstName"])
	assert.Equal(t, "required", fields["lastName"])
	assert.Equal(t, "required", fields["email"])
	assert.Equal(t, "age must be greater than 18", fields["age"])
}

func TestBindAgeAndEmailJsonFields(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService := mocks.NewService(t)

	requestBody := request.UserRequestDto{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "invalid-email",
	}

	jsonBody, _ := json.Marshal(requestBody)
	ctx.Request = &http.Request{
		Body: io.NopCloser(bytes.NewBuffer(jsonBody)),
	}

	handler := &Handler{
		userService: mockService,
	}

	handler.Save(ctx)
	var responseBody response.ErrorDto
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "Error while binding JSON", responseBody.Message)
	assert.NotEmpty(t, responseBody.Timestamp)
	assert.NotEmpty(t, responseBody.Field)

	fields, _ := convertToMap(responseBody.Field)
	assert.Equal(t, "not a valid email address", fields["email"])
	assert.Equal(t, "required", fields["age"])
}

func TestFindById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService := mocks.NewService(t)
	id := uuid.New().String()

	pathParam := []gin.Param{
		{
			Key:   "id",
			Value: id,
		},
	}
	ctx.Params = pathParam

	user, _ := models.NewUser(id, "John", "Doe", "johndoe@email.com", 18)
	mockService.EXPECT().GetUserById(ctx, id).Return(&user, nil).Once()

	handler := &Handler{
		userService: mockService,
	}

	handler.FindById(ctx)
	var responseBody response.UserResponseDto
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, user.GetID(), responseBody.ID)
	assert.Equal(t, user.GetFirstName(), responseBody.FirstName)
	assert.Equal(t, user.GetLastName(), responseBody.LastName)
	assert.Equal(t, user.GetEmail(), responseBody.Email)
	assert.Equal(t, user.GetAge(), responseBody.Age)
}

func TestUserDoesNotExist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService := mocks.NewService(t)
	id := uuid.New().String()

	pathParam := []gin.Param{
		{
			Key:   "id",
			Value: id,
		},
	}
	ctx.Params = pathParam

	err := service.ErrUserNotFound
	mockService.EXPECT().GetUserById(ctx, id).Return(nil, err).Once()

	handler := &Handler{
		userService: mockService,
	}

	handler.FindById(ctx)
	var responseBody response.ErrorDto
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, err.Error(), responseBody.Message)
	assert.NotEmpty(t, responseBody.Timestamp)
	assert.Nil(t, responseBody.Field)
}
func TestErrorFindingUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService := mocks.NewService(t)
	id := uuid.New().String()

	pathParam := []gin.Param{
		{
			Key:   "id",
			Value: id,
		},
	}
	ctx.Params = pathParam

	mockService.EXPECT().GetUserById(ctx, id).Return(nil, errors.New("some error")).Once()

	handler := &Handler{
		userService: mockService,
	}

	handler.FindById(ctx)
	var responseBody response.ErrorDto
	_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "Something went wrong", responseBody.Message)
	assert.NotEmpty(t, responseBody.Timestamp)
	assert.Nil(t, responseBody.Field)
}

func convertToMap(data interface{}) (map[string]string, error) {
	result := make(map[string]string)

	if dataMap, ok := data.(map[string]interface{}); ok {
		for key, value := range dataMap {
			if str, ok := value.(string); ok {
				result[key] = str
			} else {
				return nil, fmt.Errorf("value for key '%s' is not a string", key)
			}
		}
	} else {
		return nil, fmt.Errorf("data is not a map[string]interface{}")
	}

	return result, nil
}
