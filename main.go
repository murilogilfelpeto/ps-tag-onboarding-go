package main

import (
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/repository"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/router"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service"
)

func main() {
	appConfig, err := configuration.Init()
	if err != nil {
		panic(err)
	}
	logger := appConfig.Logger

	logger.Info("Starting application...")

	userRepository := repository.NewRepository(appConfig.Database, "onboarding", "users")
	userService := service.NewService(userRepository)
	userHandler := handler.NewHandler(userService)
	userRoute := router.NewRouter(userHandler)
	userRoute.InitServer()

	logger.Info("Application started successfully")
}
