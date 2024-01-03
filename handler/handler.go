package handler

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service"
)

var (
	logger *configuration.Logger
)

func InitializeHandler() {
	logger = configuration.GetLogger("handler")
	service.Initialize()
}
