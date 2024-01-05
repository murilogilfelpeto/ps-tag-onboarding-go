package main

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/router"
)

var (
	logger *configuration.Logger
)

func main() {
	logger = configuration.GetLogger("main")
	logger.Info("Starting application...")

	err := configuration.Init()
	if err != nil {
		logger.Errorf("Error initializing application: %v", err)
		return
	}

	router.InitServer()
	logger.Info("Application started successfully")
}
