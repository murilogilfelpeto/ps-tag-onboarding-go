package router

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/configuration"
	"github.com/murilogilfelpeto/ps-tag-onboarding-go/handler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("Error starting server: %v", err)
			panic(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Errorf("Error shutting down server: %v", err)
		panic(err)
	}

	logger.Info("Server gracefully stopped")
}
