package mocks

import (
	"context"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/ps-tag-onboarding/service/models"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (m *Repository) Save(ctx context.Context, user models.User) (models.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *Repository) GetUserById(ctx context.Context, id string) (models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *Repository) GetUserByFullName(ctx context.Context, firstName string, lastName string) (models.User, error) {
	args := m.Called(ctx, firstName, lastName)
	return args.Get(0).(models.User), args.Error(1)
}
