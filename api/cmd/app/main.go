package main

import (
	"github.com/gin-gonic/gin"
	"github.com/momoyo-droid/pismo/api/internal/config"
	"github.com/momoyo-droid/pismo/api/internal/handler"
	"go.uber.org/zap"
)

func logger() *zap.Logger {
	logger := zap.Must(zap.NewDevelopment())

	return logger
}

func main() {

	mainLogger := logger()

	defer func() {
		_ = mainLogger.Sync()
	}()

	cfg, err := config.LoadConfig()

	if err != nil {
		mainLogger.Fatal("Failed to load configuration", zap.Error(err))
	}

	router := gin.Default()

	router.GET("/health", handler.HealthCheck)

	mainLogger.Info("Starting API server...")
	if err := router.Run(":" + cfg.Port); err != nil {
		mainLogger.Fatal("Failed to start API server", zap.Error(err))
	}
}
