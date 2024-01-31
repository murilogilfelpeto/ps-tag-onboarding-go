package router

import (
	"github.com/gin-gonic/gin"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/docs"
	logger "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (r *router) initializeRoutes(router *gin.Engine) {
	logger.Info("Initializing routes...")
	basePath := "/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	v1.POST("/users", r.handler.Save)
	v1.GET("/users/:id", r.handler.FindById)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
