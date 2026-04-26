package handler

import (
	"github.com/momoyo-droid/pismo/api/internal/service"
	"go.uber.org/zap"
)

type TransactionHandler struct {
	TransactionService *service.Service
	Logger             *zap.Logger
}

func NewTransactionHandler(service *service.Service, logger *zap.Logger) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: service,
		Logger:             logger,
	}
}
