package handler

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
)

var (
	logger *configuration.Logger
)

func InitializeHandler() {
	logger = configuration.GetLogger("handler")
}
