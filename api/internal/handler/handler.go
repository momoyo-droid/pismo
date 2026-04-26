package handler

import (
	"github.com/momoyo-droid/pismo/api/internal/service"
	"go.uber.org/zap"
)

type Handler struct {
	AccountHandler     *AccountHandler
	TransactionHandler *TransactionHandler
}

func NewHandler(service *service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		AccountHandler:     NewAccountHandler(service, logger),
		TransactionHandler: NewTransactionHandler(service, logger),
	}
}
