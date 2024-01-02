package router

import (
	"github.com/gin-gonic/gin"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/docs"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	logger.Info("Initializing routes...")
	handler.InitializeHandler()

	basePath := "/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	v1.POST("/users", handler.Save)
	v1.GET("/users/:id", handler.FindById)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
