package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler/dto/request"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler/dto/response"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler/mapper"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/mocks"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Persist User", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		mockService := &mocks.Service{}

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
		mockService.On("SaveUser", ctx, user).Return(createdUser, nil)

		handler := &Handler{
			service: mockService,
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
	})

	t.Run("Error Binding Json all fields", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		mockService := &mocks.Service{}

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
			service: mockService,
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

	})
	t.Run("Error Binding Json nil age and invalid email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		mockService := &mocks.Service{}

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
			service: mockService,
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

	})
	t.Run("Error persisting user", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		mockService := &mocks.Service{}

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
		mockService.On("SaveUser", ctx, user).Return(models.User{}, errors.New("some error"))

		handler := &Handler{
			service: mockService,
		}

		handler.Save(ctx)
		var responseBody response.ErrorDto
		_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Equal(t, "Error while creating user. some error", responseBody.Message)
		assert.NotEmpty(t, responseBody.Timestamp)
		assert.Nil(t, responseBody.Field)
	})
}

func TestFindById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Find user by id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		mockService := &mocks.Service{}
		id := uuid.New().String()

		pathParam := []gin.Param{
			{
				Key:   "id",
				Value: id,
			},
		}
		ctx.Params = pathParam

		user, _ := models.NewUser(id, "John", "Doe", "johndoe@email.com", 18)
		mockService.On("GetUserById", ctx, id).Return(user, nil)

		handler := &Handler{
			service: mockService,
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
	})
	t.Run("Error finding user", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		mockService := &mocks.Service{}
		id := uuid.New().String()

		pathParam := []gin.Param{
			{
				Key:   "id",
				Value: id,
			},
		}
		ctx.Params = pathParam

		mockService.On("GetUserById", ctx, id).Return(models.User{}, errors.New("some error"))

		handler := &Handler{
			service: mockService,
		}

		handler.FindById(ctx)
		var responseBody response.ErrorDto
		_ = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
		assert.Equal(t, http.StatusNotFound, recorder.Code)
		assert.Equal(t, "some error", responseBody.Message)
		assert.NotEmpty(t, responseBody.Timestamp)
		assert.Nil(t, responseBody.Field)
	})
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
