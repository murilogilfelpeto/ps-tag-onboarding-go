package router

import (
	"github.com/gin-gonic/gin"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
)

var (
	logger *configuration.Logger
)

func InitServer() {
	logger = configuration.GetLogger("router")
	router := gin.Default()

	initializeRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		logger.Errorf("Error starting server: %v", err)
		panic(err)
	}
}
