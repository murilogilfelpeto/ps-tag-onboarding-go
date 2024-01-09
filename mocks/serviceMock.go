package mocks

import (
	"context"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (m *Service) SaveUser(ctx context.Context, user models.User) (models.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *Service) GetUserById(ctx context.Context, id string) (models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.User), args.Error(1)
}
