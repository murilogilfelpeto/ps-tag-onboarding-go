package main

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/repository"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/router"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service"
)

func main() {
	logger := configuration.NewLogger()
	logger.Info("Initializing application...")

	appConfig, err := configuration.Init()
	if err != nil {
		logger.Fatalf("Error initializing application: %v", err)
	}

	logger.Info("Starting application...")

	userRepository := repository.NewUserRepository(appConfig.Database, "onboarding", "users")
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	userRoute := router.NewRouter(userHandler)
	userRoute.InitServer()

	logger.Info("Application started successfully")
}
