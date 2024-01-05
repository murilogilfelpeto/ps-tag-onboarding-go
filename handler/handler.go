package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/service"
)

var logger = configuration.NewLogger("handler")

type Handler interface {
	Save(context *gin.Context)
	FindById(context *gin.Context)
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{
		service: service,
	}
}
