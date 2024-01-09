package service

import (
	"context"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/repository"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service/models"
)

var logger = configuration.NewLogger("service")

type Service interface {
	SaveUser(ctx context.Context, user models.User) (models.User, error)
	GetUserById(ctx context.Context, id string) (models.User, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}
