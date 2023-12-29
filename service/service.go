package service

import "github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"

var (
	logger *configuration.Logger
)

func Initialize() {
	logger = configuration.GetLogger("service")
}
