package router

import (
	"github.com/gin-gonic/gin"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler"
)

var logger = configuration.NewLogger("router")

type Router interface {
	InitServer()
	initializeRoutes(router *gin.Engine)
}

type router struct {
	handler handler.Handler
}

func NewRouter(handler handler.Handler) Router {
	return &router{
		handler: handler,
	}
}

func (r *router) InitServer() {
	logger = configuration.NewLogger("router")
	router := gin.Default()

	r.initializeRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		logger.Errorf("Error starting server: %v", err)
		panic(err)
	}
}
