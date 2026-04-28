package handler

import (
	"github.com/momoyo-droid/pismo/api/internal/service"
	"go.uber.org/zap"
)

// Handler is the main entry point for all HTTP handlers in the application.
// It contains references to specific handlers for accounts and transactions.
type Handler struct {
	AccountHandler     *AccountHandler
	TransactionHandler *TransactionHandler
}

// NewHandler creates a new instance of Handler with the provided service and logger.
// It initializes the AccountHandler and TransactionHandler with the same service and logger.
func NewHandler(service *service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		AccountHandler:     NewAccountHandler(service, logger),
		TransactionHandler: NewTransactionHandler(service, logger),
	}
}
