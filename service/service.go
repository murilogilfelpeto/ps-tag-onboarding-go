package service

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/repository"
)

var (
	logger *configuration.Logger
)

func Initialize() {
	logger = configuration.GetLogger("service")
	logger.Info("Initializing service...")
	repository.Initialize()
}
