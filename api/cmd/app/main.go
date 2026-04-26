package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/momoyo-droid/pismo/api/internal/config"
	"github.com/momoyo-droid/pismo/api/internal/handler"
	"github.com/momoyo-droid/pismo/api/internal/repository"
	"github.com/momoyo-droid/pismo/api/internal/repository/postgres"
	"github.com/momoyo-droid/pismo/api/internal/service"
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

	db, err := postgres.NewDatabaseConnection(cfg)
	if err != nil {
		mainLogger.Fatal("Failed to connect to database", zap.Error(err))
	}

	accountRepository := repository.NewAccountRepository(db, mainLogger)
	transactionRepository := repository.NewTransactionRepository(db, mainLogger)
	service := service.NewService(accountRepository, transactionRepository, mainLogger)
	handler := handler.NewHandler(service, mainLogger)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.POST("/accounts", handler.AccountHandler.CreateAccount)

	mainLogger.Info("Starting API server...")
	if err := router.Run(":" + cfg.Port); err != nil {
		mainLogger.Fatal("Failed to start API server", zap.Error(err))
	}
}
