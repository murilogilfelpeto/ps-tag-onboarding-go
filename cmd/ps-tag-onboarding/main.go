package main

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/ps-tag-onboarding/handler"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/ps-tag-onboarding/repository"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/ps-tag-onboarding/router"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/ps-tag-onboarding/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Initializing application...")

	appConfig, err := configuration.Init()
	if err != nil {
		log.Fatalf("Error initializing application: %v", err)
	}

	log.Info("Starting application...")

	userRepository := repository.NewUserRepository(appConfig.Database, "onboarding", "users")
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	userRoute := router.NewRouter(userHandler)
	userRoute.InitServer()

	log.Info("Application started successfully")
}
