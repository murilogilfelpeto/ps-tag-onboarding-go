package router

import (
	"github.com/gin-gonic/gin"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler"
)

func initializeRoutes(router *gin.Engine) {
	logger.Info("Initializing routes...")
	handler.InitializeHandler()

	v1 := router.Group("/v1")
	v1.POST("/users", handler.CreateUser)
}
